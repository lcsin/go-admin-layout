package tools

import (
	"fmt"
	"time"
)

const (
	TimeLayoutDefault = "2006-01-02 15:04:05"
	TimeLayoutYMD     = "2006-01-02"
)

func TimeFormatDefault(time *time.Time) string {
	if time == nil {
		return ""
	}
	return time.Format("2006-01-02 15:04:05")
}

func TimeFormatYMD(time *time.Time) string {
	if time == nil {
		return ""
	}
	return time.Format("2006-01-02")
}

func TimeFormatSearchTime(time *time.Time) string {
	if time == nil {
		return ""
	}
	return time.Format("2006-01-02") + " 00:00:00"
}

func GetTimeWithDurationHour(hour int) time.Time {
	str := fmt.Sprintf("%dh", hour)
	duration, _ := time.ParseDuration(str)
	return time.Now().Add(duration)
}

func GetTimeWithDurationDay(day int) time.Time {
	str := fmt.Sprintf("%dh", day*24)
	duration, _ := time.ParseDuration(str)
	return time.Now().Add(duration)
}
