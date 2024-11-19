package ago

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		a   time.Time
		b   time.Time
		out string
	}{
		// now <> past
		{time.Now(), time.Now().Add(-10 * time.Second), "10s"},
		{time.Now(), time.Now().Add(-5 * time.Minute), "5m"},
		{time.Now(), time.Now().Add(-5 * time.Hour), "5h"},
		{time.Now(), time.Now().AddDate(0, 0, -1), "1d"},
		{time.Now(), time.Now().AddDate(-1, 0, 0), "1y"},
		{time.Now().Add(-10 * time.Second), time.Now(), "10s"},
		{time.Now().AddDate(0, 0, -1), time.Now(), "1d"},
		// now <> future
		{time.Now(), time.Now().Add(10 * time.Second), "10s"},
		{time.Now(), time.Now().Add(5 * time.Minute), "5m"},
		{time.Now(), time.Now().Add(5 * time.Hour), "5h"},
		{time.Now(), time.Now().AddDate(0, 0, 1), "1d"},
		{time.Now(), time.Now().AddDate(1, 0, 0), "1y"},
		{time.Now().Add(10 * time.Second), time.Now(), "10s"},
		// past <> past
		{time.Now().Add(-10 * time.Second), time.Now().Add(-20 * time.Second), "10s"},
		{time.Now().Add(-20 * time.Second), time.Now().Add(-10 * time.Second), "10s"},
		// future <> future
		{time.Now().Add(10 * time.Second), time.Now().Add(20 * time.Second), "10s"},
		// past <> future
		{time.Now().Add(-30 * time.Second), time.Now().Add(30 * time.Second), "1m"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			s := Parse(tt.a, tt.b)
			if s != tt.out {
				t.Fatalf("%s != %s", s, tt.out)
			}
		})
	}
}

func TestParseWithContext(t *testing.T) {
	tests := []struct {
		in  time.Time
		out string
	}{
		{time.Now().Add(-10 * time.Second), "10s ago"},
		{time.Now().Add(10 * time.Second), "in 10s"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			s := ParseWithContext(tt.in)
			if s != tt.out {
				t.Fatalf("%s != %s", s, tt.out)
			}
		})
	}
}
