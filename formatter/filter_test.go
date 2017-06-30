package formatter

import (
	"testing"

	"github.com/sugilog/ltsv-cli.go/io"
)

func TestFilterFormatter(t *testing.T) {
	tests := []struct {
		Line   string
		Keys   []string
		Expect string
	}{
		{
			Line:   "test:line\tvalid:true",
			Keys:   []string{"test", "valid"},
			Expect: "test:line\tvalid:true",
		},
		{
			Line:   "test:line\tvalid:true",
			Keys:   []string{"test"},
			Expect: "test:line",
		},
		{
			Line:   "test:line\tvalid:true",
			Keys:   []string{"Test", "valid"},
			Expect: "valid:true",
		},
		{
			Line:   "test:line\tvalid:true",
			Keys:   []string{""},
			Expect: "",
		},
	}

	for _, test := range tests {
		formatter := FilterFormatter(test.Keys)
		entry, ok := formatter(io.Entry{Line: test.Line})

		if !ok {
			t.Error("Should be ok")
		}

		if test.Expect != entry.Formatted {
			t.Error("<%s> expected, but was <%s>", test.Expect, entry.Formatted)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	tests := []struct {
		LTSV   map[string]string
		Keys   []string
		Expect string
	}{
		{
			LTSV:   map[string]string{"test": "line", "valid": "true"},
			Keys:   []string{"test", "valid"},
			Expect: "test:line\tvalid:true",
		},
		{
			LTSV:   map[string]string{"test": "line", "valid": "true"},
			Keys:   []string{"test"},
			Expect: "test:line",
		},
		{
			LTSV:   map[string]string{"test": "line", "valid": "true"},
			Keys:   []string{"Test", "valid"},
			Expect: "valid:true",
		},
		{
			LTSV:   map[string]string{"test": "line", "valid": "true"},
			Keys:   []string{""},
			Expect: "",
		},
	}

	for _, test := range tests {
		converted := FilterFunc(test.LTSV, test.Keys)

		if test.Expect != converted {
			t.Error("<%s> expected, but was <%s>", test.Expect, converted)
		}
	}
}
