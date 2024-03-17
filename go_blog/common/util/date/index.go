package date

import (
	"time"
)

// DateUtils 时间
type DateUtils struct {
}

const DateTimeFormat = "2006-01-02 15:04:05"

func (c DateUtils) FormatDateTime(v time.Time) string {
	return v.Format(DateTimeFormat)
}
