package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {

	assert.Exactly(t, nil, Parse(""))
	assert.Exactly(t, nil, Parse("null"))

	assert.Exactly(t, int(1), Parse("1"))
	assert.Exactly(t, 9223372036854775807, Parse("9223372036854775807"))
	assert.Exactly(t, float32(92233720368547758071), Parse("92233720368547758071"))
	assert.Exactly(t, float32(1.11), Parse("1.11"))
	assert.Exactly(t, uint64(18446744073709551615), Parse("18446744073709551615"))

	assert.Exactly(t, true, Parse("true"))
	assert.Exactly(t, false, Parse("false"))
	assert.Exactly(t, true, Parse("TRUE"))
	assert.Exactly(t, false, Parse("FALSE"))
	assert.Exactly(t, "something", Parse("something"))

}

func TestParse_ForcedStrings(t *testing.T) {

	assert.Exactly(t, "true", Parse("'true'"))
	assert.Exactly(t, "true", Parse(`"true"`))

}
