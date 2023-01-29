package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasInt(t *testing.T) {
	assert.True(t, HasInt([]int{0, 1}, 0))
	assert.True(t, HasInt([]int{0, 1}, 1))
	assert.False(t, HasInt([]int{0, 1}, 2))
}

func TestHasInt8(t *testing.T) {
	assert.True(t, HasInt8([]int8{0, 1}, 0))
	assert.True(t, HasInt8([]int8{0, 1}, 1))
	assert.False(t, HasInt8([]int8{0, 1}, 2))
}

func TestHasInt16(t *testing.T) {
	assert.True(t, HasInt16([]int16{0, 1}, 0))
	assert.True(t, HasInt16([]int16{0, 1}, 1))
	assert.False(t, HasInt16([]int16{0, 1}, 2))
}

func TestHasInt32(t *testing.T) {
	assert.True(t, HasInt32([]int32{0, 1}, 0))
	assert.True(t, HasInt32([]int32{0, 1}, 1))
	assert.False(t, HasInt32([]int32{0, 1}, 2))
}

func TestHasInt64(t *testing.T) {
	assert.True(t, HasInt64([]int64{0, 1}, 0))
	assert.True(t, HasInt64([]int64{0, 1}, 1))
	assert.False(t, HasInt64([]int64{0, 1}, 2))
}

func TestHasUint(t *testing.T) {
	assert.True(t, HasUint([]uint{0, 1}, 0))
	assert.True(t, HasUint([]uint{0, 1}, 1))
	assert.False(t, HasUint([]uint{0, 1}, 2))
}

func TestHasUint8(t *testing.T) {
	assert.True(t, HasUint8([]uint8{0, 1}, 0))
	assert.True(t, HasUint8([]uint8{0, 1}, 1))
	assert.False(t, HasUint8([]uint8{0, 1}, 2))
}

func TestHasUint16(t *testing.T) {
	assert.True(t, HasUint16([]uint16{0, 1}, 0))
	assert.True(t, HasUint16([]uint16{0, 1}, 1))
	assert.False(t, HasUint16([]uint16{0, 1}, 2))
}

func TestHasUint32(t *testing.T) {
	assert.True(t, HasUint32([]uint32{0, 1}, 0))
	assert.True(t, HasUint32([]uint32{0, 1}, 1))
	assert.False(t, HasUint32([]uint32{0, 1}, 2))
}

func TestHasUint64(t *testing.T) {
	assert.True(t, HasUint64([]uint64{0, 1}, 0))
	assert.True(t, HasUint64([]uint64{0, 1}, 1))
	assert.False(t, HasUint64([]uint64{0, 1}, 2))
}

func TestHasFloat32(t *testing.T) {
	assert.True(t, HasUint32([]uint32{0, 1}, 0))
	assert.True(t, HasUint32([]uint32{0., 1}, 1))
	assert.False(t, HasUint32([]uint32{0., 1}, 2))
}

func TestHasBool(t *testing.T) {
	assert.True(t, HasBool([]bool{true, false}, true))
	assert.True(t, HasBool([]bool{true, false}, false))
	assert.False(t, HasBool([]bool{true}, false))
}

func TestHasByte(t *testing.T) {
	assert.True(t, HasByte([]byte{'a', 'b'}, 'a'))
	assert.True(t, HasByte([]byte{'a', 'b'}, 'b'))
	assert.False(t, HasByte([]byte{'a', 'b'}, 'c'))
}

func TestHasRune(t *testing.T) {
	assert.True(t, HasRune([]rune{'a', 'b'}, 'a'))
	assert.True(t, HasRune([]rune{'a', 'b'}, 'b'))
	assert.False(t, HasRune([]rune{'a', 'b'}, 'c'))
}

func TestHasString(t *testing.T) {
	assert.True(t, HasString([]string{"foo", "bar"}, "foo"))
	assert.True(t, HasString([]string{"foo", "bar"}, "bar"))
	assert.False(t, HasString([]string{"foo", "bar"}, "baz"))
	assert.False(t, HasString([]string{"foo", "bar"}, "FOO"))
	assert.False(t, HasString([]string{"foo", "bar"}, "BAR"))
}

func TestHasStringI(t *testing.T) {
	assert.True(t, HasStringI([]string{"foo", "bar"}, "foo"))
	assert.True(t, HasStringI([]string{"foo", "bar"}, "bar"))
	assert.False(t, HasStringI([]string{"foo", "bar"}, "baz"))
	assert.True(t, HasStringI([]string{"foo", "bar"}, "FOO"))
	assert.True(t, HasStringI([]string{"foo", "bar"}, "BAR"))
}
