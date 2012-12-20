package slice

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestObjectsAreEqual(t *testing.T) {

	type FirstInterface interface{}
	type SecondInterface interface{}
	type FirstObject struct{}
	type SecondObject struct{}

	// Integer
	assert.True(t, objectsAreEqual(1, 1))
	assert.False(t, objectsAreEqual(1, 2))

	// Integer and String
	assert.False(t, objectsAreEqual(1, "1"))

	// Float
	assert.True(t, objectsAreEqual(1.0, 1.0))
	assert.False(t, objectsAreEqual(1.0, 1.1))

	// Float and Integer (do we want a float and int of the same value to be equal?)
	assert.True(t, objectsAreEqual(1.0, 1))

	// Complex
	assert.True(t, objectsAreEqual(complex(1.0, 1.0), complex(1.0, 1.0)))
	assert.False(t, objectsAreEqual(complex(1.0, 1.0), complex(1.0, 1.1)))

	// nil
	assert.True(t, objectsAreEqual(nil, nil))
	assert.False(t, objectsAreEqual(nil, 1))

	// Strings
	assert.True(t, objectsAreEqual("slice", "slice"))
	assert.False(t, objectsAreEqual("slice", "testify"))

	// Empty Interfaces
	assert.True(t, objectsAreEqual((*FirstInterface)(nil), (*FirstInterface)(nil)))
	assert.False(t, objectsAreEqual((*FirstInterface)(nil), (*SecondInterface)(nil)))

	// Objects
	assert.True(t, objectsAreEqual(new(FirstObject), new(FirstObject)))
	assert.False(t, objectsAreEqual(new(FirstObject), new(SecondObject)))

}

func TestContains(t *testing.T) {

	type oneStruct struct{}
	type twoStruct struct{}
	type threeStruct struct{}
	type fourStruct struct{}

	one := new(oneStruct)
	two := new(twoStruct)
	three := new(threeStruct)
	four := new(fourStruct)

	// Objects
	interfaceSlice := []interface{}{one, two, three}
	assert.True(t, Contains(interfaceSlice, one))
	assert.False(t, Contains(interfaceSlice, four))

	// Strings
	stringSlice := []string{"one", "two", "three"}
	assert.True(t, Contains(stringSlice, "one"))
	assert.False(t, Contains(stringSlice, "four"))

	// Integers and constant promotion
	intSlice := []int{1, 2, 3}
	int8Slice := []int8{1, 2, 3}
	int16Slice := []int{1, 2, 3}
	int32Slice := []int{1, 2, 3}
	int64Slice := []int{1, 2, 3}

	assert.True(t, Contains(intSlice, 1))
	assert.False(t, Contains(intSlice, 4))
	assert.True(t, Contains(int8Slice, 1))
	assert.False(t, Contains(int8Slice, 4))
	assert.True(t, Contains(int16Slice, 1))
	assert.False(t, Contains(int16Slice, 4))
	assert.True(t, Contains(int32Slice, 1))
	assert.False(t, Contains(int32Slice, 4))
	assert.True(t, Contains(int64Slice, 1))
	assert.False(t, Contains(int64Slice, 4))

	uIntSlice := []uint{1, 2, 3}
	uInt8Slice := []uint8{1, 2, 3}
	uInt16Slice := []uint16{1, 2, 3}
	uInt32Slice := []uint32{1, 2, 3}
	uInt64Slice := []uint64{1, 2, 3}

	assert.True(t, Contains(uIntSlice, 1))
	assert.False(t, Contains(uIntSlice, 4))
	assert.True(t, Contains(uInt8Slice, 1))
	assert.False(t, Contains(uInt8Slice, 4))
	assert.True(t, Contains(uInt16Slice, 1))
	assert.False(t, Contains(uInt16Slice, 4))
	assert.True(t, Contains(uInt32Slice, 1))
	assert.False(t, Contains(uInt32Slice, 4))
	assert.True(t, Contains(uInt64Slice, 1))
	assert.False(t, Contains(uInt64Slice, 4))

	// Floats and constant promotion
	float32Slice := []float32{1.0, 2.0, 3.0}
	float64Slice := []float64{1.0, 2.0, 3.0}

	assert.True(t, Contains(float32Slice, 1.0))
	assert.False(t, Contains(float32Slice, 1.1))
	assert.True(t, Contains(float64Slice, 1.0))
	assert.False(t, Contains(float64Slice, 1.1))

	// Complex and constant promotion
	complex64Slice := []complex64{complex(1.0, 1.0), complex(1.0, 1.1), complex(1.0, 1.2)}
	complex128Slice := []complex128{complex(1.0, 1.0), complex(1.0, 1.1), complex(1.0, 1.2)}

	assert.True(t, Contains(complex64Slice, complex(1.0, 1.0)))
	assert.False(t, Contains(complex64Slice, complex(1.0, 2.0)))
	assert.True(t, Contains(complex128Slice, complex(1.0, 1.0)))
	assert.False(t, Contains(complex128Slice, complex(1.0, 2.0)))

}

/*
	Benchmarks. Run them with "go test -bench=.*"
*/

func BenchmarkContainsObject(b *testing.B) {
	b.StopTimer()
	//uncomment this to see how it performs on a large array
	//stringSlice := []string{"one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "ten", "three"}
	stringSlice := []string{"one", "two", "three", "ten"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ContainsObject(stringSlice, "ten")
	}

}

func BenchmarkContains(b *testing.B) {

	b.StopTimer()
	//uncomment this to see how it performs on a large array
	//stringSlice := []string{"one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "ten", "three"}
	stringSlice := []string{"one", "two", "three", "ten"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Contains(stringSlice, "ten")
	}

}

func BenchmarkContainsString(b *testing.B) {

	b.StopTimer()
	//uncomment this to see how it performs on a large array
	//stringSlice := []string{"one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "two", "one", "ten", "three"}
	stringSlice := []string{"one", "two", "three", "ten"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ContainsString(stringSlice, "ten")
	}

}
