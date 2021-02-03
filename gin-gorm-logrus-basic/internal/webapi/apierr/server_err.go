package apierr

import (
	"fmt"
	"net/http"
)

// InternalServer internal server error 500
// NOTE: クライアントに返却する Internal Server Error は原則 500 番とする
// アプリ内部で取り扱うエラーは internal/apperr で定義する
type InternalServer struct {
	ErrCode ErrCode
	StsCode int
	Cause   error
}

// ErrorCode ErrCode を返却
func (i *InternalServer) ErrorCode() ErrCode {
	if i == nil {
		return UnknownErrCode
	}
	return i.ErrCode
}

// StatusCode StsCode を返却
func (i *InternalServer) StatusCode() int {
	if i == nil {
		return 500
	}
	if i.StsCode == 0 {
		return 500
	}
	return i.StsCode
}

// Error エラー文字列を返却
func (i *InternalServer) Error() string {
	if i == nil {
		return ""
	} else if i.Cause == nil {
		return http.StatusText(i.StatusCode())
	} else {
		return i.Cause.Error()
	}
}

// NewInternalServer InternalErr を err, msg から生成
func NewInternalServer(errCode ErrCode, stsCode int, err error) *InternalServer {
	if err == nil {
		return &InternalServer{ErrCode: UnknownErrCode, StsCode: 500}
	}
	return &InternalServer{ErrCode: errCode, StsCode: stsCode, Cause: err}
}

// NewInternalServerE InternalServer を errCode, err から生成
func NewInternalServerE(errCode ErrCode, err error) *InternalServer {
	return NewInternalServer(errCode, 500, err)
}

// NewInternalServerF InternalServer を errCode, formatted err から生成
func NewInternalServerF(errCode ErrCode, format string, a ...interface{}) *InternalServer {
	return NewInternalServer(errCode, 500, fmt.Errorf(format, a...))
}
