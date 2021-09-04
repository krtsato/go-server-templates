package dater

import (
	"fmt"
	"strconv"
)

// LocalTime DB 接続時の timezone 設定による変換を回避するため使用
type LocalTime struct {
	Hour   uint
	Minute uint
	Second uint
}

// Valid 有効期間内の LocalTime を返却
func (t LocalTime) Valid() (LocalTime, error) {
	if t.Hour < MinHourOfDay || MaxHourOfDay < t.Hour {
		return t, fmt.Errorf("%d is out of LocalTime hour range", t.Hour)
	}
	if t.Minute < MinMinuteOfHour || MaxMinuteOfHour < t.Minute {
		return t, fmt.Errorf("%d is out of LocalTime minute range", t.Minute)
	}
	if t.Second < MinSecOfMinute || MaxSecOfMinute < t.Second {
		return t, fmt.Errorf("%d is out of LocalTime second range", t.Second)
	}
	return t, nil
}

// SplitString Hour, Minute, Second の桁数を揃えて返却
// ex) 12, 1, 1 -> "12", "01", "01"
func (t LocalTime) SplitString() (hourStr, minStr, secStr string) {
	var hour, min, sec string
	if t.Hour < 10 {
		hour = "0" + strconv.Itoa(int(t.Hour))
	} else {
		hour = strconv.Itoa(int(t.Hour))
	}
	if t.Minute < 10 {
		min = "0" + strconv.Itoa(int(t.Minute))
	} else {
		min = strconv.Itoa(int(t.Minute))
	}
	if t.Second < 10 {
		sec = "0" + strconv.Itoa(int(t.Second))
	} else {
		sec = strconv.Itoa(int(t.Second))
	}
	return hour, min, sec
}

// IsZero LocalTime がゼロ値のとき true を返却
func (t LocalTime) IsZero() bool {
	return t.Hour == 0 && t.Minute == 0 && t.Second == 0
}

// NewLocalTime Hour, Minute, Second から LocalTime を生成
func NewLocalTime(hour, min, sec uint) (LocalTime, error) {
	return LocalTime{Hour: hour, Minute: min, Second: sec}.Valid()
}
