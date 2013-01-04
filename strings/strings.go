package strings

import (
	"bytes"
)

// MergeStrings merges many strings together.
func MergeStrings(stringArray ...string) string {

	var buffer bytes.Buffer
	for _, v := range stringArray {
		buffer.WriteString(v)
	}
	return buffer.String()

}

// MergeStringsReversed merges many strings together backwards.
func MergeStringsReversed(stringArray ...string) string {

	var buffer bytes.Buffer
	for vi := len(stringArray); vi >= 0; vi-- {
		buffer.WriteString(stringArray[vi])
	}
	return buffer.String()

}

// JoinStrings joins many strings together separated by the specified separator.
func JoinStrings(separator string, stringArray ...string) string {

	var buffer bytes.Buffer
	var max int = len(stringArray) - 1
	for vi, v := range stringArray {
		buffer.WriteString(v)
		if vi < max {
			buffer.WriteString(separator)
		}
	}
	return buffer.String()

}

// JoinStringsReversed joins many strings together backwards separated by the specified separator.
func JoinStringsReversed(separator string, stringArray ...string) string {

	var buffer bytes.Buffer

	for vi := len(stringArray) - 1; vi >= 0; vi-- {
		buffer.WriteString(stringArray[vi])
		if vi > 0 {
			buffer.WriteString(separator)
		}
	}

	return buffer.String()

}
