package apierr

// APIErr Web API で取り扱うエラー
type APIErr interface {
	ErrorCode() ErrCode
	StatusCode() int
	Error() string
}

// ErrCode APIErr Enum
type ErrCode int

// ErrCode 種類
const (
	UnknownErrCode ErrCode = iota // 無効値
	PublicErrCode
	PrivateErrCode
)

var errCodeValueMap = map[ErrCode]string{
	PublicErrCode:  "publicErrCode",
	PrivateErrCode: "privateErrCode",
}

// String ErrCode 文字列を返却
func (c ErrCode) String() string {
	return errCodeValueMap[c]
}
