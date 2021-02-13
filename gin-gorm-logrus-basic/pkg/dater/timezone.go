package dater

import "time"

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

// NowJST 現在の JST 時間を返却
func NowJST() time.Time {
	loc := AsiaTokyo.Location()
	return time.Now().In(loc)
}

// ToJST 時間を JST に変換
func ToJST(t time.Time) time.Time {
	loc := AsiaTokyo.Location()
	return t.In(loc)
}

// NowUTC 現在の UTC 時間を返却
func NowUTC() time.Time {
	loc := UTC.Location()
	return time.Now().In(loc)
}

// ToUTC 時間を UTC に変換
func ToUTC(t time.Time) time.Time {
	loc := UTC.Location()
	return t.In(loc)
}

// ParseInJST DateFormat 形式で文字列を JST 時間に変換
func ParseInJST(fmt DateFormat, date string) (time.Time, error) {
	loc := AsiaTokyo.Location()
	return time.ParseInLocation(string(fmt), date, loc)
}

// IsFistDayOfMonth 月初日を判定
func IsFistDayOfMonth(t time.Time) bool {
	return t.Day() == 1
}

// ToMonthFirst 月初日を取得
func ToMonthFirst(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// ToLastMonthFirst 先月初日を取得
func ToLastMonthFirst(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()-1, 1, 0, 0, 0, 0, t.Location())
}

// ToLastMonthEnd 先月末日を取得
func ToLastMonthEnd(t time.Time) time.Time {
	return time.
		Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).
		AddDate(0, 0, -1)
}
