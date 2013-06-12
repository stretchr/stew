package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Teststrings_MergeBytes(t *testing.T) {

	assert.Equal(t, []byte("callback(jsonString)"), MergeBytes([]byte("callback"), []byte("("), []byte("jsonString"), []byte(")")))

}

func Teststrings_JoinBytes(t *testing.T) {

	assert.Equal(t, []byte("projects/centivus/accounts/tyler"), JoinBytes([]byte("/"), []byte("projects"), []byte("centivus"), []byte("accounts"), []byte("tyler")))

}
