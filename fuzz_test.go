package fuzz

import (
	"testing"
	"time"
)

type FuzzTest struct {
	Time      *time.Time
	Fz        Fuzzer
	Output    string
}

var TestTime, _ = time.Parse(time.RFC822, "02 Jan 06 1504 MST")

var fuzzTests = []FuzzTest{
	FuzzTest{TestTime, FuzzNone, "Mon Jan  2 15:04:00 MST 2006"},
}

func TestFuzz(t *testing.T) {
	for _, test := range fuzzTests {
		res := test.Fz.Fuzz(test.Time)
		if res != test.Output {
			t.Errorf("%s != %s", res, test.Output)
		}
	}
}

