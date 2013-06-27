package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommonStrings(t *testing.T) {

	s1 := []string{"one", "two", "three"}
	s2 := []string{"two", "three", "four"}

	c := CommonStrings(s1, s2)

	if assert.Equal(t, 2, len(c)) {
		assert.Equal(t, "two", c[0])
		assert.Equal(t, "three", c[1])
	}

}
