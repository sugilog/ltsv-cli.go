package io

import (
	"testing"

	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func TestWorker(t *testing.T) {
	iomap := IOMap{
		Out: new(bytes.Buffer),
		Err: new(bytes.Buffer),
		In:  bytes.NewBufferString("1 aaa\n2 bbb\n3 ccc\n4 ddd"),
	}

	for n := 0; n < 100; n++ {
		rand.Seed(time.Now().UnixNano())

		Worker(iomap, func(entry Entry) (Entry, bool) {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			entry.Formatted = fmt.Sprintf("%03d\t%s", entry.LineNo, entry.Line)
			return entry, true
		})

		out := fmt.Sprintf("%v", iomap.Out)
		err := fmt.Sprintf("%v", iomap.Err)

		if out != "001\t1 aaa\n002\t2 bbb\n003\t3 ccc\n004\t4 ddd\n" {
			t.Errorf("Some lines incorrect\n%s", out)
		}

		if err != "" {
			t.Errorf("Error should be emtpy\n%s", err)
		}
	}
}
