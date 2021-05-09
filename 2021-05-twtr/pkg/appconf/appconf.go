package appconf

import (
	"reflect"

	"github.com/krtsato/go-server-templates/2021-05-twtr/configs"

	"github.com/imdario/mergo"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/apperr"
)

// boolTransformer extends mergo.merge.
type boolTransformer struct{}

// loadAppConf returns AppConf loaded from yml.
func loadAppConf(e AppEnv) (*configs.AppConf, error) {
	appConfs, err := configs.UnmarshalAppConfs()
	if err != nil {
		return nil, apperr.ErrorF(apperr.Config, "failed to unmarshal app config")
	}

	mergedConf, err := merge(e, appConfs)
	if err != nil {
		return nil, err
	}

	return mergedConf, nil
}

// merge integrates default config with target one.
// zero values in targetConfig are overwritten by values in defaultConfig.
func merge(e AppEnv, cs configs.AppConfs) (*configs.AppConf, error) {
	var defaultConf, targetConf, emptyConf configs.AppConf

	for _, c := range cs {
		if c.Env() == e.String() {
			targetConf = *c
		} else if c.Env() == "default" {
			defaultConf = *c
		}
	}

	if targetConf == emptyConf {
		return nil, apperr.ErrorF(apperr.Config, "failed to merge default config with empty one")
	}

	if err := mergo.Merge(&targetConf, defaultConf, mergo.WithTransformers(boolTransformer{})); err != nil {
		return nil, apperr.ErrorF(apperr.Config, "failed to merge default config with %s one", e.String())
	}

	return &targetConf, nil
}

// Transformer expresses duck typing for mergo.merge because of overwriting WebAPI.Auth with zero value.
func (b boolTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	switch typ {
	case reflect.TypeOf(configs.WebAPI{}):
		return func(dst, src reflect.Value) error { return nil }
	default:
		return nil
	}
}
