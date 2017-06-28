package io

import (
	"os"
	"testing"
)

func TestIOMap(t *testing.T) {
	iomap := IOMap{
		Out: os.Stdout,
		Err: os.Stderr,
		In:  os.Stdin,
	}

	if iomap.Out == nil || iomap.Err == nil || iomap.In == nil {
		t.Error("Something wrong")
	}
}
