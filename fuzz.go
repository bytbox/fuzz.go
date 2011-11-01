package fuzz

import (
	"os"
	"time"
)

type Fuzziness uint8

const (
	FUZZ_NONE = iota
)

type FuzzyTime struct {
	time.Time
}

func (t *FuzzyTime) String() {

}

func Parse(str string) (*FuzzyTime, os.Error) {
	return nil, nil
}

