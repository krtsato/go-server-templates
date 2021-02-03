package apierr

import (
	"errors"
	"fmt"
)

// ------------------------------------------------------------
// BadRequest
// ------------------------------------------------------------

// BadRequest 400
type BadRequest struct {
	ErrCode ErrCode
	Cause   error
}

// ErrorCode ErrCode を返却
func (b *BadRequest) ErrorCode() ErrCode {
	return b.ErrCode
}

// StatusCode 400 を返却
func (*BadRequest) StatusCode() int {
	return 400
}

// Error エラー文字列を返却
func (b *BadRequest) Error() string {
	return b.Cause.Error()
}

// NewBadRequest BadRequest を err から生成
func NewBadRequest(errCode ErrCode, err error) *BadRequest {
	return &BadRequest{ErrCode: errCode, Cause: err}
}

// NewBadRequestM BadRequest を msg から生成
func NewBadRequestM(errCode ErrCode, msg string) *BadRequest {
	return NewBadRequest(errCode, errors.New(msg))
}

// NewBadRequestF BadRequest を formatted err から生成
func NewBadRequestF(errCode ErrCode, format string, a ...interface{}) *BadRequest {
	return NewBadRequest(errCode, fmt.Errorf(format, a...))
}

// ------------------------------------------------------------
// Unauthorized
// ------------------------------------------------------------

// Unauthorized 401
type Unauthorized struct {
	ErrCode ErrCode
	Cause   error
}

// ErrorCode ErrCode を返却
func (u *Unauthorized) ErrorCode() ErrCode {
	return u.ErrCode
}

// StatusCode 401 を返却
func (*Unauthorized) StatusCode() int {
	return 401
}

// Error エラー文字列を返却
func (u *Unauthorized) Error() string {
	return u.Cause.Error()
}

// NewUnauthorized Unauthorized を err から生成
func NewUnauthorized(errCode ErrCode, err error) *Unauthorized {
	return &Unauthorized{ErrCode: errCode, Cause: err}
}

// NewUnauthorizedM Unauthorized を msg から生成
func NewUnauthorizedM(errCode ErrCode, msg string) *BadRequest {
	return NewBadRequest(errCode, errors.New(msg))
}

// NewUnauthorizedF Unauthorized を formatted err から生成
func NewUnauthorizedF(errCode ErrCode, format string, a ...interface{}) *BadRequest {
	return NewBadRequest(errCode, fmt.Errorf(format, a...))
}

// ------------------------------------------------------------
// Forbidden
// ------------------------------------------------------------

// Forbidden 403
type Forbidden struct {
	ErrCode ErrCode
	Cause   error
}

// ErrorCode ErrCode を返却
func (f *Forbidden) ErrorCode() ErrCode {
	return f.ErrCode
}

// StatusCode 403 を返却
func (*Forbidden) StatusCode() int {
	return 403
}

// Error エラー文字列を返却
func (f *Forbidden) Error() string {
	return f.Cause.Error()
}

// NewForbidden Forbidden を err から生成
func NewForbidden(errCode ErrCode, err error) *Forbidden {
	return &Forbidden{ErrCode: errCode, Cause: err}
}

// NewForbiddenM Forbidden を msg から生成
func NewForbiddenM(errCode ErrCode, msg string) *Forbidden {
	return NewForbidden(errCode, errors.New(msg))
}

// NewForbiddenF Forbidden を formatted err から生成
func NewForbiddenF(errCode ErrCode, format string, a ...interface{}) *Forbidden {
	return NewForbidden(errCode, fmt.Errorf(format, a...))
}

// ------------------------------------------------------------
// NotFound
// ------------------------------------------------------------

// NotFound 404
type NotFound struct {
	ErrCode ErrCode
	Cause   error
}

// ErrCode ErrCode を返却
func (n *NotFound) ErrorCode() ErrCode {
	return n.ErrCode
}

// StatusCode 404 を返却
func (*NotFound) StatusCode() int {
	return 404
}

// Error エラー文字列を返却
func (n *NotFound) Error() string {
	return n.Cause.Error()
}

// NewNotFound NotFound を err から生成
func NewNotFound(errCode ErrCode, err error) *NotFound {
	return &NotFound{ErrCode: errCode, Cause: err}
}

// NewNotFoundM NotFound を msg から生成
func NewNotFoundM(errCode ErrCode, msg string) *NotFound {
	return NewNotFound(errCode, errors.New(msg))
}

// NewNotFoundF NotFound を formatted err から生成
func NewNotFoundF(errCode ErrCode, format string, a ...interface{}) *NotFound {
	return NewNotFound(errCode, fmt.Errorf(format, a...))
}
