package fuzz

import (
	"os"
	"time"
)

type Fuzziness uint8

const (
	// No fuzzing
	FUZZ_NONE = iota

	// Time-scale fuzzing
	FUZZ_FIVE
	FUZZ_TEN
	FUZZ_QUARTER
	FUZZ_HALF
	FUZZ_HOUR
	FUZZ_PERIOD // afternoon, evening, etc...

	// Date-scale fuzzing
	FUZZ_MONTH
	FUZZ_YEAR

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

type FuzzyTime struct {
	time.Time
}

func (t *FuzzyTime) Fuzz(f Fuzziness) string {
	return ""
}

func Parse(str string) (*FuzzyTime, os.Error) {
	return nil, nil
}

