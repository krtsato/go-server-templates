package config

import "github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/infra"

// DB database env
type DB struct {
	Env    infra.DBEnv      `yaml:"env"`
	GglbDB infra.Datasource `yaml:"gglbdb"`
}
