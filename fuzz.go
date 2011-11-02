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
	// Relative fuzzing
	REL_DEFAULT = iota // some sensible default, depending on the distance
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
	fmt, _ := compileFormat(format, Fields[startIndex:stopIndex+1])
	return t.Format(fmt)
}

type VagueFuzzer struct {
	StartField string
	StopField  string
}

func Vague(start string, stop string) VagueFuzzer {
	return VagueFuzzer{
		start,
		stop,
	}
}

func (f VagueFuzzer) Fuzz(t *time.Time) string {
	return ""
}

// Fuzzes relative to the current time
type RelativeFuzzer struct {
	StartField string
}

func Relative(start string) RelativeFuzzer {
	return RelativeFuzzer{
		StartField: start,
	}
}

func (f RelativeFuzzer) Fuzz(t *time.Time) string {
	return ""
}

// Perform a sensible amount of relative fuzzing
func RelativeFuzz(t *time.Time) string {
	return ""
}
