package appconf

import (
	"os"
	"strings"

	"github.com/krtsato/go-server-templates/202105-twtr/pkg/apperr"
)

// AppEnv is enum.
type AppEnv int

// AppEnv is constant.
const (
	AppEnvUnknown AppEnv = iota
	AppEnvLocal
	AppEnvDev
	AppEnvPrd
)

var appEnvValueMap = map[AppEnv]string{
	AppEnvLocal: "local",
	AppEnvDev:   "dev",
	AppEnvPrd:   "prd",
}

// String returns AppEnv value.
func (e AppEnv) String() string {
	return appEnvValueMap[e]
}

// confirmAppEnv returns AppEnv which equals to the argument.
func confirmAppEnv(value string) (AppEnv, error) {
	for k, v := range appEnvValueMap {
		if strings.EqualFold(v, value) {
			return k, nil
		}
	}

	return AppEnvUnknown, apperr.Errorf(apperr.Config, "failed to confirm the unknown AppEnv: %s", value)
}

// loadOSAppEnv returns OS APP_ENV value.
func loadOSAppEnv() string {
	return os.Getenv("APP_ENV")
}
