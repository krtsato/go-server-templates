package configs

import (
	_ "embed" //nolint

	"go.uber.org/zap"

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/apperr"

	"gopkg.in/yaml.v3"
)

type (
	// AppConfs contains all environmental app configs.
	AppConfs []*AppConf

	// AppConf is the config for app.
	AppConf struct {
		Env  string      `yaml:"env"`
		Log  *zap.Config `yaml:"log"`
		REST REST        `yaml:"rest"`
	}

	// REST is the rest api setting.
	REST struct {
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
		return nil, apperr.ErrorF(apperr.Unmarshal, "failed to unmarshal app config: %s", err)
	}

	return appConfs, nil
}
