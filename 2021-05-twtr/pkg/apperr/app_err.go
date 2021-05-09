package apperr

import (
	"fmt"
)

type (
	// AppErr is error inside app.
	AppErr interface {
		ErrorCode() Code
		Error() string
		Unwrap() error
	}

	appErrImpl struct {
		code  Code
		cause error
	}
)

// ErrorCode returns error Code enum.
func (a appErrImpl) ErrorCode() Code {
	if a.cause == nil {
		return OK
	}
	return a.code
}

// Error returns Code and origin error.
func (a appErrImpl) Error() string {
	return fmt.Sprintf("code: %s, cause: %s", a.ErrorCode(), a.Unwrap().Error())
}

// Unwrap returns original error.
func (a appErrImpl) Unwrap() error {
	if err, ok := a.cause.(interface{ Unwrap() error }); ok {
		return err.Unwrap()
	}
	return a.cause
}

// ErrorF generates app error with format string.
func ErrorF(c Code, format string, v ...interface{}) error {
	if c == OK {
		return nil
	}
	return appErrImpl{
		code:  c,
		cause: fmt.Errorf(format, v...),
	}
}
