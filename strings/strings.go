package stringy

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
