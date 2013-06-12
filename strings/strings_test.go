package strings

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStrings_SplitBy(t *testing.T) {

	segs := SplitBy("ThisWillSplitByCase", func(r rune) bool {
		return strings.ToUpper(string(r)) == string(r)
	})

	if assert.Equal(t, len(segs), 5) {
		assert.Equal(t, segs[0], "This")
		assert.Equal(t, segs[1], "Will")
		assert.Equal(t, segs[2], "Split")
		assert.Equal(t, segs[3], "By")
		assert.Equal(t, segs[4], "Case")
	}

}

func TestStrings_SplitByCamelCase(t *testing.T) {

	segs := SplitByCamelCase("ThisWillSplitByCase")

	if assert.Equal(t, len(segs), 5) {
		assert.Equal(t, segs[0], "This")
		assert.Equal(t, segs[1], "Will")
		assert.Equal(t, segs[2], "Split")
		assert.Equal(t, segs[3], "By")
		assert.Equal(t, segs[4], "Case")
	}

}

func Teststrings_MergeStrings(t *testing.T) {

	assert.Equal(t, "callback(jsonString)", MergeStrings("callback", "(", "jsonString", ")"))

}

func Teststrings_MergeStringsReversed(t *testing.T) {

	assert.Equal(t, "(jsonString)callback", MergeStringsReversed("callback", "(", "jsonString", ")"))

}

func Teststrings_JoinStrings(t *testing.T) {

	assert.Equal(t, "projects/centivus/accounts/tyler", JoinStrings("/", "projects", "centivus", "accounts", "tyler"))

}

func TestStrings_JoinStringsReversed(t *testing.T) {

	assert.Equal(t, "tyler/accounts/centivus/projects", JoinStringsReversed("/", "projects", "centivus", "accounts", "tyler"))

}

func Benchmark_SprintF(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("projects/%s/accounts/%s", string(i), string(i))
	}

}

func Benchmark_Strings_Join(b *testing.B) {

	for i := 0; i < b.N; i++ {
		strings.Join([]string{"projects", string(i), "accounts", string(i)}, "/")
	}

}

func Benchmark_Stew_JoinStrings(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = JoinStrings("/", "projects", string(i), "accounts", string(i))
	}

}

func Benchmark_Stew_Bytes(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString("/")
		buffer.WriteString("projects")
		buffer.WriteString(string(i))
		buffer.WriteString("accounts")
		buffer.WriteString(string(i))
		buffer.String()
	}

}
