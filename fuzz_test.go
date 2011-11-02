package fuzz

import (
	"testing"
	"time"
)

type FuzzTest struct {
	Time   *time.Time
	Fz     Fuzzer
	Output string
}

var AbsTestTime, _ = time.Parse(time.RFC822, "06 Feb 09 1129 EST")

var fuzzTests = []FuzzTest{
	FuzzTest{AbsTestTime, FuzzNone, "11:29:00 EST, February 6, 2009"},
	FuzzTest{AbsTestTime, Cutoff("Year", "Month"), "February 2009"},
	FuzzTest{AbsTestTime, Cutoff("Year", "Minute"), "11:29, February 6, 2009"},
	FuzzTest{AbsTestTime, Cutoff("Month", "Minute"), "11:29, February 6"},
}

func TestFuzz(t *testing.T) {
	for _, test := range fuzzTests {
		res := test.Fz.Fuzz(test.Time)
		if res != test.Output {
			t.Errorf("%s != %s", res, test.Output)
		}
	}
}
