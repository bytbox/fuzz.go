package fuzz

import (
	"testing"
)

type FuzzTest struct {
	Time      FuzzyTime
	FNess     Fuzziness
	Output    string
}

var fuzzTests = []FuzzTest{}

func TestFuzz(t *testing.T) {
	for _, test := range fuzzTests {
		res := test.Time.Fuzz(test.FNess)
		if res != test.Output {
			t.Errorf("%s != %s", res, test.Output)
		}
	}
}

