package configs

import "gopkg.in/yaml.v3"

type (
	// DBConfs contains all environmental DB configs.
	DBConfs []*DBConf

	// DBConf is the config for DB.
	DBConf struct {
		Env         string  `yaml:"env"` // TODO: enum
		TwtrDataSrc DataSrc `yaml:"twtr"`
	}

	// DataSrc is DB data source
	DataSrc struct {
		Driver             string `yaml:"driver"`
		UseConnPool        bool   `yaml:"useConnPool"`
		MaxIdleConnSize    int    `yaml:"maxIdleConnSize"`
		MaxOpenConnSize    int    `yaml:"maxOpenConnSize"`
		ConnMaxLifetimeMin int    `yaml:"connMaxLifetimeMin"`
		Reader             Reader `yaml:"reader"`
		Writer             Writer `yaml:"writer"`
	}

	// Reader is DB settings for Read Connection.
	Reader struct {
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		Params string `yaml:"params"`
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
	}

	// Writer is DB settings for Read Connection.
	Writer struct {
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		Params string `yaml:"params"`
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
	}
)

//go:embed dbconf.yml
var ymlDBConfs []byte

// UnmarshalDBConfs scans DBConfs from yml
func UnmarshalDBConfs() (DBConfs, error) {
	var dbConfs DBConfs
	if err := yaml.Unmarshal(ymlDBConfs, &dbConfs); err != nil {
		return nil, err
	}
	return dbConfs, nil
}
