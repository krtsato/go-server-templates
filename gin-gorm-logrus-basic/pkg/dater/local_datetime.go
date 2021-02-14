package dater

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
		return fmt.Errorf("nil receiver of LocalDatetime is invalid")
	}
	if value == nil {
		return fmt.Errorf("failed to Scan the empty interface argument")
	}
	convVal, convErr := driver.String.ConvertValue(value)
	if convErr != nil {
		return fmt.Errorf("failed to convert LocalDatetime: %s", convErr.Error())
	}
	val, ok := convVal.(string)
	if !ok {
		return fmt.Errorf("failed to assert LocalDatetime type")
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
		return fmt.Errorf("nil receiver of LocalDatetime is invalid")
	}
	if len(data) == 0 {
		return fmt.Errorf("failed to UnmarshalJSON LocalDatetime because of zero length data")
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("failed to UnmarshalJSON LocalDatetime: %s", err.Error())
	}
	timeUTC, err := time.ParseInLocation(DateTimeHyphen.String(), str, UTC.Location())
	if err != nil {
		return fmt.Errorf("failed to parse LocalDatetime: %s", err.Error())
	}
	*d = applyLocalDatetime(timeUTC)
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
	return d.IsBefore(targetDateTime) || d.Equal(targetDateTime)
}

// IsAfterEqual LocalDatetime が引数と同日または進んだ日付のとき true を返却
func (d LocalDatetime) IsAfterEqual(targetDateTime LocalDatetime) bool {
	return d.IsAfter(targetDateTime) || d.Equal(targetDateTime)
}

// IsEqual LocalDatetime が引数と同じ日付のとき true を返却
func (d LocalDatetime) Equal(targetDateTime LocalDatetime) bool {
	return d.ToTimeUTC().Equal(targetDateTime.ToTimeUTC())
}

// IsBetween LocalDatetime が引数の範囲内にあるとき true を返却
func (d LocalDatetime) IsBetween(start, end LocalDatetime) bool {
	return (d.IsAfter(start) || d.Equal(start)) && (d.Equal(end) || d.IsBefore(end))
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

// Add localDatetime add duration
func (d LocalDatetime) Add(duration time.Duration) LocalDatetime {
	loc := UTC.Location()
	addedTime := d.ToTime(loc).Add(duration)
	return applyLocalDatetime(addedTime)
}

// AddDate localDatetime add date
func (d LocalDatetime) AddDate(year, month, day int) LocalDatetime {
	loc := UTC.Location()
	addedTime := d.ToTime(loc).AddDate(year, month, day)
	return applyLocalDatetime(addedTime)
}

// IsZero localDatetime is zero?
func (d LocalDatetime) IsZero() bool {
	return d.LocalDate.IsZero() && d.LocalTime.IsZero()
}

// IsNotZero localDatetime is not zero?
func (d LocalDatetime) IsNotZero() bool {
	return !d.IsZero()
}

// NewLocalDatetime new localDatetime.
// day: 40のように最大値を超えた数値が渡された場合, カレンダー計算を行い初期化します.
// カレンダー計算を行った結果、紀元前になる場合,空を返却します
func NewLocalDatetime(year, month, day uint, hour, min, sec int) LocalDatetime {
	tm := time.Date(int(year), time.Month(month), int(day), hour, min, sec, 0, UTC.Location())
	if tm.Unix() < FirstUnixInAD {
		return LocalDatetime{}
	}
	localTime := LocalTime{Hour: uint(tm.Hour()), Minute: uint(tm.Minute()), Second: uint(tm.Second())}
	localDate := LocalDate{Year: uint(tm.Year()), Month: uint(tm.Month()), Day: uint(tm.Day())}
	return LocalDatetime{LocalDate: localDate, LocalTime: localTime}
}

// applyLocalDatetime timeからLocalDateTimeへ変換します.(tzは無視します)
func applyLocalDatetime(tm time.Time) LocalDatetime {
	// 発生エラーは日付としての正当性によるエラーのためtimeからの変換では不要
	return NewLocalDatetime(uint(tm.Year()), uint(tm.Month()), uint(tm.Day()), tm.Hour(), tm.Minute(), tm.Second())
}

// NowLocalDatetimeJst now localDatetime jst
func NowLocalDatetimeJst() LocalDatetime {
	return applyLocalDatetime(NowJST())
}

// NowLocalDatetimeUTC now localDateime utc
func NowLocalDatetimeUTC() LocalDatetime {
	return applyLocalDatetime(NowUTC())
}

// ParseLocalDatetime parse localDatetime by string
func ParseLocalDatetime(fmt DateFormat, t string) (LocalDatetime, error) {
	loc := UTC.Location() //localdatetimeのため、このtimezoneは使用しない

	tm, err := time.ParseInLocation(fmt.String(), t, loc)
	if err != nil {
		return LocalDatetime{}, err
	}
	return applyLocalDatetime(tm), nil
}

// ========= nullable LocalDatetimeを表現します.  =========

// NullLocalDatetime nullable localDatetime
type NullLocalDatetime struct {
	LocalDatetime LocalDatetime
	Valid         bool
}

// Value for go-sql-driver to value
func (ndt NullLocalDatetime) Value() (driver.Value, error) {
	if ndt.Valid {
		return ndt.LocalDatetime.Value()
	}
	return nil, nil
}

// Scan for go-sql-driver
func (ndt *NullLocalDatetime) Scan(value interface{}) error {
	if ndt == nil || value == nil {
		ndt.LocalDatetime, ndt.Valid = LocalDatetime{}, false
		return nil
	}
	scanErr := ndt.LocalDatetime.Scan(value)
	if scanErr != nil {
		ndt.Valid = false
	} else {
		ndt.Valid = true
	}
	return scanErr
}

// MarshalJSON for json return format: yyyy-MM-dd hh:mm:ss
func (ndt NullLocalDatetime) MarshalJSON() ([]byte, error) {
	if ndt.Valid {
		return ndt.LocalDatetime.MarshalJSON()
	}
	return json.Marshal(nil)
}

// UnmarshalJSON for json default format: yyyy-MM-dd hh:mm:ss
func (ndt *NullLocalDatetime) UnmarshalJSON(data []byte) error {
	if ndt == nil {
		return fmt.Errorf("failed to UnmarshalJSON NullLocalDatetime. receiver is nil")
	}
	if len(data) == 0 || strings.EqualFold(string(data), "null") {
		ndt.LocalDatetime, ndt.Valid = LocalDatetime{}, false
		return nil
	}
	if err := ndt.LocalDatetime.UnmarshalJSON(data); err != nil {
		ndt.LocalDatetime, ndt.Valid = LocalDatetime{}, false
		return err
	}
	ndt.Valid = true
	return nil
}

// NewNullLocalDatetime new LocalDatetime
func NewNullLocalDatetime(year, month, day uint, hour, min, sec int) NullLocalDatetime {
	dtm := NewLocalDatetime(year, month, day, hour, min, sec)
	if dtm.IsZero() {
		return NullLocalDatetime{Valid: false}
	}
	return NullLocalDatetime{LocalDatetime: dtm, Valid: true}
}

// ApplyNullLocalDatetimeT time to LocalDatetime
func ApplyNullLocalDatetimeT(t time.Time) NullLocalDatetime {
	if t.IsZero() {
		return NullLocalDatetime{Valid: false}
	}
	return NullLocalDatetime{
		LocalDatetime: applyLocalDatetime(t),
		Valid:         true,
	}
}

/// ApplyNullLocalDatetimeD localDate to LocalDatetime
func ApplyNullLocalDatetimeD(d LocalDate) NullLocalDatetime {
	dtm := d.LocalDatetime()
	if dtm.IsZero() {
		return NullLocalDatetime{Valid: false}
	}
	return NullLocalDatetime{LocalDatetime: dtm, Valid: true}
}

// ApplyNullLocalDatetimeDt dtm to NullLocalDatetime
func ApplyNullLocalDatetimeDt(d LocalDatetime) NullLocalDatetime {
	if d.IsZero() {
		return NullLocalDatetime{Valid: false}
	}
	return NullLocalDatetime{LocalDatetime: d, Valid: true}
}
