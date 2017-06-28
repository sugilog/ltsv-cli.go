package io

import (
	"sort"
)

type Entry struct {
	LineNo    int
	Line      string
	Formatted string
}

var NullEntry = Entry{LineNo: -1}

func Sort(slice []Entry) []Entry {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].LineNo < slice[j].LineNo
	})

	return slice
}

func Nullify(slice []Entry, i int) []Entry {
	slice[i] = NullEntry
	return slice
}

func Compact(slice []Entry) []Entry {
	var new []Entry

	for _, entry := range slice {
		if entry.LineNo != NullEntry.LineNo {
			new = append(new, entry)
		}
	}

	return new
}
