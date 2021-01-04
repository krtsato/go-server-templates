package config

import (
	"fmt"
	"github.com/imdario/mergo"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/apperr"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const (
	configDir     = "configs"
	appConfigPath = configDir + "/application.yml"
)

// AppConfig 環境単位の設定値
type AppConfig struct {
	AppEnv string `yaml:"appEnv"`
	// DB  GglbDB      `yaml:"db"`
	// Log log.Config `yaml:"log"`
	// Web Web        `yaml:"web"`
}

// AppConfigs 全環境の設定値群
type AppConfigs []AppConfig

// targetConfig ゼロ値フィールドにのみ defaultConfig をマージ
func (configs AppConfigs) merge(e AppEnv) (AppConfig, error) {
	var defaultConfig, targetConfig, emptyConfig AppConfig

	for _, config := range configs {
		if config.AppEnv == e.String() {
			targetConfig = config
		} else if config.AppEnv == "default" {
			defaultConfig = config
		}
	}
	if emptyConfig == targetConfig {
		return AppConfig{}, apperr.NewConfigErrorF("unknown profile %s", e.String())
	}

	err := mergo.Merge(&targetConfig, defaultConfig)
	return targetConfig, err
}

// LoadConfig アプリケーションで使用する Config を決定
func LoadConfig(ae AppEnv) (AppConfig, error) {
	fmt.Printf("start to load config... %v\n", ae)

	ymlConfigs, readErr := ioutil.ReadFile(appConfigPath)
	if readErr != nil {
		return AppConfig{}, readErr
	}

	appConfigs := new(AppConfigs)
	if unmErr := yaml.Unmarshal(ymlConfigs, &appConfigs); unmErr != nil {
		return AppConfig{}, unmErr
	}

	mergedConfig, mergeErr := appConfigs.merge(ae)
	if mergeErr != nil {
		return AppConfig{}, mergeErr
	}

	fmt.Printf("Succeed at %v config setup\n", mergedConfig.AppEnv)
	return mergedConfig, nil
}
