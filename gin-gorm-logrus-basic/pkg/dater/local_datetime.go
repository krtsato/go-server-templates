package dater

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

// const values used by LocalDate, LocalTime and LocalDatetime
const (
	LocalDateRegex                   = "^(\\d{4})-(\\d{1,2})-(\\d{1,2})"
	LocalDateTimeRegex               = "^(\\d{4})-(\\d{1,2})-(\\d{1,2})\\ (\\d{1,2}):(\\d{1,2}):(\\d{1,2})"
	MaxYear            uint          = 999999999 // mysqlの最大値とは異なるため注意
	MinYear            uint          = 0
	MaxMonthOfYear     uint          = 12
	MinMonthOfYear     uint          = 1
	MaxDayOfMonth      uint          = 31
	MinDayOfMonth      uint          = 1
	MaxHourOfDay       uint          = 23
	MinHourOfDay       uint          = 1
	MaxMinuteOfHour    uint          = 59
	MinMinuteOfHour    uint          = 0
	MaxSecOfMinute     uint          = 59
	MinSecOfMinute     uint          = 0
	MinDuration        time.Duration = -1 << 63
	MaxDuration        time.Duration = 1<<63 - 1
	FirstUnixInAD      int64         = -62135596800
)

// ------------------------------------------------------------
// LocalDatetime
// ------------------------------------------------------------

// LocalDatetime local datetime
type LocalDatetime struct {
	LocalDate LocalDate
	LocalTime LocalTime
}

// Value LocalDatetime に応じた日付時刻 YYYY-MM-DD HH:MM:SS を返却
// go-sql-driver で使用するためダックタイピング
func (d LocalDatetime) Value() (driver.Value, error) {
	year, mon, day := d.LocalDate.SplitString()
	hour, min, sec := d.LocalTime.SplitString()
	return fmt.Sprintf("%s-%s-%s %s:%s:%s", year, mon, day, hour, min, sec), nil
}

// Scan DB レコードをメモリ上にスキャン
// go-sql-driver で使用するためダックタイピング
func (d *LocalDatetime) Scan(value interface{}) error {
	if d == nil {
		return errors.New("nil receiver of LocalDatetime is invalid")
	}
	if value == nil {
		return errors.New("failed to Scan the empty interface argument")
	}
	convVal, convErr := driver.String.ConvertValue(value)
	if convErr != nil {
		return fmt.Errorf("failed to convert LocalDatetime: %s", convErr.Error())
	}
	val, ok := convVal.(string)
	if !ok {
		return errors.New("failed to assert LocalDatetime type")
	}
	matchVals, matchErr := groupSubMatch(val, LocalDateTimeRegex)
	if matchErr != nil {
		return fmt.Errorf("failed to match LocalDate: %s", matchErr.Error())
	}
	if len(matchVals) < 7 {
		return fmt.Errorf("failed to match LocalDatetime in group len: %d", len(matchVals))
	}
	year, yErr := strconv.Atoi(matchVals[1])
	mon, monErr := strconv.Atoi(matchVals[2])
	day, dErr := strconv.Atoi(matchVals[3])
	hour, hErr := strconv.Atoi(matchVals[4])
	min, minErr := strconv.Atoi(matchVals[5])
	sec, sErr := strconv.Atoi(matchVals[6])
	if yErr != nil || monErr != nil || dErr != nil || hErr != nil || minErr != nil || sErr != nil {
		return fmt.Errorf("failed to convert LocalDatetime matchVals [ %s, %s, %s, %s, %s, %s ]", matchVals[1], matchVals[2], matchVals[3], matchVals[4], matchVals[5], matchVals[6])
	}
	*d = LocalDatetime{
		LocalDate: LocalDate{Year: uint(year), Month: uint(mon), Day: uint(day)},
		LocalTime: LocalTime{Hour: uint(hour), Minute: uint(min), Second: uint(sec)},
	}
	return nil
}

// String LocalDatetime に応じた文字列を返却
func (d LocalDatetime) String() string {
	val, _ := d.Value()
	return val.(string)
}

// MarshalJSON JSON に YYYY-MM-DD HH:MM:SS を書き込む
// encoding/json で使用するためダックタイピング
func (d LocalDatetime) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(d.String())
}

// UnmarshalJSON JSON の YYYY-MM-DD HH:MM:SS を UTC 時間として読み取る
// encoding/json で使用するためダックタイピング
func (d *LocalDatetime) UnmarshalJSON(data []byte) error {
	if d == nil {
		return errors.New("nil receiver of LocalDatetime is invalid")
	}
	if len(data) == 0 {
		return errors.New("failed to UnmarshalJSON LocalDatetime because of zero length data")
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("failed to UnmarshalJSON LocalDatetime: %s", err.Error())
	}
	timeUTC, err := time.ParseInLocation(DateTimeHyphen.String(), str, UTC.Location())
	if err != nil {
		return fmt.Errorf("failed to parse LocalDatetime: %s", err.Error())
	}
	*d = ApplyLocalDatetime(timeUTC)
	return nil
}

// ToTime Location に応じた日時を返却
func (d LocalDatetime) ToTime(loc *time.Location) time.Time {
	return time.Date(
		int(d.LocalDate.Year), time.Month(int(d.LocalDate.Month)), int(d.LocalDate.Day),
		int(d.LocalTime.Hour), int(d.LocalTime.Minute), int(d.LocalTime.Second), 0, loc)
}

// ToTimeUTC UTC 時間を返却
func (d LocalDatetime) ToTimeUTC() time.Time {
	loc := UTC.Location()
	return d.ToTime(loc)
}

// IsBefore LocalDatetime が引数よりも遅れた日付のとき true を返却
func (d LocalDatetime) IsBefore(targetDateTime LocalDatetime) bool {
	firstDateTime := d.ToTimeUTC()
	secondDateTime := targetDateTime.ToTimeUTC()
	return firstDateTime.Before(secondDateTime)
}

// IsAfter LocalDate が引数よりも進んだ日付のとき true を返却
func (d LocalDatetime) IsAfter(targetDateTime LocalDatetime) bool {
	firstDateTime := d.ToTimeUTC()
	secondDateTime := targetDateTime.ToTimeUTC()
	return firstDateTime.After(secondDateTime)
}

// IsBeforeEqual LocalDatetime が引数と同日または遅れた日付のとき true を返却
func (d LocalDatetime) IsBeforeEqual(targetDateTime LocalDatetime) bool {
	return d.IsBefore(targetDateTime) || d.IsEqual(targetDateTime)
}

// IsAfterEqual LocalDatetime が引数と同日または進んだ日付のとき true を返却
func (d LocalDatetime) IsAfterEqual(targetDateTime LocalDatetime) bool {
	return d.IsAfter(targetDateTime) || d.IsEqual(targetDateTime)
}

// IsEqual LocalDatetime が引数と同じ日付のとき true を返却
func (d LocalDatetime) IsEqual(targetDateTime LocalDatetime) bool {
	return d.ToTimeUTC().Equal(targetDateTime.ToTimeUTC())
}

// IsBetween LocalDatetime が引数の範囲内にあるとき true を返却
func (d LocalDatetime) IsBetween(start, end LocalDatetime) bool {
	return (d.IsAfter(start) || d.IsEqual(start)) && (d.IsEqual(end) || d.IsBefore(end))
}

// Sub LocalDatetime から引数を引いた期間を返却
// 290 年以上の期間は扱わない
func (d LocalDatetime) Sub(target LocalDatetime) (time.Duration, bool) {
	duration := d.ToTimeUTC().Sub(target.ToTimeUTC())
	if duration == MaxDuration || duration == MinDuration {
		return duration, false
	}
	return duration, true
}

// Add LocalDatetime を引数の Duration だけ進める
func (d LocalDatetime) Add(duration time.Duration) LocalDatetime {
	loc := UTC.Location()
	addedTime := d.ToTime(loc).Add(duration)
	return ApplyLocalDatetime(addedTime)
}

// AddDate LocalDatetime を引数の Year, Month, Day だけ進める
func (d LocalDatetime) AddDate(year, month, day int) LocalDatetime {
	loc := UTC.Location()
	addedTime := d.ToTime(loc).AddDate(year, month, day)
	return ApplyLocalDatetime(addedTime)
}

// IsZero LocalDatetime がゼロ値のとき true を返却
func (d LocalDatetime) IsZero() bool {
	return d.LocalDate.IsZero() && d.LocalTime.IsZero()
}

// NewLocalDatetime Year, Month, Day, Hour, Minute, Second から LocalDatetime を生成
// 引数の日時が最大値を超過する場合 time.Date() によって標準化
// 引数の日時が紀元前になる場合 LocalDatetime のゼロ値を返却
func NewLocalDatetime(year, month, day uint, hour, min, sec int, loc *time.Location) LocalDatetime {
	t := time.Date(int(year), time.Month(month), int(day), hour, min, sec, 0, loc)
	if t.Unix() < FirstUnixInAD {
		return LocalDatetime{}
	}
	localDate := LocalDate{Year: uint(t.Year()), Month: uint(t.Month()), Day: uint(t.Day())}
	localTime := LocalTime{Hour: uint(t.Hour()), Minute: uint(t.Minute()), Second: uint(t.Second())}
	return LocalDatetime{LocalDate: localDate, LocalTime: localTime}
}

// NowLocalDatetimeJST JST の現在時刻から LocalDatetime を生成
func NowLocalDatetimeJST() LocalDatetime {
	return ApplyLocalDatetime(NowJST())
}

// NowLocalDatetimeUTC UTC の現在時刻から LocalDatetime を生成
func NowLocalDatetimeUTC() LocalDatetime {
	return ApplyLocalDatetime(NowUTC())
}

// ParseLocalDatetime DateFormat 形式 Location 時間に文字列を変換
func ParseLocalDatetime(fmt DateFormat, dateStr string, loc *time.Location) (LocalDatetime, error) {
	t, err := time.ParseInLocation(fmt.String(), dateStr, loc)
	if err != nil {
		return LocalDatetime{}, err
	}
	return ApplyLocalDatetime(t), nil
}

// ApplyLocalDatetime Time から LocalDatetime を生成
func ApplyLocalDatetime(tm time.Time) LocalDatetime {
	return NewLocalDatetime(uint(tm.Year()), uint(tm.Month()), uint(tm.Day()), tm.Hour(), tm.Minute(), tm.Second(), tm.Location())
}

// ------------------------------------------------------------
// NullLocalDatetime
// ------------------------------------------------------------

// NullLocalDatetime nullable localDatetime
type NullLocalDatetime struct {
	LocalDatetime LocalDatetime
	Valid         bool
}

// Value LocalDatetime に応じた日付 YYYY-MM-DD HH:MM:SS を返却
// go-sql-driver で使用するためダックタイピング
func (nd NullLocalDatetime) Value() (driver.Value, error) {
	if nd.Valid {
		return nd.LocalDatetime.Value()
	}
	return nil, nil
}

// Scan DB レコードをメモリ上にスキャン
// go-sql-driver で使用するためダックタイピング
func (nd *NullLocalDatetime) Scan(value interface{}) error {
	if nd == nil || value == nil {
		nd.LocalDatetime, nd.Valid = LocalDatetime{}, false
		return nil
	}
	err := nd.LocalDatetime.Scan(value)
	if err != nil {
		nd.Valid = false
		return err
	}
	nd.Valid = true
	return nil
}

// MarshalJSON JSON に YYYY-MM-DD HH:MM:SS を書き込む
// encoding/json で使用するためダックタイピング
func (nd NullLocalDatetime) MarshalJSON() ([]byte, error) {
	if nd.Valid {
		return nd.LocalDatetime.MarshalJSON()
	}
	return json.Marshal(nil)
}

// UnmarshalJSON JSON の YYYY-MM-DD HH:MM:SS を UTC 時間として読み取る
// encoding/json で使用するためダックタイピング
func (nd *NullLocalDatetime) UnmarshalJSON(data []byte) error {
	if nd == nil {
		return errors.New("nil receiver of NullLocalDatetime is invalid")
	}
	if len(data) == 0 || strings.EqualFold(string(data), "null") {
		nd.LocalDatetime, nd.Valid = LocalDatetime{}, false
		return nil
	}
	if err := nd.LocalDatetime.UnmarshalJSON(data); err != nil {
		nd.LocalDatetime, nd.Valid = LocalDatetime{}, false
		return err
	}
	nd.Valid = true
	return nil
}

// NewNullLocalDatetime Year, Month, Day, Hour, Minutes, Second から UTC 時間の LocalDatetime を生成
func NewNullLocalDatetime(year, month, day uint, hour, min, sec int) NullLocalDatetime {
	datetimeUTC := NewLocalDatetime(year, month, day, hour, min, sec, UTC.Location())
	if datetimeUTC.IsZero() {
		return NullLocalDatetime{Valid: false}
	}
	return NullLocalDatetime{LocalDatetime: datetimeUTC, Valid: true}
}

// ApplyNullLocalDatetimeT Time から NullLocalDatetime を生成
func ApplyNullLocalDatetimeT(t time.Time) NullLocalDatetime {
	if t.IsZero() {
		return NullLocalDatetime{Valid: false}
	}
	return NullLocalDatetime{
		LocalDatetime: ApplyLocalDatetime(t),
		Valid:         true,
	}
}

// ApplyNullLocalDatetimeD LocalDate から NullLocalDatetime を生成
func ApplyNullLocalDatetimeD(d LocalDate) NullLocalDatetime {
	dtm := d.LocalDatetime()
	if dtm.IsZero() {
		return NullLocalDatetime{Valid: false}
	}
	return NullLocalDatetime{LocalDatetime: dtm, Valid: true}
}

// ApplyNullLocalDatetimeDt LocalDatetime から NullLocalDatetime を生成
func ApplyNullLocalDatetimeDt(d LocalDatetime) NullLocalDatetime {
	if d.IsZero() {
		return NullLocalDatetime{Valid: false}
	}
	return NullLocalDatetime{LocalDatetime: d, Valid: true}
}
