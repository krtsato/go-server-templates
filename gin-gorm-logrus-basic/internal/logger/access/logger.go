package access

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/logger"
	"github.com/sirupsen/logrus"
)

// Logger access logger
type Logger struct {
	Logger *logrus.Logger
}

// Log 呼び出される Logger インスタンス
var Log = defaultLogger()

// defaultLogger nil 回避のため Logger 初期化
func defaultLogger() *Logger {
	return newLogger(logger.Config{Level: logger.Info, Format: logger.JSON})
}

// newLogger 引数となる設定値から Logger を生成
func newLogger(cfg logger.Config) *Logger {
	l := logrus.New()
	l.SetLevel(cfg.Level.LogrusLevel())
	l.SetFormatter(cfg.Format.LogrusFormatter())
	l.SetOutput(os.Stdout)
	return &Logger{Logger: l}
}

// ApplyLogger 引数となる設定値からデフォルト Logger を上書き
func ApplyLogger(cfg logger.Config) {
	Log = newLogger(cfg)
	fmt.Println("succeed at access logger setup")
}

// Attributes AccessLogRecord 生成時に log として含める
type Attributes struct {
	Request    *http.Request
	Start      time.Time
	End        time.Time
	Body       string
	StatusCode int
}

// LogrusFields log 出力に載せる情報を返却
func (attr *Attributes) LogrusFields() logrus.Fields {
	req := attr.Request
	return logrus.Fields{
		"type":        "access",
		"time":        attr.Start,
		"method":      req.Method,
		"path":        req.URL.Path,
		"query":       req.URL.RawQuery,
		"body":        attr.Body,
		"remote_ip":   req.RemoteAddr,
		"status_code": attr.StatusCode,
		"duration_ms": int64(attr.End.Sub(attr.Start) / time.Millisecond),
		"headers":     req.Header,
		// "request_id":  req.Header.Get(header.RequestID), TODO: 認証用リクエスト ID 追加
	}
}

// Info Attributes を info レベルの logEntry にする
func (l *Logger) Info(attr *Attributes) {
	l.log(logger.Info, attr, nil)
}

// log logEntry を生成
func (l *Logger) log(level logger.Level, attr *Attributes, err error) {
	logEntry := l.Logger.WithFields(attr.LogrusFields())
	if err != nil {
		logEntry = logEntry.WithError(err)
	}
	logEntry.Log(level.LogrusLevel(), "")
}

// NewWithFields fields を追加した上で新規インスタンスを生成
func (l *Logger) NewWithFields(fields logger.Fields) *Logger {
	return &Logger{Logger: l.Logger.WithFields(fields.LogrusFields()).Logger}
}
