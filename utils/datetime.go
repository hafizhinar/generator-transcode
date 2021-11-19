package utils

import "time"

func ParseTimeZone(t time.Time) time.Time {
	loc := "Asia/Jakarta"
	tz := time.FixedZone(loc, 7*60*60)

	utc := t.UTC()
	jstTime := utc.In(tz)

	return jstTime
}

func FormatDate(t time.Time) string {
	dateFormat := "060102050415" // YYMMDDssmmhh
	return t.Format(dateFormat)
}
