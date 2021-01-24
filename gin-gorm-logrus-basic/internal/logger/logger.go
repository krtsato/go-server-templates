package logger

import (
	"strings"

	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/apperr"

	"github.com/sirupsen/logrus"
)

// Config ログ設定
type Config struct {
	Level    Level  `yaml:"level"`
	Format   Format `yaml:"format"`
	FullPath string `yaml:"fullPath"`
}

// ------------------------------------------------------------
// Level
// ------------------------------------------------------------

// Level ログ出力レベル
type Level int

// Level 種類
const (
	UnknownLevel Level = iota // 無効知 0
	Debug
	Info
	Warn
	Error
)

var (
	levelLogrusLevelMap = map[Level]logrus.Level{
		Debug: logrus.DebugLevel,
		Info:  logrus.InfoLevel,
		Warn:  logrus.WarnLevel,
		Error: logrus.ErrorLevel,
	}

	levelValueMap = map[Level]string{
		Debug: "debug",
		Info:  "info",
		Warn:  "warn",
		Error: "error",
	}
)

// LogrusLevel Level 種類に応じた logrus level を返却
func (l Level) LogrusLevel() logrus.Level {
	return levelLogrusLevelMap[l]
}

// applyLevel 入力文字列に応じた Level を返却
func applyLevel(value string) (Level, error) {
	for l, v := range levelValueMap {
		if strings.EqualFold(v, value) {
			return l, nil
		}
	}
	return UnknownLevel, apperr.NewConfigErrF("unknown log level: %s", value)
}

// UnmarshalYAML Yaml ファイルを読み込むためダックタイピング
func (l *Level) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}

	level, applyErr := applyLevel(str)
	if applyErr != nil {
		return applyErr
	}

	*l = level
	return nil
}

// ------------------------------------------------------------
// Format
// ------------------------------------------------------------

// Format ログフォーマット
type Format int

// Format 種類
const (
	UnknownFormat Format = iota
	Console
	JSON
)

var (
	formatValueMap = map[Format]string{
		Console: "console",
		JSON:    "json",
	}

	formatLogrusFormatterMap = map[Format]logrus.Formatter{
		Console: &logrus.TextFormatter{
			ForceColors:    true,
			FullTimestamp:  true,
			DisableSorting: true,
		},
		JSON: &logrus.JSONFormatter{
			DisableHTMLEscape: false,
		},
	}
)

// LogrusFormatter Format 種類に応じた Formatter を返却
func (f Format) LogrusFormatter() logrus.Formatter {
	return formatLogrusFormatterMap[f]
}

// ApplyLevel 入力文字列に応じた Format を返却
func applyFormat(value string) (Format, error) {
	for f, v := range formatValueMap {
		if strings.EqualFold(v, value) {
			return f, nil
		}
	}
	return UnknownFormat, apperr.NewConfigErrF("unknown log format: %s", value)
}

// UnmarshalYAML Yaml ファイルを読み込むためダックタイピング
func (f *Format) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}

	format, applyErr := applyFormat(str)
	if applyErr != nil {
		return applyErr
	}

	*f = format
	return nil
}

// Fields ログ出力情報は関数 LogrusFields で定義
type Fields interface {
	LogrusFields() logrus.Fields
}
