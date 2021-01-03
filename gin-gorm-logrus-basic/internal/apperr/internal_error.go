package apperr

import "fmt"

// InternalError 初期化後のアプリ内部エラー
type InternalError struct {
	Cause error
	Msg   string
}

// Error stringify error
func (e *InternalError) Error() string {
	if e == nil {
		return ""
	} else if e.Cause == nil {
		return fmt.Sprintf("Msg: %s", e.Msg)
	} else {
		return fmt.Sprintf("Cause: %s, Msg: %s", e.Unwrap().Error(), e.Msg)
	}
}

// Unwrap unwrap error
func (e *InternalError) Unwrap() error {
	wErr, ok := e.Cause.(AppError)
	if !ok {
		return nil
	}
	return wErr.Unwrap()
}

// NewInternalError InternalError を err, msg から生成
func NewInternalError(err error, msg string) error {
	if err == nil {
		return &InternalError{Msg: msg}
	}
	return &InternalError{Cause: err, Msg: msg}
}

// NewInternalErrorM InternalError を msg から生成
func NewInternalErrorM(msg string) error {
	return NewInternalError(nil, msg)
}

// NewInternalErrorF InternalError を formatted msg から生成
func NewInternalErrorF(format string, a ...interface{}) error {
	return NewInternalError(nil, fmt.Sprintf(format, a...))
}

// NewInternalErrorEF InternalError を err, formatted msg から生成
func NewInternalErrorEF(err error, format string, a ...interface{}) error {
	return NewInternalError(err, fmt.Sprintf(format, a...))
}
