package formatter

import (
	"testing"

	"github.com/sugilog/ltsv-cli.go/io"
)

func TestFilterFormatter(t *testing.T) {
	line := "test:line\tvalid:true"

	tests := []struct {
		Keys   []string
		Expect string
	}{
		{
			Keys:   []string{"test", "valid"},
			Expect: "test:line\tvalid:true",
		},
		{
			Keys:   []string{"test"},
			Expect: "test:line",
		},
		{
			Keys:   []string{"Test", "valid"},
			Expect: "valid:true",
		},
		{
			Keys:   []string{""},
			Expect: "",
		},
	}

	for _, test := range tests {
		formatter := FilterFormatter(test.Keys)
		entry, ok := formatter(io.Entry{Line: line})

		if !ok {
			t.Error("Should be ok")
		}

		if test.Expect != entry.Formatted {
			t.Errorf("<%s> expected, but was <%s>", test.Expect, entry.Formatted)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	ltsv := map[string]string{"test": "line", "valid": "true"}

	tests := []struct {
		Keys   []string
		Expect string
	}{
		{
			Keys:   []string{"test", "valid"},
			Expect: "test:line\tvalid:true",
		},
		{
			Keys:   []string{"test"},
			Expect: "test:line",
		},
		{
			Keys:   []string{"Test", "valid"},
			Expect: "valid:true",
		},
		{
			Keys:   []string{""},
			Expect: "",
		},
	}

	for _, test := range tests {
		converted := FilterFunc(ltsv, test.Keys)

		if test.Expect != converted {
			t.Errorf("<%s> expected, but was <%s>", test.Expect, converted)
		}
	}
}
