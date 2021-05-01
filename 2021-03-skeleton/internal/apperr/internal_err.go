package apperr

import "fmt"

// InternalErr 初期化後のアプリ内部エラー
// クライアント側にエラー内容を隠蔽するとき Msg をセットする
type InternalErr struct {
	Code  ErrCode
	Cause error
	Msg   string
}

// ErrorCode ErrCode を返却
func (e *InternalErr) ErrorCode() ErrCode {
	if e == nil {
		return UnknownErrCode
	}
	return e.Code
}

// Error エラー文字列を返却
func (e *InternalErr) Error() string {
	if e == nil {
		return ""
	}
	if e.Msg == "" {
		return fmt.Sprintf("Error Cause: %s", e.Unwrap().Error())
	}
	return fmt.Sprintf("Error Message: %s", e.Msg)
}

// Unwrap InternalErr をアンラップ
func (e *InternalErr) Unwrap() error {
	if e == nil {
		return nil
	}
	if err, ok := e.Cause.(interface{ Unwrap() error }); ok {
		return err.Unwrap()
	}
	return e.Cause
}

// NewInternalErr InternalErr を err, msg から生成
func NewInternalErr(err error, msg string) error {
	if err == nil {
		return &InternalErr{Code: UnknownErrCode, Msg: msg}
	}
	return &InternalErr{Code: InternalErrCode, Cause: err, Msg: msg}
}

// NewInternalErrM InternalErr を msg から生成
func NewInternalErrM(msg string) error {
	return NewInternalErr(nil, msg)
}

// NewInternalErrF InternalErr を formatted msg から生成
func NewInternalErrF(format string, a ...interface{}) error {
	return NewInternalErr(nil, fmt.Sprintf(format, a...))
}

// NewInternalErrEF InternalErr を err, formatted msg から生成
func NewInternalErrEF(err error, format string, a ...interface{}) error {
	return NewInternalErr(err, fmt.Sprintf(format, a...))
}
