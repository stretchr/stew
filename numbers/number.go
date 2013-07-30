package numbers

import (
	"github.com/stretchr/stew/strings"
	"strconv"
)

type Number float64

var NumberZero Number = Number(0)
var NumberTrue Number = Number(1)
var NumberFalse Number = NumberZero

func (n Number) Float64() float64 {
	return float64(n)
}

// FromInterface creates a Number from the given value.
func FromInterface(value interface{}) (Number, error) {

	switch value.(type) {
	case string:
		value = strings.Parse(value.(string))
	}

	switch value.(type) {
	case string:
		n, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			return NumberZero, err
		}
		return Number(n), nil
	case bool:
		if value.(bool) {
			return NumberTrue, nil
		} else {
			return NumberFalse, nil
		}
	case int:
		return Number(value.(int)), nil
	case int8:
		return Number(value.(int8)), nil
	case int16:
		return Number(value.(int16)), nil
	case int32:
		return Number(value.(int32)), nil
	case int64:
		return Number(value.(int64)), nil
	case uint:
		return Number(value.(uint)), nil
	case uint8:
		return Number(value.(uint8)), nil
	case uint16:
		return Number(value.(uint16)), nil
	case uint32:
		return Number(value.(uint32)), nil
	case uint64:
		return Number(value.(uint64)), nil
	case float32:
		return Number(value.(float32)), nil
	case float64:
		return Number(value.(float64)), nil
	}

	return NumberZero, nil
}
