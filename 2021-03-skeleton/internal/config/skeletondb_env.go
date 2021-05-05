package config

import "github.com/krtsato/go-server-templates/2021-03-skeleton/pkg/infra"

// DB database env
type DB struct {
	Env        infra.DBEnv      `yaml:"env"`
	SkeletonDB infra.Datasource `yaml:"skeletondb"`
}
