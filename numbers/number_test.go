package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromInterface(t *testing.T) {

	var num Number
	var err error

	// string
	num, err = FromInterface("100")
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}

	// bool
	num, err = FromInterface(true)
	if assert.NoError(t, err) {
		assert.Equal(t, Number(1), num)
	}
	num, err = FromInterface(false)
	if assert.NoError(t, err) {
		assert.Equal(t, Number(0), num)
	}

	// int
	num, err = FromInterface(100)
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}
	num, err = FromInterface(int8(100))
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}
	num, err = FromInterface(int16(100))
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}
	num, err = FromInterface(int32(100))
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}
	num, err = FromInterface(int64(100))
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}
	num, err = FromInterface(uint8(100))
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}
	num, err = FromInterface(uint16(100))
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}
	num, err = FromInterface(uint32(100))
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}
	num, err = FromInterface(uint64(100))
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}

	num, err = FromInterface(float32(100))
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}
	num, err = FromInterface(float64(100))
	if assert.NoError(t, err) {
		assert.Equal(t, Number(100), num)
	}

}

func TestFloat64(t *testing.T) {

	assert.Equal(t, Number(123).Float64(), float64(123))

}
