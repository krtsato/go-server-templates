package config

import (
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/apperr"
	"os"
	"strings"
)

// AppEnv Enum
type AppEnv int

// AppEnv 種類
const (
	_ AppEnv = iota // 0 無効値
	Local
	Dev
	Prd
)

var appEnvValueMap = map[AppEnv]string{
	Local: "local",
	Dev:   "dev",
	Prd:   "prd",
}

// String stringify AppEnv
func (e AppEnv) String() string {
	return appEnvValueMap[e]
}

// ApplyAppEnv 入力文字列に応じた AppEnv を返却
func ApplyAppEnv(value string) (AppEnv, error) {
	for k, v := range appEnvValueMap {
		if strings.EqualFold(v, value) {
			return k, nil
		}
	}
	return 0, apperr.NewConfigErrorF("unknown AppEnv %s", value)
}

// GetOSEnv get os env from APP_ENV
func GetOSEnv() string {
	return os.Getenv("APP_ENV")
}
