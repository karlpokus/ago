package ago

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in  time.Time
		out string
	}{
		{time.Now().Add(-10 * time.Second), "10s"},
		{time.Now().Add(10 * time.Second), "10s"},
		{time.Now().Add(5 * time.Minute), "5m"},
		{time.Now().Add(-5 * time.Minute), "5m"},
		{time.Now().Add(5 * time.Hour), "5h"},
		{time.Now().Add(-5 * time.Hour), "5h"},
		{time.Now().AddDate(0, 0, 1), "1d"},
		{time.Now().AddDate(0, 0, -1), "1d"},
		{time.Now().AddDate(1, 0, 0), "1y"},
		{time.Now().AddDate(-1, 0, 0), "1y"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			s := Parse(tt.in)
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
