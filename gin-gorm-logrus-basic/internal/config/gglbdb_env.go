package config

import "github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/infra"

// DB database env
type DB struct {
	DBEnv infra.DBEnv `yaml:"dbEnv"`
	// GglbDB infra.Datasource `yaml:"gglbdb"`
}
