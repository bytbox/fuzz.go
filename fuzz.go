package fuzz

import (
	"time"
)

// TODO support interval fuzz

/*
	Year                 int64  // 2006 is 2006
	Month, Day           int    // Jan-2 is 1, 2
	Hour, Minute, Second int    // 15:04:05 is 15, 4, 5.
	Nanosecond           int    // Fractional second.
	Weekday              int    // Sunday, Monday, ...
	ZoneOffset           int    // seconds east of UTC, e.g. -7*60*60 for -0700
	Zone                 string // e.g., "MST"
*/

type Fuzziness uint8

const (
	// No fuzzing
	FUZZ_NONE = iota

	// Time-scale fuzzing
	FUZZ_FIVE    // to the nearest five minutes
	FUZZ_QUARTER // to the nearest quarter hour
	FUZZ_HOUR    // to the nearest hour
	FUZZ_PERIOD  // afternoon, evening, etc...

	// Date-scale fuzzing
	FUZZ_MONTH
	FUZZ_MONTH_ONLY // FUZZ_MONTH without "early" etc.
	FUZZ_YEAR
	FUZZ_YEAR_ONLY // FUZZ_YEAR without "early" etc.

	// Relative fuzzing
	REL_PRECISE
	REL_DEFAULT // some sensible default, depending on the distance
	REL_SECOND
	REL_MINUTE
	REL_HOUR
	REL_DAY
	REL_WEEK
	REL_MONTH
	REL_YEAR
)

var Fields = []string{
	"Year", "Month", "Day", "Hour", "Minute", "Second", "Nanosecond",
	"Weekday",
	"ZoneOffset", "Zone",
}

type Fuzzer interface {
	Fuzz(t *time.Time) string
}

type NoFuzzer struct{}

var FuzzNone = NoFuzzer{}

func (NoFuzzer) Fuzz(t *time.Time) string {
	return t.Format("15:04:05 MST, January 2, 2006")
}

type CutoffFuzzer struct {
	StartField string
	StopField  string
}

func Cutoff(start string, stop string) CutoffFuzzer {
	return CutoffFuzzer{start, stop}
}

func (f CutoffFuzzer) Fuzz(t *time.Time) string {
	var startIndex, stopIndex int
	for i, field := range Fields {
		if field == f.StartField {
			startIndex = i
		}
		if field == f.StopField {
			stopIndex = i
		}
	}
	return t.Format(compileFormat(format, Fields[startIndex:stopIndex+1]))
}

type RelativeFuzzer struct {

}
