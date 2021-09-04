package apperr

// Code is string error code.
type Code string

// Code is constant.
const (
	Unknown   Code = "unknown"
	OK        Code = "ok"
	Config    Code = "config"
	Unmarshal Code = "unmarshal"
)
