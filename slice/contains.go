package slice

import (
	"fmt"
	"reflect"
)

// objectsAreEqual uses multiple methods of testing equality between two
// interface{} objects. So far, it always succeeds with no false positives.
// This function is very quick if the first path is taken, otherwise it gets
// extremely slow. Up to 200 to 300 times slower.
func objectsAreEqual(left, right interface{}) bool {

	// Test for simple equality. This will not succeed for all types
	if left == right {
		return true
	}

	// Deep equality check. This will almost always succeed.
	if reflect.DeepEqual(left, right) {
		return true
	}

	// Last ditch effort. This will always succeed where the others fail
	// (at least so far)
	return fmt.Sprintf("%#v", left) == fmt.Sprintf("%#v", right)

}

// Contains determines if the "contains" argument is contained within the
// "slice" argument. The slice argument must be a slice or array.
//
// This function is performant for builtin types such as int, string, etc, though
// it is not quite as fast as a direct comparison loop would be, due to some 
// type assertion necessary to make it generic.
//
// If the type passed is not a builtin type, it is significantly slower due to 
// deep equality checks.
//
// This function is practically equal in performance to a a direct comparison loop
// for large arrays as the type assertion becomes a miniscule part of the overall
// time spent. For small arrays, it is about 7 times slower than using the
// equivalent direct call function.
//
// If you need bleeding edge performance, use one of the direct
// functions provided.
func Contains(slice, contains interface{}) bool {

	// Determine what type the "contains" variable is, then act on it
	switch slice.(type) {
	case []bool:
		return ContainsBool(slice.([]bool), contains.(bool))
	case []int:
		return ContainsInt(slice.([]int), contains.(int))
	case []int8:
		var typedContains int8
		if _, ok := contains.(int); ok {
			typedContains = int8(contains.(int))
		} else {
			typedContains = contains.(int8)
		}
		return ContainsInt8(slice.([]int8), typedContains)
	case []int16:
		var typedContains int16
		if _, ok := contains.(int); ok {
			typedContains = int16(contains.(int))
		} else {
			typedContains = contains.(int16)
		}
		return ContainsInt16(slice.([]int16), typedContains)
	case []int32:
		var typedContains int32
		if _, ok := contains.(int); ok {
			typedContains = int32(contains.(int))
		} else {
			typedContains = contains.(int32)
		}
		return ContainsInt32(slice.([]int32), typedContains)
	case []int64:
		var typedContains int64
		if _, ok := contains.(int); ok {
			typedContains = int64(contains.(int))
		} else {
			typedContains = contains.(int64)
		}
		return ContainsInt64(slice.([]int64), typedContains)
	case []uint:
		var typedContains uint
		if _, ok := contains.(int); ok {
			typedContains = uint(contains.(int))
		} else {
			typedContains = contains.(uint)
		}
		return ContainsUInt(slice.([]uint), typedContains)
	case []uint8:
		var typedContains uint8
		if _, ok := contains.(int); ok {
			typedContains = uint8(contains.(int))
		} else {
			typedContains = contains.(uint8)
		}
		return ContainsUInt8(slice.([]uint8), typedContains)
	case []uint16:
		var typedContains uint16
		if _, ok := contains.(int); ok {
			typedContains = uint16(contains.(int))
		} else {
			typedContains = contains.(uint16)
		}
		return ContainsUInt16(slice.([]uint16), typedContains)
	case []uint32:
		var typedContains uint32
		if _, ok := contains.(int); ok {
			typedContains = uint32(contains.(int))
		} else {
			typedContains = contains.(uint32)
		}
		return ContainsUInt32(slice.([]uint32), typedContains)
	case []uint64:
		var typedContains uint64
		if _, ok := contains.(int); ok {
			typedContains = uint64(contains.(int))
		} else {
			typedContains = contains.(uint64)
		}
		return ContainsUInt64(slice.([]uint64), typedContains)
	case []float32:
		var typedContains float32
		if _, ok := contains.(float64); ok {
			typedContains = float32(contains.(float64))
		} else {
			typedContains = contains.(float32)
		}
		return ContainsFloat32(slice.([]float32), typedContains)
	case []float64:
		if typedSlice, ok := slice.([]float64); ok {
			return ContainsFloat64(typedSlice, contains.(float64))
		}
	case []complex64:
		var typedContains complex64
		if _, ok := contains.(complex128); ok {
			typedContains = complex64(contains.(complex128))
		} else {
			typedContains = contains.(complex64)
		}
		return ContainsComplex64(slice.([]complex64), typedContains)
	case []complex128:
		return ContainsComplex128(slice.([]complex128), contains.(complex128))
	case []string:
		return ContainsString(slice.([]string), contains.(string))
	default:
		return ContainsObject(slice, contains)

	}

	return false

}

// ContainsBool checks if the slice has the contains value in it.
func ContainsBool(slice []bool, contains bool) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsInt checks if the slice has the contains value in it.
func ContainsInt(slice []int, contains int) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsInt8 checks if the slice has the contains value in it.
func ContainsInt8(slice []int8, contains int8) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsInt16 checks if the slice has the contains value in it.
func ContainsInt16(slice []int16, contains int16) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsInt32 checks if the slice has the contains value in it.
func ContainsInt32(slice []int32, contains int32) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsInt64 checks if the slice has the contains value in it.
func ContainsInt64(slice []int64, contains int64) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsUInt checks if the slice has the contains value in it.
func ContainsUInt(slice []uint, contains uint) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsUInt8 checks if the slice has the contains value in it.
func ContainsUInt8(slice []uint8, contains uint8) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsUInt16 checks if the slice has the contains value in it.
func ContainsUInt16(slice []uint16, contains uint16) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsUInt32 checks if the slice has the contains value in it.
func ContainsUInt32(slice []uint32, contains uint32) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsUInt64 checks if the slice has the contains value in it.
func ContainsUInt64(slice []uint64, contains uint64) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsFloat32 checks if the slice has the contains value in it.
func ContainsFloat32(slice []float32, contains float32) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsFloat64 checks if the slice has the contains value in it.
func ContainsFloat64(slice []float64, contains float64) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsComplex64 checks if the slice has the contains value in it.
func ContainsComplex64(slice []complex64, contains complex64) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsComplex128 checks if the slice has the contains value in it.
func ContainsComplex128(slice []complex128, contains complex128) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsString checks if the slice has the contains value in it.
func ContainsString(slice []string, contains string) bool {
	for _, value := range slice {
		if value == contains {
			return true
		}
	}
	return false
}

// ContainsObject checks if the slice has the contains value in it.
func ContainsObject(slice interface{}, contains interface{}) bool {
	reflectedSlice := reflect.ValueOf(slice)
	for i := 0; i < reflectedSlice.Len(); i++ {
		if objectsAreEqual(reflectedSlice.Index(i).Interface(), contains) {
			return true
		}
	}
	return false
}
