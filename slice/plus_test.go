package slice

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestPlusStrings(t *testing.T) {

	s1 := []string{"one", "two"}
	s2 := []string{"three", "four"}

	all := PlusStrings(s1, s2)

	if assert.Equal(t, 4, len(all)) {
		assert.Equal(t, "one", all[0])
		assert.Equal(t, "two", all[1])
		assert.Equal(t, "three", all[2])
		assert.Equal(t, "four", all[3])
	}

}
