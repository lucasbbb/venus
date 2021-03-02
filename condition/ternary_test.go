package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteIf(t *testing.T) {
	assert.Equal(t, ByteIf(true, 1, 0), byte(1))
	assert.Equal(t, ByteIf(false, 1, 0), byte(0))
}

func TestComplex128If(t *testing.T) {
	assert.Equal(t, Complex128If(true, 1+1i, 0+0i), 1+1i)
	assert.Equal(t, Complex128If(false, 1+1i, 0+0i), 0+0i)
}

func TestComplex64If(t *testing.T) {
	assert.Equal(t, Complex64If(true, 1+1i, 0+0i), complex64(1+1i))
	assert.Equal(t, Complex64If(false, 1+1i, 0+0i), complex64(0+0i))
}

func TestFloat32If(t *testing.T) {
	assert.Equal(t, Float32If(true, 1, 0), float32(1))
	assert.Equal(t, Float32If(false, 1, 0), float32(0))
}

func TestFloat64If(t *testing.T) {
	assert.Equal(t, Float64If(true, 1, 0), float64(1))
	assert.Equal(t, Float64If(false, 1, 0), float64(0))
}

func TestInt16If(t *testing.T) {
	assert.Equal(t, Int16If(true, 1, 0), int16(1))
	assert.Equal(t, Int16If(false, 1, 0), int16(0))
}

func TestInt32If(t *testing.T) {
	assert.Equal(t, Int32If(true, 1, 0), int32(1))
	assert.Equal(t, Int32If(false, 1, 0), int32(0))
}

func TestInt64If(t *testing.T) {
	assert.Equal(t, Int64If(true, 1, 0), int64(1))
	assert.Equal(t, Int64If(false, 1, 0), int64(0))
}

func TestInt8If(t *testing.T) {
	assert.Equal(t, Int8If(true, 1, 0), int8(1))
	assert.Equal(t, Int8If(false, 1, 0), int8(0))
}

func TestIntIf(t *testing.T) {
	assert.Equal(t, IntIf(true, 1, 0), 1)
	assert.Equal(t, IntIf(false, 1, 0), 0)
}

func TestRuneIf(t *testing.T) {
	assert.Equal(t, RuneIf(true, 1, 0), rune(1))
	assert.Equal(t, RuneIf(false, 1, 0), rune(0))
}

func TestStringIf(t *testing.T) {
	assert.Equal(t, StringIf(true, "true", "false"), "true")
	assert.Equal(t, StringIf(false, "false", "false"), "false")
}

func TestUint16If(t *testing.T) {
	assert.Equal(t, Uint16If(true, 1, 0), uint16(1))
	assert.Equal(t, Uint16If(false, 1, 0), uint16(0))
}

func TestUint32If(t *testing.T) {
	assert.Equal(t, Uint32If(true, 1, 0), uint32(1))
	assert.Equal(t, Uint32If(false, 1, 0), uint32(0))
}

func TestUint64If(t *testing.T) {
	assert.Equal(t, Uint64If(true, 1, 0), uint64(1))
	assert.Equal(t, Uint64If(false, 1, 0), uint64(0))
}

func TestUint8If(t *testing.T) {
	assert.Equal(t, Uint8If(true, 1, 0), uint8(1))
	assert.Equal(t, Uint8If(false, 1, 0), uint8(0))
}

func TestUintIf(t *testing.T) {
	assert.Equal(t, UintIf(true, 1, 0), uint(1))
	assert.Equal(t, UintIf(false, 1, 0), uint(0))
}
