package formatter

import (
	"regexp"
	"testing"

	"github.com/sugilog/ltsv-cli.go/io"
)

func TestGrepFormatterWithKeys(t *testing.T) {
	line := "test:line\tvalid:true"

	tests := []struct {
		Pattern *regexp.Regexp
		Keys    []string
		Expect  string
		Ok      bool
	}{
		{
			Pattern: regexp.MustCompile("in"),
			Keys:    []string{"test", "valid"},
			Expect:  "test:line\tvalid:true",
			Ok:      true,
		},
		{
			Pattern: regexp.MustCompile("in"),
			Keys:    []string{"test"},
			Expect:  "test:line\tvalid:true",
			Ok:      true,
		},
		{
			Pattern: regexp.MustCompile("in"),
			Keys:    []string{"valid"},
			Expect:  "",
			Ok:      false,
		},
		{
			Pattern: regexp.MustCompile("in"),
			Keys:    []string{"unknown"},
			Expect:  "",
			Ok:      false,
		},
		{
			Pattern: regexp.MustCompile(".*"),
			Keys:    []string{"test"},
			Expect:  "test:line\tvalid:true",
			Ok:      true,
		},
		{
			Pattern: regexp.MustCompile(".*"),
			Keys:    []string{"unknown"},
			Expect:  "",
			Ok:      false,
		},
		{
			Pattern: regexp.MustCompile("test"),
			Keys:    []string{"test", "valid"},
			Expect:  "",
			Ok:      false,
		},
		{
			Pattern: regexp.MustCompile("test"),
			Keys:    []string{"unknown"},
			Expect:  "",
			Ok:      false,
		},
	}

	for _, test := range tests {
		formatter := GrepFormatter(test.Pattern, test.Keys)
		entry, ok := formatter(io.Entry{Line: line})

		if test.Expect != entry.Formatted {
			t.Errorf("<%s> expected, but was <%s>", test.Expect, line)
		}

		if ok != test.Ok {
			t.Errorf("<%v> expected, but was <%v>", test.Ok, ok)
		}
	}
}

func TestGrepFormatterNoKeys(t *testing.T) {
	line := "test:line\tvalid:true"

	tests := []struct {
		Pattern *regexp.Regexp
		Expect  string
		Ok      bool
	}{
		{
			Pattern: regexp.MustCompile("in"),
			Expect:  "test:line\tvalid:true",
			Ok:      true,
		},
		{
			Pattern: regexp.MustCompile(".*"),
			Expect:  "test:line\tvalid:true",
			Ok:      true,
		},
		{
			Pattern: regexp.MustCompile("test"),
			Expect:  "",
			Ok:      false,
		},
	}

	for _, test := range tests {
		formatter := GrepFormatter(test.Pattern, []string{""})
		entry, ok := formatter(io.Entry{Line: line})

		if test.Expect != entry.Formatted {
			t.Errorf("<%s> expected, but was <%s>", test.Expect, line)
		}

		if ok != test.Ok {
			t.Errorf("<%v> expected, but was <%v>", test.Ok, ok)
		}
	}

	for _, test := range tests {
		formatter := GrepFormatter(test.Pattern, []string{})
		entry, ok := formatter(io.Entry{Line: line})

		if test.Expect != entry.Formatted {
			t.Errorf("<%s> expected, but was <%s>", test.Expect, line)
		}

		if ok != test.Ok {
			t.Errorf("<%v> expected, but was <%v>", test.Ok, ok)
		}
	}
}

func TestGrepFuncWithKeys(t *testing.T) {
	ltsv := map[string]string{"test": "line", "valid": "true"}

	tests := []struct {
		Pattern *regexp.Regexp
		Keys    []string
		Expect  string
		Ok      bool
	}{
		{
			Pattern: regexp.MustCompile("in"),
			Keys:    []string{"test", "valid"},
			Expect:  "test:line\tvalid:true",
			Ok:      true,
		},
		{
			Pattern: regexp.MustCompile("in"),
			Keys:    []string{"test"},
			Expect:  "test:line\tvalid:true",
			Ok:      true,
		},
		{
			Pattern: regexp.MustCompile("in"),
			Keys:    []string{"valid"},
			Expect:  "",
			Ok:      false,
		},
		{
			Pattern: regexp.MustCompile("in"),
			Keys:    []string{"unknown"},
			Expect:  "",
			Ok:      false,
		},
		{
			Pattern: regexp.MustCompile(".*"),
			Keys:    []string{"test"},
			Expect:  "test:line\tvalid:true",
			Ok:      true,
		},
		{
			Pattern: regexp.MustCompile(".*"),
			Keys:    []string{"unknown"},
			Expect:  "",
			Ok:      false,
		},
		{
			Pattern: regexp.MustCompile("test"),
			Keys:    []string{"test", "valid"},
			Expect:  "",
			Ok:      false,
		},
		{
			Pattern: regexp.MustCompile("test"),
			Keys:    []string{"unknown"},
			Expect:  "",
			Ok:      false,
		},
	}

	for _, test := range tests {
		line, ok := GrepFuncWithKeys(ltsv, test.Pattern, test.Keys)

		if test.Expect != line {
			t.Errorf("<%s> expected, but was <%s>", test.Expect, line)
		}

		if ok != test.Ok {
			t.Errorf("<%v> expected, but was <%v>", test.Ok, ok)
		}
	}
}

func TestGrepFuncNoKeys(t *testing.T) {
	ltsv := map[string]string{"test": "line", "valid": "true"}

	tests := []struct {
		Pattern *regexp.Regexp
		Expect  string
		Ok      bool
	}{
		{
			Pattern: regexp.MustCompile("in"),
			Expect:  "test:line\tvalid:true",
			Ok:      true,
		},
		{
			Pattern: regexp.MustCompile(".*"),
			Expect:  "test:line\tvalid:true",
			Ok:      true,
		},
		{
			Pattern: regexp.MustCompile("test"),
			Expect:  "",
			Ok:      false,
		},
	}

	for _, test := range tests {
		line, ok := GrepFuncNoKeys(ltsv, test.Pattern)

		if test.Expect != line {
			t.Errorf("<%s> expected, but was <%s>", test.Expect, line)
		}

		if ok != test.Ok {
			t.Errorf("<%v> expected, but was <%v>", test.Ok, ok)
		}
	}
}
