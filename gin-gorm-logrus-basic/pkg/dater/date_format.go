package dater

import "time"

// DateFormat 日付フォーマット
type DateFormat string

// String DateFormat に応じた文字列を返却
func (f DateFormat) String() string {
	return string(f)
}

// DateFormat 種類
const (
	Month           = DateFormat("200601")
	MonthHyphen     = DateFormat("2006-01")
	DateAbbreviated = DateFormat("060102")
	Date            = DateFormat("20060102")
	DateHyphen      = DateFormat("2006-01-02")
	DateSlash       = DateFormat("2006/01/02")
	DateHour        = DateFormat("2006010215")
	DateHourHyphen  = DateFormat("2006-01-02 15")
	DateHourSlash   = DateFormat("2006/01/02 15")
	DateRFC3339     = DateFormat("2006-01-02 Z07:00")
	DateTime        = DateFormat("20060102150405")
	DateTimeHyphen  = DateFormat("2006-01-02 15:04:05")
	DateTimeSlash   = DateFormat("2006/01/02 15:04:05")
	ANSIC           = DateFormat(time.ANSIC)
	UnixDate        = DateFormat(time.UnixDate)
	RubyDate        = DateFormat(time.RubyDate)
	RFC822          = DateFormat(time.RFC822)
	RFC822Z         = DateFormat(time.RFC822Z)
	RFC850          = DateFormat(time.RFC850)
	RFC1123         = DateFormat(time.RFC1123)
	RFC1123Z        = DateFormat(time.RFC1123Z)
	RFC3339         = DateFormat(time.RFC3339)
	RFC3339Nano     = DateFormat(time.RFC3339Nano)
)
