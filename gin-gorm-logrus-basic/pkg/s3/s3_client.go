package s3

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Client S3 Client
type Client interface {
	Put(file *File) error
	List(prefix string) ([]string, error)
	Get(key string) (*File, error)
	IsModified(key string, time time.Time) (bool, error)
	Head(key string) (*s3.HeadObjectOutput, error)
}

type s3ClientImpl struct {
	s3     *s3.S3
	bucket string
}

// NewS3Client S3Client を Config から生成
func NewS3Client(cfg *Config) (Client, error) {
	awsConfig := &aws.Config{
		Region: aws.String(cfg.Region),
	}
	if cfg.AccessKeyID != "" && cfg.SecretAccessKey != "" {
		awsConfig.Credentials = credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, "")
	}
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}
	return &s3ClientImpl{s3: s3.New(sess), bucket: cfg.BucketName}, nil
}

// Put File を S3 に配置
func (s *s3ClientImpl) Put(file *File) error {
	_, err := s.s3.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(file.Name),
		Body:        bytes.NewReader(file.Body),
		ContentType: aws.String(file.ContentType),
	})
	return err
}

// List S3 オブジェクトキーの一覧を取得
func (s *s3ClientImpl) List(prefix string) ([]string, error) {
	req := &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(prefix),
	}
	keys := make([]string, 0)
	err := s.s3.ListObjectsV2Pages(req, func(page *s3.ListObjectsV2Output, isLastPage bool) bool {
		for _, c := range page.Contents {
			keys = append(keys, *c.Key)
		}
		return !isLastPage // return true if we should continue with the next page
	})
	if err != nil {
		return nil, err
	}
	return keys, nil
}

// Get S3 オブジェクトキーから S3 オブジェクトを取得
func (s *s3ClientImpl) Get(key string) (*File, error) {
	resp, getErr := s.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if getErr != nil {
		return nil, getErr
	}
	defer closeS3(resp.Body)
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	return &File{
		Name:         key,
		Body:         body,
		ContentType:  *resp.ContentType,
		LastModified: *resp.LastModified,
	}, nil
}

// IsModified S3 オブジェクトが変更されていれば true を返却
func (s *s3ClientImpl) IsModified(key string, time time.Time) (bool, error) {
	headRes, err := s.Head(key)
	if err != nil {
		return false, err
	}
	return headRes.LastModified.After(time), nil
}

// Head S3 オブジェクトのメタデータを取得
func (s *s3ClientImpl) Head(key string) (*s3.HeadObjectOutput, error) {
	resp, err := s.s3.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return resp, err
}

// closeS3 S3 コネクションを終了
// s3 パッケージ内部からの呼び出しを想定
func closeS3(respBody io.ReadCloser) {
	if err := respBody.Close(); err != nil {
		log.Fatalf("closing S3 error: %v", err)
	}
}
