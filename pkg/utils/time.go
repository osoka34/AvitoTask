package utils

import "time"

func BeginningOfMonth(date time.Time) string {
	return date.AddDate(0, 0, -date.Day()+1).Format(time.DateOnly)
}

func EndOfMonth(date time.Time) string {
	return date.AddDate(0, 1, -date.Day()).Format(time.DateOnly)
}
