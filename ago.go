package ago

import (
	"fmt"
	"math"
	"time"
)

// time units in seconds
const MINUTE = int64(60)
const HOUR = MINUTE * 60
const DAY = HOUR * 24
const YEAR = DAY * 365

// Parse returns the duration between now and t formatted as
// 5s, 10m, 20h, 30d, 5y. t can be in the future or in the past.
func Parse(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	diff := int64(math.Abs(float64(time.Now().Unix() - t.Unix())))
	var out string
	switch {
	case diff < MINUTE:
		out = format("s", diff)
	case diff < HOUR:
		out = format("m", diff/MINUTE)
	case diff < DAY:
		out = format("h", diff/HOUR)
	case diff < YEAR:
		out = format("d", diff/DAY)
	default:
		out = format("y", diff/YEAR)
	}
	return out
}

func format(suffix string, v int64) string {
	return fmt.Sprintf("%d%s", v, suffix)
}
