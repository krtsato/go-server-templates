package app

import (
	"fmt"
	"io"
	"os"

	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/apperr"
	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/logger"
	"github.com/sirupsen/logrus"
)

// Logger app logger
type Logger struct {
	*logrus.Entry
}

// Log 呼び出される Logger インスタンス
var Log = defaultLogger()

// defaultLogger nil 回避のため Logger 初期化
func defaultLogger() *Logger {
	l, _ := newLogger(logger.Config{Level: logger.Info, Format: logger.JSON})
	return l
}

// newLogger 引数となる設定値から Logger を生成
func newLogger(cfg logger.Config) (*Logger, error) {
	l := logrus.New()
	l.SetLevel(cfg.Level.LogrusLevel())
	l.SetFormatter(cfg.Format.LogrusFormatter())

	if len(cfg.FullPath) > 0 {
		logfile, openErr := os.OpenFile(cfg.FullPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if openErr != nil {
			if logfile != nil {
				defer logfile.Close()
			}
			return nil, apperr.NewConfigErr(openErr, "failed to open log file")
		}
		l.SetOutput(io.MultiWriter(logfile, os.Stdout))
	} else {
		l.SetOutput(os.Stdout)
	}
	return &Logger{Entry: l.WithField("type", "app")}, nil
}

// ApplyLogger 引数となるデフォルト Logger を上書き
func ApplyLogger(cfg logger.Config) error {
	l, err := newLogger(cfg)
	if err != nil {
		return apperr.NewConfigErr(err, "failed to apply app logger")
	}
	Log = l
	fmt.Println("succeed at app logger setup")
	return nil
}
