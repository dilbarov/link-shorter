package date

import (
	"gorm.io/datatypes"
	"time"
)

func ToStartOfDay(date *datatypes.Date) time.Time {
	t := time.Time(*date)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func ToEndOfDay(date *datatypes.Date) time.Time {
	t := time.Time(*date)
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}
