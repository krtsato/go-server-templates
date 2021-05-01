package infra

import (
	"strings"

	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/apperr"
)

// DBEnv Enum
type DBEnv int

// DBEnv 種別
const (
	UnknownEnv DBEnv = iota
	Local
	Dev
	Prd
)

var dbEnvStringMap = map[DBEnv]string{
	Local: "local",
	Dev:   "dev",
	Prd:   "prd",
}

// String DBEnv に応じた文字列を返却
func (e DBEnv) String() string {
	return dbEnvStringMap[e]
}

// ApplyDBEnv 入力文字列に応じた DBEnv を返却
func ApplyDBEnv(value string) (DBEnv, error) {
	for k, v := range dbEnvStringMap {
		if strings.EqualFold(v, value) {
			return k, nil
		}
	}
	return UnknownEnv, apperr.NewConfigErrF("unknown DBEnv %s", value)
}

// UnmarshalYAML Yaml ファイルを読み込むためダックタイピング
func (e *DBEnv) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}

	dbEnv, err := ApplyDBEnv(str)
	if err != nil {
		return err
	}

	*e = dbEnv
	return nil
}
