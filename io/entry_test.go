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

func TestSort(t *testing.T) {
	entries := make([]Entry, 5, 5)
	entries[0] = Entry{LineNo: 1}
	entries[1] = Entry{LineNo: 2}
	entries[2] = NullEntry
	entries[3] = Entry{LineNo: 0}
	entries[4] = Entry{LineNo: 3}

	newEntries := Sort(entries)

	for i := -1; i < len(entries)-1; i++ {
		if entries[i+1].LineNo != i {
			t.Error("Should be sorted")
		}

		if entries[i+1] != newEntries[i+1] {
			t.Error("Should be rewrited")
		}
	}
}

func TestNullify(t *testing.T) {
	entries := make([]Entry, 5, 5)
	entries[0] = Entry{LineNo: 1}
	entries[1] = Entry{LineNo: 2}
	entries[2] = NullEntry
	entries[3] = Entry{LineNo: 0}
	entries[4] = Entry{LineNo: 3}

	Nullify(entries, 4)

	if entries[0].LineNo != 1 {
		t.Error("[0] Should not be changed")
	}

	if entries[1].LineNo != 2 {
		t.Error("[1] Should not be changed")
	}

	if entries[2] != NullEntry {
		t.Error("[2] Should not be changed")
	}

	if entries[3].LineNo != 0 {
		t.Error("[3] Should not be changed")
	}

	if entries[4] != NullEntry {
		t.Error("[4] Should be NullEntry")
	}
}

func TestCompact(t *testing.T) {
	entries := make([]Entry, 5, 5)
	entries[0] = Entry{LineNo: 1}
	entries[1] = Entry{LineNo: 2}
	entries[2] = NullEntry
	entries[3] = Entry{LineNo: 0}
	entries[4] = Entry{LineNo: 3}

	newEntries := Compact(entries)

	if len(entries) != 5 {
		t.Error("Original slice should not be changed")
	}

	if entries[0].LineNo != 1 {
		t.Error("[0] Should not be changed")
	}

	if entries[1].LineNo != 2 {
		t.Error("[1] Should not be changed")
	}

	if entries[2] != NullEntry {
		t.Error("[2] Should not be changed")
	}

	if entries[3].LineNo != 0 {
		t.Error("[3] Should not be changed")
	}

	if entries[4].LineNo != 3 {
		t.Error("[4] Should be NullEntry")
	}

	if len(newEntries) != 4 {
		t.Error("new slice should be changed")
	}

	if newEntries[0].LineNo != 1 {
		t.Error("[0] Should not be changed")
	}

	if newEntries[1].LineNo != 2 {
		t.Error("[1] Should not be changed")
	}

	if newEntries[2].LineNo != 0 {
		t.Error("[2] Should not be changed")
	}

	if newEntries[3].LineNo != 3 {
		t.Error("[3] Should not be changed")
	}
}
