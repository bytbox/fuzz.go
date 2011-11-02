package fuzz

import (
	"testing"
)

type formatTest struct {
	fmt string
	fld []string
	res string
}

var formatTests = []formatTest{
	formatTest{"a", []string{}, "a"},
	formatTest{"a{B|a}", []string{}, "a"},
	formatTest{"a{B|b}", []string{"B"}, "ab"},
	formatTest{"a{B|b{C|c}}d", []string{"B"}, "abd"},
	formatTest{"a{B|b}{C|c{C|d}}", []string{"A", "B", "C"}, "abcd"},
}

func TestFormat(t *testing.T) {
	for _, test := range formatTests {
		res := compileFormat(test.fmt, test.fld)
		if res != test.res {
			t.Errorf("%s != %s", res, test.res)
		}
	}
}
