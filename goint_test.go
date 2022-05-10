package goint_test

import (
	"fmt"
	"goint"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This example returns a pointer to a variable holding the `int64` constant `2018`.
func ExampleTo_int() {
	foo := goint.To(2022)
	fmt.Println("foo contains value:", *foo)
}

// This example returns a pointer to a variable holding the `string` constant
// "point to me".
func ExampleTo_string() {
	bar := goint.To("point to me")
	fmt.Println("bar contains value:", *bar)
}

// This example returns a variable from a pointer holding the `string` constant
// "point to me".
func ExampleFrom_string() {
	val := "point to me"
	pointerToVal := &val
	bar := goint.From(pointerToVal, "Default!")
	fmt.Println("bar contains value:", bar)
}

func TestBool(t *testing.T) {
	var value bool = true
	var fallback bool = false

	result := goint.To(true)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestByte(t *testing.T) {
	var value byte = 'a'
	var fallback byte = 'b'

	result := goint.To(byte('a'))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestComplex128(t *testing.T) {
	var value complex128 = 42i
	var fallback complex128 = 10i

	result := goint.To(42i)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestComplex64(t *testing.T) {
	var value complex64 = 42i
	var fallback complex64 = 10i

	result := goint.To(complex64(42i))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestFloat32(t *testing.T) {
	var value float32 = 42.42
	var fallback float32 = 83.12

	result := goint.To(float32(42.42))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestFloat64(t *testing.T) {
	var value float64 = 42.42
	var fallback float64 = 83.12

	result := goint.To(42.42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}
func TestInt(t *testing.T) {
	var value int = 42
	var fallback int = 83

	result := goint.To(42)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestInt8(t *testing.T) {
	var value int8 = 42
	var fallback int8 = 83

	result := goint.To(int8(42))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestInt16(t *testing.T) {
	var value int16 = 42
	var fallback int16 = 83

	result := goint.To(int16(42))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestInt32(t *testing.T) {
	var value int32 = 42
	var fallback int32 = 83

	result := goint.To(int32(42))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestInt64(t *testing.T) {
	var value int64 = 42
	var fallback int64 = 83

	result := goint.To(int64(42))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestUint(t *testing.T) {
	var value uint = 42
	var fallback uint = 83

	result := goint.To(uint(42))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestUint8(t *testing.T) {
	var value uint8 = 42
	var fallback uint8 = 83

	result := goint.To(uint8(42))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestUint16(t *testing.T) {
	var value uint16 = 42
	var fallback uint16 = 83

	result := goint.To(uint16(42))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestUint32(t *testing.T) {
	var value uint32 = 42
	var fallback uint32 = 83

	result := goint.To(uint32(42))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestUint64(t *testing.T) {
	var value uint64 = 42
	var fallback uint64 = 83

	result := goint.To(uint64(42))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestString(t *testing.T) {
	var value string = "foo"
	var fallback string = "bar"

	result := goint.To("foo")

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}

func TestRune(t *testing.T) {
	var value rune = 'a'
	var fallback rune = 'b'

	result := goint.To(rune('a'))

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, goint.From(result, fallback))
	assert.Exactly(t, fallback, goint.From(nil, fallback))
}
