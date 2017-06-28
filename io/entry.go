package io

type Entry struct {
	LineNo    int
	Line      string
	Formatted string
}

var NullEntry = Entry{LineNo: -1}
