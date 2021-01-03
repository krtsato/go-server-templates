package apperr

// AppError アプリ内部で引き回すエラー
type AppError interface {
	Error() string
	Unwrap() error
}
