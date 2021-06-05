package appconf

import (
	"github.com/imdario/mergo"

	"github.com/krtsato/go-server-templates/2021-05-twtr/configs"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/apperr"
)

// loadDBConf returns DBConf loaded from yml.
func loadDBConf(e AppEnv) (*configs.DBConf, error) {
	dbConfs, err := configs.UnmarshalDBConfs()
	if err != nil {
		return nil, err
	}

	dbConf, err := mergeDBConf(e, dbConfs)
	if err != nil {
		return nil, err
	}

	return dbConf, nil
}

// mergeDBConf integrates default DB config with target one.
// zero values in targetConfig are overwritten by values in defaultConfig.
func mergeDBConf(e AppEnv, cs configs.DBConfs) (*configs.DBConf, error) {
	var defaultConf, targetConf, emptyConf configs.DBConf

	for _, c := range cs {
		if c.Env == e.String() {
			targetConf = *c
		} else if c.Env == "default" {
			defaultConf = *c
		}
	}

	if targetConf == emptyConf {
		return nil, apperr.ErrorF(apperr.Config, "failed to merge default DB config with empty one")
	}

	if err := mergo.Merge(&targetConf, defaultConf, mergo.WithTransformers(boolTransformer{})); err != nil {
		return nil, apperr.ErrorF(apperr.Config, "failed to merge default DB config with %s one: %s", e.String(), err)
	}

	return &targetConf, nil
}
