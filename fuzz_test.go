package fuzz

import (
	"testing"
)

type FuzzTest struct {

}

var fuzzTests = []FuzzTest{}

func TestFuzz(t *testing.T) {

}

type ParseTest struct {

}

var parseTests = []ParseTest{}

func TestParse(t *testing.T) {
	Parse("")
}

