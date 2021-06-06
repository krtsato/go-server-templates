//go:generate mockgen -destination=$PRJ_ROOT/pkg/mock/$GOPACKAGE/$GOFILE -package=$GOPACKAGE -source=$GOFILE

package apperr

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

type (
	// AppErr is error inside app.
	AppErr interface {
		ErrorCode() Code
		Error() string
		Unwrap() error
	}

	appErr struct {
		code  Code
		cause error
	}
)

// ErrorCode returns error Code enum.
func (a appErr) ErrorCode() Code {
	if a.cause == nil {
		return OK
	}
	return a.code
}

// Error returns Code and origin error.
func (a appErr) Error() string {
	return fmt.Sprintf("code: %s, cause: %s", a.ErrorCode(), a.Unwrap().Error())
}

// Unwrap returns original error.
func (a appErr) Unwrap() error {
	if err, ok := a.cause.(interface{ Unwrap() error }); ok {
		return err.Unwrap()
	}
	return a.cause
}

// Error generates app error with format string.
func Error(c Code, str string) error {
	if c == OK {
		return nil
	}
	return appErr{
		code:  c,
		cause: xerrors.New(str),
	}
}

// Errorf generates app error with format string.
func Errorf(c Code, format string, v ...interface{}) error {
	if c == OK {
		return nil
	}
	return appErr{
		code:  c,
		cause: xerrors.Errorf(format, v...),
	}
}

// GetCode takes error and returns Code accordingly.
func GetCode(err error) Code {
	if err == nil {
		return OK
	}
	var e appErr
	if errors.As(err, &e) {
		return e.code
	}
	return Unknown
}
