package util

import (
	"fmt"
	"time"
)

func ConvertIntToTime(date string) (time.Time, error) {
	return time.Parse(time.RFC3339, fmt.Sprintf("%sT12:00:00+09:00", date))
}

func ConvertTimeToStr(t time.Time) string {
	return t.String()
}
