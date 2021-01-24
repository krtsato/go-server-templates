package config

import (
	"fmt"
	"io/ioutil"

	"github.com/imdario/mergo"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/apperr"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/logger"
	"gopkg.in/yaml.v3"
)

const (
	configDir     = "configs"
	appConfigPath = configDir + "/application.yml"
)

// AppConfig 環境単位の設定値
type AppConfig struct {
	AppEnv string `yaml:"appEnv"`
	// DB  GglbDB      `yaml:"db"`
	Log logger.Config `yaml:"log"`
	Web Web           `yaml:"web"`
}

// AppConfigs 全環境の設定値群
type AppConfigs []*AppConfig

// LoadConfig アプリケーションで使用する Config を決定
func LoadConfig(e AppEnv) (*AppConfig, error) {
	fmt.Printf("start to load config... %v\n", e)

	ymlConfigs, readErr := ioutil.ReadFile(appConfigPath)
	if readErr != nil {
		return &AppConfig{}, readErr
	}

	appConfigs := new(AppConfigs)
	if err := yaml.Unmarshal(ymlConfigs, appConfigs); err != nil {
		return &AppConfig{}, err
	}

	mergedConfig, mergeErr := appConfigs.merge(e)
	if mergeErr != nil {
		return &AppConfig{}, mergeErr
	}

	fmt.Printf("succeed at %v config setup\n", mergedConfig.AppEnv)
	return mergedConfig, nil
}

// targetConfig ゼロ値フィールドにのみ defaultConfig をマージ
func (configs AppConfigs) merge(e AppEnv) (*AppConfig, error) {
	var defaultConfig, targetConfig, emptyConfig *AppConfig
	for _, config := range configs {
		if config.AppEnv == e.String() {
			targetConfig = config
		} else if config.AppEnv == "default" {
			defaultConfig = config
		}
	}

	if targetConfig == emptyConfig {
		return &AppConfig{}, apperr.NewConfigErrF("unknown profile %s", e.String())
	}
	if err := mergo.Merge(targetConfig, defaultConfig); err != nil {
		return &AppConfig{}, apperr.NewConfigErr(err, "failed to Merge by mergo")
	}
	return targetConfig, nil
}
