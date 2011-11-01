package fuzz

import (
	"time"
)

type Fuzziness uint8

const (
	// No fuzzing
	FUZZ_NONE = iota

	// Time-scale fuzzing
	FUZZ_FIVE // to the nearest five minutes
	FUZZ_QUARTER // to the nearest quarter hour
	FUZZ_HOUR
	FUZZ_PERIOD // afternoon, evening, etc...

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

type Fuzzer interface {
	Fuzz(t *time.Time) string
}

type NoFuzzer struct{}
var FuzzNone = NoFuzzer{}
func (NoFuzzer) Fuzz(t *time.Time) string {
	return t.String()
}


