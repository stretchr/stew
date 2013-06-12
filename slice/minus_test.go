package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinusStrings(t *testing.T) {

	s := []string{"zero", "one", "two", "three"}
	s2 := []string{"two", "three", "four"}

	complete := MinusStrings(s, s2)

	if assert.Equal(t, 2, len(complete)) {
		assert.Equal(t, "zero", complete[0])
		assert.Equal(t, "one", complete[1])
	}

}
