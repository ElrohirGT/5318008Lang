package lib

import "strings"

// Efficiently constructs a string using the supplied builder.
type StringConstructor func(*strings.Builder)

// Constructs an empty string.
func EmptyStringConstructor(*strings.Builder) {}

// Encapsulates all data needed to construct an error.
type Error struct {
	Line        uint
	LineColumns Range[uint]
	Overview    StringConstructor
	Description StringConstructor
}

func NewError(
	line uint,
	colStart uint,
	colEnd uint,
	overview StringConstructor,
	description StringConstructor,
) Error {
	return Error{
		Line:        line,
		LineColumns: NewRange(colStart, colEnd),
		Overview:    overview,
		Description: description,
	}
}
