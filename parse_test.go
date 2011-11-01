package fuzz

import (
	"testing"
)

type ParseTest struct {
	Str  string
	Time FuzzyTime
}

var parseTests = []ParseTest{}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		res, err := Parse(test.Str)
		if err != nil {
			t.Errorf("ERR: %s", err)
		}
		if res.Seconds() != test.Time.Seconds() {
			t.Errorf("%s != %s", res, test.Time)
		}
	}
}

