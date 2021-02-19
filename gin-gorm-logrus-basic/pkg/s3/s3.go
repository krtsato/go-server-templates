package s3

import "time"

// Config S3 の設定値
type Config struct {
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	Region          string `yaml:"region" validate:"required"`
	BucketName      string `yaml:"bucket_name" validate:"required"`
}

// Object S3 に配置したファイル
type Object struct {
	Name         string
	Body         []byte
	ContentType  string
	LastModified time.Time
}
