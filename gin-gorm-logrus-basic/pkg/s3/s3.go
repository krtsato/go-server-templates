package s3

import "time"

type Config struct {
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	Region          string `yaml:"region" validate:"required"`
	BucketName      string `yaml:"bucket_name" validate:"required"`
}

type File struct {
	Name         string
	Body         []byte
	ContentType  string
	LastModified time.Time
}
