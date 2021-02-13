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

// LocalDate const
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

// LocalDate DB 接続時の timezone 設定による変換を回避
// Go の time 型は timezone 情報を持っている
type LocalDate struct {
	Year  uint
	Month uint
	Day   uint
}

// Value LocalDate に応じた日付 YYYY-MM-DD を返却
// go-sql-driver で使用するためダックタイピングz
func (d *LocalDate) Value() (driver.Value, error) {
	if d == nil {
		return nil, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
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
		return fmt.Errorf("failed to scan the empty interface argument")
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
		return fmt.Errorf("failed to match LocalDate: %v", matchErr.Error())
	}
	if len(matchVals) < 4 {
		return fmt.Errorf("failed to match LocalDate in the group len: %d", len(matchVals))
	}

	year, yErr := strconv.Atoi(matchVals[1])
	month, mErr := strconv.Atoi(matchVals[2])
	day, dErr := strconv.Atoi(matchVals[3])
	if yErr != nil || mErr != nil || dErr != nil {
		return fmt.Errorf("failed to convert LocalDate groups [ %s, %s, %s ]", matchVals[1], matchVals[2], matchVals[3])
	}

	*d = LocalDate{Year: uint(year), Month: uint(month), Day: uint(day)}
}

// Valid 有効期間内の LocalDate を返却
func (d *LocalDate) Valid() (*LocalDate, error) {
	if d == nil {
		return &LocalDate{}, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	if d.Year < MinYear || MaxYear < d.Year {
		return &LocalDate{}, fmt.Errorf("%d is out of range", d.Year)
	}
	if d.Month < MinMonthOfYear || MaxMonthOfYear < d.Month {
		return &LocalDate{}, fmt.Errorf("%d is out of range", d.Month)
	}
	if d.Day < MinDayOfMonth || MaxDayOfMonth < d.Day {
		return &LocalDate{}, fmt.Errorf("%d is out of range", d.Day)
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
func (d *LocalDate) String() (string, error) {
	if d == nil {
		return "", fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	val, err := d.Value()
	if err != nil {
		return "", err
	}
	return val.(string), nil
}

// SplitString Year, Month, Day の桁数を揃えて返却
// ex) 2021, 1, 1 -> "2021", "01", "01"
func (d *LocalDate) SplitString() (yearStr, monthStr, dayStr string, err error) {
	if d == nil {
		return "", "", "", fmt.Errorf("nil receiver of LocalDate is invalid")
	}

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

	return year, month, day, nil
}

// ToTime Location に応じた時間を返却
func (d *LocalDate) ToTime(loc *time.Location) (time.Time, error) {
	if d == nil {
		return time.Time{}, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	return time.Date(int(d.Year), time.Month(int(d.Month)), int(d.Day), 0, 0, 0, 0, loc), nil
}

// ToTimeUTC UTC 時間を返却
func (d *LocalDate) ToTimeUTC() (time.Time, error) {
	if d == nil {
		return time.Time{}, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	loc := UTC.Location()
	return d.ToTime(loc)
}

// Before LocalDate が引数よりも遅れた日付のとき true を返却
func (d *LocalDate) Before(targetDate LocalDate) (bool, error) {
	if d == nil {
		return false, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	firstDate, firstErr := d.ToTimeUTC()
	if firstErr != nil {
		return false, firstErr
	}
	secondDate, secondErr := targetDate.ToTimeUTC()
	if secondErr != nil {
		return false, secondErr
	}
	return firstDate.Before(secondDate), nil
}

// After LocalDate が引数よりも進んだ日付のとき true を返却
func (d *LocalDate) After(targetDate LocalDate) (bool, error) {
	if d == nil {
		return false, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	firstDate, firstErr := d.ToTimeUTC()
	if firstErr != nil {
		return false, firstErr
	}
	secondDate, secondErr := targetDate.ToTimeUTC()
	if secondErr != nil {
		return false, secondErr
	}
	return firstDate.After(secondDate), nil
}

// Equal localDate equal?
func (d *LocalDate) Equal(targetDate LocalDate) (bool, error) {
	if d == nil {
		return false, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	return d.ToTimeUTC().Equal(targetDate.ToTimeUTC()), nil
}

// Between localDate between ?
func (d *LocalDate) Between(start, end LocalDate) (bool, error) {
	return (d.After(start) || d.Equal(start)) && (d.Equal(end) || d.Before(end)), nil
}

// IsZero localDate is zero?
func (d *LocalDate) IsZero() (bool, error) {
	if d == nil {
		return false, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	return d.Year == 0 && d.Month == 0 && d.Day == 0, nil
}

// LocalDatetime date to LocalDatetime
func (d *LocalDate) LocalDatetime() (*LocalDatetime, error) {
	if d == nil {
		return &LocalDatetime{}, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	return &LocalDatetime{
		LocalDate: *d,
		LocalTime: LocalTime{Hour: 0, Minute: 0, Second: 0},
	}, nil
}

// Sub dtm - target 290年以上の期間は扱えません
func (d *LocalDate) Sub(target LocalDate) (time.Duration, bool, error) {
	if d == nil {
		return 0, false, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	duration := d.ToTimeUTC().Sub(target.ToTimeUTC())
	if duration == MaxDuration || duration == MinDuration {
		return duration, false, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	return duration, true, nil
}

// MarshalJSON for json return format: yyyy-MM-dd
func (d *LocalDate) MarshalJSON() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	if d.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(d.String())
}

// UnmarshalJSON for json default format yyyy-MM-dd
func (d *LocalDate) UnmarshalJSON(data []byte) error {
	if d == nil {
		return fmt.Errorf("nil receiver of LocalDate is invalid")
	}
	if len(data) == 0 {
		return fmt.Errorf("failed to UnmarshalJSON LocalDate. receiver is nil or data len is 0")
	}
	var str string
	if unmarshalErr := json.Unmarshal(data, &str); unmarshalErr != nil {
		return fmt.Errorf("failed to UnmarshalJSON LocalDate, err: %v", unmarshalErr)
	}
	time, parseErr := time.ParseInLocation(DateHyphen.String(), str, UTC.Location())
	if parseErr != nil {
		return fmt.Errorf("failed to parse LocalDate, err: %v", parseErr)
	}
	*d = ApplyLocalDateByTime(time)
	return nil
}

// NewLocalDate new localDate
func NewLocalDate(year, month, day int) *LocalDate {
	tm := time.Date(year, time.Month(month), day, 0, 0, 0, 0, UTC.Location())
	if tm.Unix() < FirstUnixInAD {
		return &LocalDate{}
	}
	return &LocalDate{Year: uint(tm.Year()), Month: uint(tm.Month()), Day: uint(tm.Day())}
}

// ========= nullable LocalDateを表現します.  =========

// NullLocalDate nullable localDate
type NullLocalDate struct {
	LocalDate LocalDate
	Valid     bool
}

// Value for go-sql-driver
func (nd NullLocalDate) Value() (driver.Value, error) {
	if nd.Valid {
		return nd.LocalDate.Value()
	}
	return nil, nil
}

// Scan for go-sql-driver DBの値からの変換の際に使用します.
func (nd *NullLocalDate) Scan(value interface{}) error {
	if nd == nil || value == nil {
		nd.LocalDate, nd.Valid = LocalDate{}, false
		return nil
	}
	scanErr := nd.LocalDate.Scan(value)
	if scanErr != nil {
		nd.Valid = false
	} else {
		nd.Valid = true
	}
	return scanErr
}

// MarshalJSON for json return format: yyyy-MM-dd
func (nd NullLocalDate) MarshalJSON() ([]byte, error) {
	if nd.Valid {
		return nd.LocalDate.MarshalJSON()
	}
	return json.Marshal(nil)
}

// UnmarshalJSON for json default format yyyy-MM-dd
func (nd *NullLocalDate) UnmarshalJSON(data []byte) error {
	if nd == nil {
		return fmt.Errorf("failed to UnmarshalJSON NullLocalDate. receiver is nil")
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

// NewNullLocalDate new NullLocalDate
func NewNullLocalDate(year, month, day int) NullLocalDate {
	localDate := NewLocalDate(year, month, day)
	if localDate.IsZero() {
		return NullLocalDate{Valid: false}
	}
	return NullLocalDate{LocalDate: localDate, Valid: true}
}

// ======= LocalTimeを表現します. =========

// LocalTime 以下functionは必要に応じて追加してください
type LocalTime struct {
	Hour   uint
	Minute uint
	Second uint
}

// LocalTime validate localTime
func (t LocalTime) Valid() (LocalTime, error) {
	if t.Hour < MinHourOfDay || MaxHourOfDay < t.Hour {
		return t, fmt.Errorf("out of range ! hour :%d", t.Hour)
	}
	if t.Minute < MinMinuteOfHour || MaxMinuteOfHour < t.Minute {
		return t, fmt.Errorf("out of range ! minute: %d", t.Minute)
	}
	if t.Second < MinSecOfMinute || MaxSecOfMinute < t.Second {
		return t, fmt.Errorf("out of range ! second: %d", t.Second)
	}
	return t, nil
}

// SplitString (hour, min, sec)= (12,1,3) ---> "12", "01", "03"
func (t LocalTime) SplitString() (hour, min, sec string) {
	var h, m, d string
	if t.Hour < 10 {
		h = "0" + strconv.Itoa(int(t.Hour))
	} else {
		h = strconv.Itoa(int(t.Hour))
	}
	if t.Minute < 10 {
		m = "0" + strconv.Itoa(int(t.Minute))
	} else {
		m = strconv.Itoa(int(t.Minute))
	}

	if t.Second < 10 {
		d = "0" + strconv.Itoa(int(t.Second))
	} else {
		d = strconv.Itoa(int(t.Second))
	}
	return h, m, d
}

// IsZero localTime is Zero
func (t LocalTime) IsZero() bool {
	return t.Hour == 0 && t.Minute == 0 && t.Second == 0
}

func ApplyLocalDateByTime(tm time.Time) LocalDate {
	return LocalDate{
		Year:  uint(tm.Year()),
		Month: uint(tm.Month()),
		Day:   uint(tm.Day()),
	}
}

// NewLocalTime new localTime
func NewLocalTime(hour, min, sec uint) (LocalTime, error) {
	return LocalTime{Hour: hour, Minute: min, Second: sec}.Valid()
}

// ======= LocalDateTimeを表現します. =========

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

// AddDate localDateime add date
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
func ParseLocalDatetime(fmt Format, t string) (LocalDatetime, error) {
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
