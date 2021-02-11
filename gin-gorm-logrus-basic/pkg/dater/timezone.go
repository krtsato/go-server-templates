package dater

import "time"

// Format ~ エイリアス
type Format string

// String Format 文字列を返却
func (f Format) String() string {
	return string(f)
}

// Format 種類
const (
	Month           = Format("200601")
	MonthHyphen     = Format("2006-01")
	DateAbbreviated = Format("060102")
	Date            = Format("20060102")
	DateHyphen      = Format("2006-01-02")
	DateSlash       = Format("2006/01/02")
	DateHour        = Format("2006010215")
	DateHourHyphen  = Format("2006-01-02 15")
	DateHourSlash   = Format("2006/01/02 15")
	DateRFC3339     = Format("2006-01-02 Z07:00")
	DateTime        = Format("20060102150405")
	DateTimeHyphen  = Format("2006-01-02 15:04:05")
	DateTimeSlash   = Format("2006/01/02 15:04:05")
	ANSIC           = Format(time.ANSIC)
	UnixDate        = Format(time.UnixDate)
	RubyDate        = Format(time.RubyDate)
	RFC822          = Format(time.RFC822)
	RFC822Z         = Format(time.RFC822Z)
	RFC850          = Format(time.RFC850)
	RFC1123         = Format(time.RFC1123)
	RFC1123Z        = Format(time.RFC1123Z)
	RFC3339         = Format(time.RFC3339)
	RFC3339Nano     = Format(time.RFC3339Nano)
)

// Timezone timezone Enum
type Timezone int

// timezone 種類
const (
	UTC Timezone = iota
	AsiaTokyo
)

var timezoneValueMap = map[Timezone]string{
	UTC:       "UTC",
	AsiaTokyo: "Asia/Tokyo",
}
var timezoneOffsetMap = map[Timezone]int{
	UTC:       0 * 60 * 60,
	AsiaTokyo: +9 * 60 * 60,
}

// String Timezone に応じた文字列を返却
func (t Timezone) String() string {
	return timezoneValueMap[t]
}

// Offset Timezone に応じたオフセットを返却
func (t Timezone) Offset() int {
	return timezoneOffsetMap[t]
}

// Location Timezone に応じたロケーションを返却
func (t Timezone) Location() *time.Location {
	return time.FixedZone(t.String(), t.Offset())
}

// NowJST now JST time
func NowJST() time.Time {
	loc := AsiaTokyo.Location()
	return time.Now().In(loc)
}

// NowUTC now UTC time
func NowUTC() time.Time {
	loc := UTC.Location()
	return time.Now().In(loc)
}

// ToUTC UTC に変換
func ToUTC(t time.Time) time.Time {
	loc := UTC.Location()
	return t.In(loc)
}

// ToJST JST に変換
func ToJST(t time.Time) time.Time {
	loc := AsiaTokyo.Location()
	return t.In(loc)
}

// ParseInJST string to jst time
func ParseInJST(fmt Format, t string) (time.Time, error) {
	loc := AsiaTokyo.Location()
	return time.ParseInLocation(string(fmt), t, loc)
}

// IsFistDayOfMonth 対象日が月初かどうか判定します。
func IsFistDayOfMonth(t time.Time) bool {
	return t.Day() == 1
}

// ToMonthFirst 指定日付の月初を取得します。
func ToMonthFirst(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// ToLastMonthFirst 指定日付の先月月初を取得します。
func ToLastMonthFirst(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()-1, 1, 0, 0, 0, 0, t.Location())
}

// ToLastMonthEnd 指定日付の先月月末を取得します。
func ToLastMonthEnd(t time.Time) time.Time {
	return time.
		Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).
		AddDate(0, 0, -1)
}
