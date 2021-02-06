package apperr

import "fmt"

// ConfigErr サーバ起動時の設定初期化エラー
// クライアント側にエラー内容を隠蔽するとき Msg をセットする
type ConfigErr struct {
	Code  ErrCode
	Cause error
	Msg   string
}

// ErrorCode ErrCode を返却
func (e *ConfigErr) ErrorCode() ErrCode {
	if e == nil {
		return UnknownErrCode
	}
	return e.Code
}

// Error エラー文字列を返却
func (e *ConfigErr) Error() string {
	if e == nil {
		return ""
	}
	if e.Msg == "" {
		return fmt.Sprintf("Error Cause: %s", e.Unwrap().Error())
	}
	return fmt.Sprintf("Error Message: %s", e.Msg)
}

// Unwrap ConfigErr をアンラップ
func (e *ConfigErr) Unwrap() error {
	if e == nil {
		return nil
	}
	if err, ok := e.Cause.(interface{ Unwrap() error }); ok {
		return err.Unwrap()
	}
	return e.Cause
}

// NewConfigErr ConfigErr を err, msg から生成
func NewConfigErr(err error, msg string) error {
	if err == nil {
		return &ConfigErr{Code: UnknownErrCode, Msg: msg}
	}
	return &ConfigErr{Code: ConfigErrCode, Cause: err, Msg: msg}
}

// NewConfigErrM ConfigErr を msg から生成
func NewConfigErrM(msg string) error {
	return NewConfigErr(nil, msg)
}

// NewConfigErrF ConfigErr を formatted msg から生成
func NewConfigErrF(format string, a ...interface{}) error {
	return NewConfigErr(nil, fmt.Sprintf(format, a...))
}

// NewConfigErrEF ConfigErr を err, formatted msg から生成
func NewConfigErrEF(err error, format string, a ...interface{}) error {
	return NewConfigErr(err, fmt.Sprintf(format, a...))
}
