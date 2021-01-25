package config

import (
	"os"
	"strings"

	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/apperr"
)

// AppEnv Enum
type AppEnv int

// AppEnv 種類
const (
	UnknownEnv AppEnv = iota // 無効値
	Local
	Dev
	Prd
)

var appEnvValueMap = map[AppEnv]string{
	Local: "local",
	Dev:   "dev",
	Prd:   "prd",
}

// String AppEnv に応じた文字列を返却
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
	return UnknownEnv, apperr.NewConfigErrF("unknown AppEnv %s", value)
}

// GetOSEnv 環境変数から APP_ENV を取得
func GetOSEnv() string {
	return os.Getenv("APP_ENV")
}
