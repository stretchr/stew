package stringy

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestStringy_MergeBytes(t *testing.T) {

	assert.Equal(t, []byte("callback(jsonString)"), MergeBytes([]byte("callback"), []byte("("), []byte("jsonString"), []byte(")")))

}

func TestStringy_JoinBytes(t *testing.T) {

	assert.Equal(t, []byte("projects/centivus/accounts/tyler"), JoinBytes([]byte("/"), []byte("projects"), []byte("centivus"), []byte("accounts"), []byte("tyler")))

}
