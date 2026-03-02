package util

import (
	"time"
)

func ParseTime(now time.Time) string {
	dateTime := now.Format("2006-01-02")
	return dateTime
}