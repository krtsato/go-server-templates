package dater

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// LocalDate DB 接続時の timezone 設定による変換を回避するため使用
// Go の time 型は timezone 情報を持っている
type LocalDate struct {
	Year  uint
	Month uint
	Day   uint
}

// ------------------------------------------------------------
// LocalDate
// ------------------------------------------------------------

// Value LocalDate に応じた日付 YYYY-MM-DD を返却
// go-sql-driver で使用するためダックタイピング
func (d LocalDate) Value() (driver.Value, error) {
	year, month, day := d.SplitString()
	return year + "-" + month + "-" + day, nil
}

// Scan DB レコードをメモリ上にスキャン
// go-sql-driver で使用するためダックタイピング
func (d *LocalDate) Scan(value interface{}) error {
	if d == nil {
		return fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	if value == nil {
		return fmt.Errorf("failed to Scan the empty interface argument")
	}
	convVal, convErr := driver.String.ConvertValue(value)
	if convErr != nil {
		return fmt.Errorf("failed to convert LocalDate: %s", convErr.Error())
	}
	val, ok := convVal.(string)
	if !ok {
		return fmt.Errorf("failed to assert LocalDate type")
	}
	matchVals, matchErr := groupSubMatch(val, LocalDateRegex)
	if matchErr != nil {
		return fmt.Errorf("failed to match LocalDate: %s", matchErr.Error())
	}
	if len(matchVals) < 4 {
		return fmt.Errorf("failed to match LocalDate in the group len: %d", len(matchVals))
	}
	year, yErr := strconv.Atoi(matchVals[1])
	month, mErr := strconv.Atoi(matchVals[2])
	day, dErr := strconv.Atoi(matchVals[3])
	if yErr != nil || mErr != nil || dErr != nil {
		return fmt.Errorf("failed to convert LocalDate matchVals [ %s, %s, %s ]", matchVals[1], matchVals[2], matchVals[3])
	}
	*d = LocalDate{Year: uint(year), Month: uint(month), Day: uint(day)}
	return nil
}

// Valid 有効期間内の LocalDate を返却
func (d LocalDate) Valid() (LocalDate, error) {
	if d.Year < MinYear || MaxYear < d.Year {
		return LocalDate{}, fmt.Errorf("%d is out of LocalDate year range", d.Year)
	}
	if d.Month < MinMonthOfYear || MaxMonthOfYear < d.Month {
		return LocalDate{}, fmt.Errorf("%d is out of LocalDate month range", d.Month)
	}
	if d.Day < MinDayOfMonth || MaxDayOfMonth < d.Day {
		return LocalDate{}, fmt.Errorf("%d is out of LocalDate day range", d.Day)
	}
	return d, nil
}

// groupSubMatch 正規表現にマッチする文字列グループを返却
func groupSubMatch(target, regex string) ([]string, error) {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return make([]string, 0), err
	}
	return reg.FindStringSubmatch(target), nil
}

// String LocalDate に応じた文字列を返却
func (d LocalDate) String() string {
	val, _ := d.Value()
	return val.(string)
}

// SplitString Year, Month, Day の桁数を揃えて返却
// ex) 2021, 1, 1 -> "2021", "01", "01"
func (d LocalDate) SplitString() (yearStr, monthStr, dayStr string) {
	var year, month, day string
	switch {
	case d.Year < 10:
		year = "000" + strconv.Itoa(int(d.Year))
	case d.Year < 100:
		year = "00" + strconv.Itoa(int(d.Year))
	case d.Year < 1000:
		year = "0" + strconv.Itoa(int(d.Year))
	default:
		year = strconv.Itoa(int(d.Year))
	}
	if d.Month < 10 {
		month = "0" + strconv.Itoa(int(d.Month))
	} else {
		month = strconv.Itoa(int(d.Month))
	}
	if d.Day < 10 {
		day = "0" + strconv.Itoa(int(d.Day))
	} else {
		day = strconv.Itoa(int(d.Day))
	}
	return year, month, day
}

// ToTime Location に応じた日付を返却
func (d LocalDate) ToTime(loc *time.Location) time.Time {
	return time.Date(int(d.Year), time.Month(int(d.Month)), int(d.Day), 0, 0, 0, 0, loc)
}

// ToTimeUTC UTC 時間を返却
func (d LocalDate) ToTimeUTC() time.Time {
	loc := UTC.Location()
	return d.ToTime(loc)
}

// IsBefore LocalDate が引数よりも遅れた日付のとき true を返却
func (d LocalDate) IsBefore(targetDate LocalDate) bool {
	firstDate := d.ToTimeUTC()
	secondDate := targetDate.ToTimeUTC()
	return firstDate.Before(secondDate)
}

// IsAfter LocalDate が引数よりも進んだ日付のとき true を返却
func (d LocalDate) IsAfter(targetDate LocalDate) bool {
	firstDate := d.ToTimeUTC()
	secondDate := targetDate.ToTimeUTC()
	return firstDate.After(secondDate)
}

// IsEqual LocalDate が引数と同じ日付のとき true を返却
func (d LocalDate) IsEqual(targetDate LocalDate) bool {
	return d.ToTimeUTC().Equal(targetDate.ToTimeUTC())
}

// IsBetween LocalDate が引数の範囲内にあるとき true を返却
func (d LocalDate) IsBetween(start, end LocalDate) bool {
	return (d.IsAfter(start) || d.IsEqual(start)) && (d.IsEqual(end) || d.IsBefore(end))
}

// IsZero LocalDate がゼロ値のとき true を返却
func (d LocalDate) IsZero() bool {
	return d.Year == 0 && d.Month == 0 && d.Day == 0
}

// LocalDatetime LocalDate とゼロ値の LocalTime から LocalDateTime を返却
func (d LocalDate) LocalDatetime() LocalDatetime {
	return LocalDatetime{
		LocalDate: d,
		LocalTime: LocalTime{Hour: 0, Minute: 0, Second: 0},
	}
}

// Sub LocalDate から引数を引いた期間を返却
// 290 年以上の期間は扱わない
func (d LocalDate) Sub(targetDate LocalDate) (time.Duration, bool) {
	duration := d.ToTimeUTC().Sub(targetDate.ToTimeUTC())
	if duration == MaxDuration || duration == MinDuration {
		return duration, false
	}
	return duration, true
}

// MarshalJSON JSON に YYYY-MM-DD を書き込む
// encoding/json で使用するためダックタイピング
func (d LocalDate) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(d.String())
}

// UnmarshalJSON JSON の YYYY-MM-DD を UTC 時間として読み取る
// encoding/json で使用するためダックタイピング
func (d *LocalDate) UnmarshalJSON(data []byte) error {
	if d == nil {
		return fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	if len(data) == 0 {
		return fmt.Errorf("failed to UnmarshalJSON LocalDate because of zero length data")
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("failed to UnmarshalJSON LocalDate: %s", err.Error())
	}
	timeUTC, err := time.ParseInLocation(DateHyphen.String(), str, UTC.Location())
	if err != nil {
		return fmt.Errorf("failed to parse LocalDate: %s", err.Error())
	}
	*d = applyLocalDate(timeUTC)
	return nil
}

// applyLocalDate Time 型から LocalDate を生成
func applyLocalDate(t time.Time) LocalDate {
	return LocalDate{
		Year:  uint(t.Year()),
		Month: uint(t.Month()),
		Day:   uint(t.Day()),
	}
}

// NewLocalDate Year, month, Day から UTC 時間の LocalDate を生成
func NewLocalDate(year, month, day int) LocalDate {
	timeUTC := time.Date(year, time.Month(month), day, 0, 0, 0, 0, UTC.Location())
	if timeUTC.Unix() < FirstUnixInAD {
		return LocalDate{}
	}
	return LocalDate{Year: uint(timeUTC.Year()), Month: uint(timeUTC.Month()), Day: uint(timeUTC.Day())}
}

// ------------------------------------------------------------
// NullLocalDate
// ------------------------------------------------------------

// NullLocalDate nullable LocalDate
type NullLocalDate struct {
	LocalDate LocalDate
	Valid     bool
}

// Value LocalDate に応じた日付 YYYY-MM-DD を返却
// go-sql-driver で使用するためダックタイピング
func (nd NullLocalDate) Value() (driver.Value, error) {
	if nd.Valid {
		return nd.LocalDate.Value()
	}
	return nil, nil
}

// Scan DB レコードをメモリ上にスキャン
// go-sql-driver で使用するためダックタイピング
func (nd *NullLocalDate) Scan(value interface{}) error {
	if nd == nil || value == nil {
		nd.LocalDate, nd.Valid = LocalDate{}, false
		return nil
	}
	err := nd.LocalDate.Scan(value)
	if err != nil {
		nd.Valid = false
		return fmt.Errorf("failed to Scan NullLocalDate: %s", err.Error())
	}
	nd.Valid = true
	return nil
}

// MarshalJSON JSON に YYYY-MM-DD を書き込む
// encoding/json で使用するためダックタイピング
func (nd NullLocalDate) MarshalJSON() ([]byte, error) {
	if nd.Valid {
		return nd.LocalDate.MarshalJSON()
	}
	return json.Marshal(nil)
}

// UnmarshalJSON JSON の YYYY-MM-DD を UTC 時間として読み取る
// encoding/json で使用するためダックタイピング
func (nd *NullLocalDate) UnmarshalJSON(data []byte) error {
	if nd == nil {
		return fmt.Errorf("nil receiver of NullLocalDate is invalid")
	}
	if len(data) == 0 || strings.EqualFold(string(data), "null") {
		nd.LocalDate, nd.Valid = LocalDate{}, false
		return nil
	}
	if err := nd.LocalDate.UnmarshalJSON(data); err != nil {
		nd.LocalDate, nd.Valid = LocalDate{}, false
		return err
	}
	nd.Valid = true
	return nil
}

// NewNullLocalDate Year, Month, Day から NullLocalDate を生成
func NewNullLocalDate(year, month, day int) NullLocalDate {
	localDate := NewLocalDate(year, month, day)
	if localDate.IsZero() {
		return NullLocalDate{Valid: false}
	}
	return NullLocalDate{LocalDate: localDate, Valid: true}
}
