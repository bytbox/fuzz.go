package fuzz

import (
	"os"
	"time"
)

func Parse(str string) (*time.Time, os.Error) {
	return nil, os.NewError("hi")
}

