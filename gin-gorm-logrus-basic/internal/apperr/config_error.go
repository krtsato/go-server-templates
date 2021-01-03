package apperr

import "fmt"

// ConfigError サーバ起動時に読み込む設定値エラー
type ConfigError struct {
	Cause error
	Msg   string
}

// Error stringify error
func (e *ConfigError) Error() string {
	if e == nil {
		return ""
	} else if e.Cause == nil {
		return fmt.Sprintf("Msg: %s", e.Msg)
	} else {
		return fmt.Sprintf("Cause: %s, Msg: %s", e.Unwrap().Error(), e.Msg)
	}
}

// Unwrap unwrap error
func (e *ConfigError) Unwrap() error {
	wErr, ok := e.Cause.(AppError)
	if !ok {
		return nil
	}
	return wErr.Unwrap()
}

// NewConfigError ConfigError を err, msg から生成
func NewConfigError(err error, msg string) error {
	if err == nil {
		return &ConfigError{Msg: msg}
	}
	return &ConfigError{Cause: err, Msg: msg}
}

// NewConfigErrorM ConfigError を msg から生成
func NewConfigErrorM(msg string) error {
	return NewConfigError(nil, msg)
}

// NewConfigErrorF ConfigError を formatted msg から生成
func NewConfigErrorF(format string, a ...interface{}) error {
	return NewConfigError(nil, fmt.Sprintf(format, a...))
}

// NewConfigErrorEF ConfigError を err, formatted msg から生成
func NewConfigErrorEF(err error, format string, a ...interface{}) error {
	return NewConfigError(err, fmt.Sprintf(format, a...))
}
