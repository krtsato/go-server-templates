package configs

import (
	_ "embed" //nolint

	"go.uber.org/zap"

	"github.com/krtsato/go-server-templates/202105-twtr/pkg/apperr"

	"gopkg.in/yaml.v3"
)

type (
	// AppConfs contains all environmental app configs.
	AppConfs []*AppConf

	// AppConf is the config for app.
	AppConf struct {
		Env  string      `yaml:"env"`
		Log  *zap.Config `yaml:"log"`
		Rest Rest        `yaml:"rest"`
	}

	// Rest is the rest api setting.
	Rest struct {
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
		return nil, apperr.Errorf(apperr.Unmarshal, "failed to unmarshal app config: %s", err)
	}

	return appConfs, nil
}
