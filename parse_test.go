package fuzz

import (
	"testing"
	"time"
)

type ParseTest struct {
	Str  string
	Time *time.Time
}

/*
	Year                 int64  // 2006 is 2006
	Month, Day           int    // Jan-2 is 1, 2
	Hour, Minute, Second int    // 15:04:05 is 15, 4, 5.
	Nanosecond           int    // Fractional second.
	Weekday              int    // Sunday, Monday, ...
	ZoneOffset           int    // seconds east of UTC, e.g. -7*60*60 for -0700
	Zone                 string // e.g., "MST"
*/

var parseTests = []ParseTest{
	ParseTest{
		"now",
		time.LocalTime(),
	},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		res, err := Parse(test.Str)
		if err != nil {
			t.Errorf("ERR: %s", err)
			continue
		}
		if res.Seconds() != test.Time.Seconds() {
			t.Errorf("%s != %s", res, test.Time)
		}
	}
}

