package dater

import (
	"fmt"
	"strconv"
	"time"
)

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
