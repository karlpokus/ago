package ago

import (
	"fmt"
	"time"
)

// time units in seconds
const MINUTE = int64(60)
const HOUR = MINUTE * 60
const DAY = HOUR * 24
const YEAR = DAY * 365

// Parse returns the duration between a and b formatted as
// 5s, 10m, 20h, 30d, 5y.
//
// Time arguments can be in the future or in the past and
// in any order.
func Parse(a, b time.Time) string {
	if a.IsZero() || b.IsZero() {
		return ""
	}
	v := diff(a, b)
	var s string
	switch {
	case v < MINUTE:
		s = format("s", v)
	case v < HOUR:
		s = format("m", v/MINUTE)
	case v < DAY:
		s = format("h", v/HOUR)
	case v < YEAR:
		s = format("d", v/DAY)
	default:
		s = format("y", v/YEAR)
	}
	return s
}

func diff(a, b time.Time) int64 {
	if a.Before(b) {
		return b.Unix() - a.Unix()
	}
	return a.Unix() - b.Unix()
}

func format(suffix string, v int64) string {
	return fmt.Sprintf("%d%s", v, suffix)
}

// ParseWithContext returns the duration between t and now
// and a context string showing if t is in the future or
// the past.
func ParseWithContext(t time.Time) string {
	now := time.Now()
	s := Parse(now, t)
	if t.Before(now) {
		return fmt.Sprintf("%s ago", s)
	}
	return fmt.Sprintf("in %s", s)
}
