package configs

import (
	_ "embed" //nolint

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/apperr"

	"gopkg.in/yaml.v3"
)

type (
	// AppConfs contains all environmental app configs.
	AppConfs []*AppConf

	// AppConf is the config for app.
	AppConf struct {
		Env    string `yaml:"env"`
		Logger Logger `yaml:"logger"`
		WebAPI WebAPI `yaml:"webapi"`
	}

	// Logger is the logger setting.
	Logger struct {
		Level    string `yaml:"level"`
		Format   string `yaml:"format"`
		FullPath string `yaml:"fullPath"`
	}

	// WebAPI is the webapi setting.
	WebAPI struct {
		Port string `yaml:"port"`
		Auth bool   `yaml:"auth"`
	}
)

//go:embed appconf.yml
var ymlAppConfs []byte

// UnmarshalAppConfs scans AppConfs from yml.
func UnmarshalAppConfs() (AppConfs, error) {
	var appConfs AppConfs
	if err := yaml.Unmarshal(ymlAppConfs, &appConfs); err != nil {
		return nil, apperr.ErrorF(apperr.Unmarshal, "failed to unmarshal app config: %s", err.Error())
	}

	return appConfs, nil
}
