package io

import (
	"testing"
)

func TestEntry(t *testing.T) {
	entry := Entry{
		LineNo: 1,
		Line:   "test line",
	}

	if entry.LineNo != 1 {
		t.Error(".LineNo should be 1")
	}

	if entry.Line != "test line" {
		t.Error(".Line should be given string")
	}

	if entry.Formatted != "" {
		t.Error(".Formatted should be empty")
	}

	entry.Formatted = "formatted:test line"

	if entry.LineNo != 1 {
		t.Error(".LineNo should be 1")
	}

	if entry.Line != "test line" {
		t.Error(".Line should be given string")
	}

	if entry.Formatted != "formatted:test line" {
		t.Error(".Formatted should be given string")
	}
}

func TestNullEntry(t *testing.T) {
	if NullEntry.LineNo != -1 {
		t.Error("LineNo should be -1")
	}

	if NullEntry.Line != "" {
		t.Error("NullEntry.Line should be empty")
	}

	if NullEntry.Formatted != "" {
		t.Error("NullEntry.Formatted should be empty")
	}
}
