package apperr

// APPErr アプリ内部で取り扱うエラー
type APPErr interface {
	ErrorCode() ErrCode
	Error() string
	Unwrap() error
}

// ErrCode APPErr Enum
type ErrCode int

// ErrCode 種類
const (
	UnknownErrCode ErrCode = iota // 無効値
	ConfigErrCode
	InternalErrCode
)

var errCodeValueMap = map[ErrCode]string{
	ConfigErrCode:   "configErrCode",
	InternalErrCode: "internalErrCode",
}

// String ErrCode 文字列を返却
func (c ErrCode) String() string {
	return errCodeValueMap[c]
}
