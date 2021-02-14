package dater

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// LocalDatetime local datetime
type LocalDatetime struct {
	LocalDate LocalDate
	LocalTime LocalTime
}

// Value for go-sql-driver DB用の値へ変換します.
func (dt LocalDatetime) Value() (driver.Value, error) {
	y, m, d := dt.LocalDate.SplitString()
	h, min, sec := dt.LocalTime.SplitString()
	return y + "-" + m + "-" + d + " " + h + ":" + min + ":" + sec, nil
}

// String localDatetime to string
func (dt LocalDatetime) String() string {
	val, _ := dt.Value()
	return val.(string)
}

// Scan for go-sql-driver DBの戻り値から返還の際に使用します.
func (dt *LocalDatetime) Scan(value interface{}) error {
	if dt == nil || value == nil {
		return fmt.Errorf(fmt.Sprint(" nil value ", value))
	}
	if sv, ce := driver.String.ConvertValue(value); ce == nil {
		if v, ok := sv.(string); ok {
			groups, ge := groupSubMatch(v, LocalDateTimeRegex)
			if ge != nil {
				return fmt.Errorf("failed to convert LocalDatetime!" + ge.Error())
			} else if len(groups) < 7 {
				return fmt.Errorf("failed to convert LocalDatetime! (in grouping) len: " + strconv.Itoa(len(groups)))
			}
			y, ye := strconv.Atoi(groups[1])
			m, me := strconv.Atoi(groups[2])
			d, de := strconv.Atoi(groups[3])
			h, he := strconv.Atoi(groups[4])
			min, minErr := strconv.Atoi(groups[5])
			sec, se := strconv.Atoi(groups[6])

			if ye != nil || me != nil || de != nil || he != nil || minErr != nil || se != nil {
				return fmt.Errorf("failed to convert LocalDatetime! groups [ " + groups[1] + ", " + groups[2] + ", " + groups[3] + ", " + groups[4] + ", " + groups[5] + ", " + groups[6] + " ]")
			}
			*dt = LocalDatetime{
				LocalDate: LocalDate{Year: uint(y), Month: uint(m), Day: uint(d)},
				LocalTime: LocalTime{Hour: uint(h), Minute: uint(min), Second: uint(sec)},
			}
			return nil
		}
	}
	return fmt.Errorf("failed to scan LocalDatetime")
}

// MarshalJSON for json return format: yyyy-MM-dd hh:mm:ss
func (dt LocalDatetime) MarshalJSON() ([]byte, error) {
	if dt.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(dt.String())
}

// UnmarshalJSON for json default format: yyyy-MM-dd hh:mm:ss
//  個別にFMTを指定したい場合は, 呼び出し側で以下のようにカスタムUnmarshalレシーバを追加してください.
//  type FmtLocalDatetime LocalDatetime
//  func (dt *LocalDatetime) UnmarshalJSON(data []byte) error {
//     custom unmarshal
//  }
func (dt *LocalDatetime) UnmarshalJSON(data []byte) error {
	if dt == nil || len(data) == 0 {
		return fmt.Errorf("failed to UnmarshalJSON LocalDatetime. receiver is nil or data len is 0")
	}
	var str string
	if unmarshalErr := json.Unmarshal(data, &str); unmarshalErr != nil {
		return fmt.Errorf("failed to UnmarshalJSON LocalDatetime, err: %v", unmarshalErr)
	}
	time, parseErr := time.ParseInLocation(DateTimeHyphen.String(), str, UTC.Location())
	if parseErr != nil {
		return fmt.Errorf("failed to parse LocalDate, err: %v", parseErr)
	}
	*dt = ApplyLocalDatetimeTm(time)
	return nil
}

// ToTime locを元にTime型へ変換します.
func (dt LocalDatetime) ToTime(loc *time.Location) time.Time {
	return time.Date(int(dt.LocalDate.Year), time.Month(int(dt.LocalDate.Month)), int(dt.LocalDate.Day),
		int(dt.LocalTime.Hour), int(dt.LocalTime.Minute), int(dt.LocalTime.Second), 0, loc)
}

// ToTimeUTC UTCベースでTime型へ変換します.
func (dt LocalDatetime) ToTimeUTC() time.Time {
	loc := UTC.Location()
	return dt.ToTime(loc)
}

// Before localDatetime is before
func (dt LocalDatetime) Before(targetDateTime LocalDatetime) bool {
	firstDateTime := dt.ToTimeUTC()
	secondDateTime := targetDateTime.ToTimeUTC()
	return firstDateTime.Before(secondDateTime)
}

// After localDatetime is after
func (dt LocalDatetime) After(targetDateTime LocalDatetime) bool {
	firstDateTime := dt.ToTimeUTC()
	secondDateTime := targetDateTime.ToTimeUTC()

	return firstDateTime.After(secondDateTime)
}

// BeforeEqual localDatetime is before or equal
func (dt LocalDatetime) BeforeEqual(targetDateTime LocalDatetime) bool {
	return dt.Before(targetDateTime) || dt.Equal(targetDateTime)
}

// AfterEqual localDatetime is after or equal
func (dt LocalDatetime) AfterEqual(targetDateTime LocalDatetime) bool {
	return dt.After(targetDateTime) || dt.Equal(targetDateTime)
}

// Equal localDateme is equal
func (dt LocalDatetime) Equal(targetDateTime LocalDatetime) bool {
	return dt.ToTimeUTC().Equal(targetDateTime.ToTimeUTC())
}

// Between localDateime is between
func (dt LocalDatetime) Between(start, end LocalDatetime) bool {
	return (dt.After(start) || dt.Equal(start)) && (dt.Equal(end) || dt.Before(end))
}

// Sub dtm - target 290年以上の期間は扱えません
func (dt LocalDatetime) Sub(target LocalDatetime) (time.Duration, bool) {
	duration := dt.ToTimeUTC().Sub(target.ToTimeUTC())
	if duration == MaxDuration || duration == MinDuration {
		return duration, false
	}
	return duration, true
}

// Add localDateime add duration
func (dt LocalDatetime) Add(d time.Duration) LocalDatetime {
	loc := UTC.Location()
	addedTime := dt.ToTime(loc).Add(d)
	return ApplyLocalDatetimeTm(addedTime)
}

// AddDate localDatetime add date
func (dt LocalDatetime) AddDate(year, month, day int) LocalDatetime {
	loc := UTC.Location()
	addedTime := dt.ToTime(loc).AddDate(year, month, day)
	return ApplyLocalDatetimeTm(addedTime)
}

// IsZero localDatetime is zero?
func (dt LocalDatetime) IsZero() bool {
	return dt.LocalDate.IsZero() && dt.LocalTime.IsZero()
}

// IsNotZero localDatetime is not zero?
func (dt LocalDatetime) IsNotZero() bool {
	return !dt.IsZero()
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

// ApplyLocalDatetimeTm timeからLocalDateTimeへ変換します.(tzは無視します)
func ApplyLocalDatetimeTm(tm time.Time) LocalDatetime {
	// 発生エラーは日付としての正当性によるエラーのためtimeからの変換では不要
	return NewLocalDatetime(uint(tm.Year()), uint(tm.Month()), uint(tm.Day()), tm.Hour(), tm.Minute(), tm.Second())
}

// NowLocalDatetimeJst now localDatetime jst
func NowLocalDatetimeJst() LocalDatetime {
	return ApplyLocalDatetimeTm(NowJST())
}

// NowLocalDatetimeUTC now localDateime utc
func NowLocalDatetimeUTC() LocalDatetime {
	return ApplyLocalDatetimeTm(NowUTC())
}

// ParseLocalDatetime parse localDatetime by string
func ParseLocalDatetime(fmt DateFormat, t string) (LocalDatetime, error) {
	loc := UTC.Location() //localdatetimeのため、このtimezoneは使用しない

	tm, err := time.ParseInLocation(fmt.String(), t, loc)
	if err != nil {
		return LocalDatetime{}, err
	}
	return ApplyLocalDatetimeTm(tm), nil
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
		LocalDatetime: ApplyLocalDatetimeTm(t),
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
func ApplyNullLocalDatetimeDt(dt LocalDatetime) NullLocalDatetime {
	if dt.IsZero() {
		return NullLocalDatetime{Valid: false}
	}
	return NullLocalDatetime{LocalDatetime: dt, Valid: true}
}
