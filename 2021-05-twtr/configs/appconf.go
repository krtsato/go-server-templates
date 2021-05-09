package configs

import (
	_ "embed" //nolint

	"gopkg.in/yaml.v3"
)

type (
	// AppConfs contains all environmental app configs.
	AppConfs []*AppConf

	// AppConf is the config for app.
	AppConf struct {
		env string `yaml:"env"`
		// logger Logger `yaml:"logger"`
		webAPI WebAPI `yaml:"webapi"`
	}

	// Logger is the logger setting.
	//Logger struct {
	//	level    string `yaml:"level"`
	//	format   string `yaml:"format"`
	//	fullPath string `yaml:"fullPath"`
	//}

	// WebAPI is the webapi setting.
	WebAPI struct {
		port string `yaml:"port"`
		//auth bool   `yaml:"auth"`
	}
)

//go:embed appconf.yml
var ymlAppConfs []byte

// UnmarshalAppConfs scans AppConfs from yml.
func UnmarshalAppConfs() (AppConfs, error) {
	var appConfs AppConfs
	if err := yaml.Unmarshal(ymlAppConfs, &appConfs); err != nil {
		return nil, err
	}

	return appConfs, nil
}

// Env returns app env
func (a AppConf) Env() string {
	return a.env
}

// WebAPIPort returns the webapi port
func (a AppConf) WebAPIPort() string {
	return a.webAPI.port
}
