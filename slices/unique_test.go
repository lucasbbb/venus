package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueBool(t *testing.T) {
	assert.Equal(t, []bool{}, UniqueBool([]bool{}))
	assert.Equal(t, []bool{true, false}, UniqueBool([]bool{true, false}))
	assert.Equal(t, []bool{true}, UniqueBool([]bool{true, true}))
}

func TestUniqueByte(t *testing.T) {
	assert.Equal(t, []byte{}, UniqueByte([]byte{}))
	assert.Equal(t, []byte{'a', 'b'}, UniqueByte([]byte{'a', 'b'}))
	assert.Equal(t, []byte{'a', 'b'}, UniqueByte([]byte{'a', 'b', 'b'}))
}

func TestUniqueInt(t *testing.T) {
	assert.Equal(t, []int{}, UniqueInt([]int{}))
	assert.Equal(t, []int{0, 1}, UniqueInt([]int{0, 1}))
	assert.Equal(t, []int{0, 1}, UniqueInt([]int{0, 1, 1}))
}

func TestUniqueInt16(t *testing.T) {
	assert.Equal(t, []int16{}, UniqueInt16([]int16{}))
	assert.Equal(t, []int16{0, 1}, UniqueInt16([]int16{0, 1}))
	assert.Equal(t, []int16{0, 1}, UniqueInt16([]int16{0, 1, 1}))
}

func TestUniqueInt32(t *testing.T) {
	assert.Equal(t, []int32{}, UniqueInt32([]int32{}))
	assert.Equal(t, []int32{0, 1}, UniqueInt32([]int32{0, 1}))
	assert.Equal(t, []int32{0, 1}, UniqueInt32([]int32{0, 1, 1}))
}

func TestUniqueInt64(t *testing.T) {
	assert.Equal(t, []int64{}, UniqueInt64([]int64{}))
	assert.Equal(t, []int64{0, 1}, UniqueInt64([]int64{0, 1}))
	assert.Equal(t, []int64{0, 1}, UniqueInt64([]int64{0, 1, 1}))
}

func TestUniqueInt8(t *testing.T) {
	assert.Equal(t, []int8{}, UniqueInt8([]int8{}))
	assert.Equal(t, []int8{0, 1}, UniqueInt8([]int8{0, 1}))
	assert.Equal(t, []int8{0, 1}, UniqueInt8([]int8{0, 1, 1}))
}

func TestUniqueRune(t *testing.T) {
	assert.Equal(t, []rune{}, UniqueRune([]rune{}))
	assert.Equal(t, []rune{'a', 'b'}, UniqueRune([]rune{'a', 'b'}))
	assert.Equal(t, []rune{'a', 'b'}, UniqueRune([]rune{'a', 'b', 'b'}))
}

func TestUniqueString(t *testing.T) {
	assert.Equal(t, []string{}, UniqueString([]string{}))
	assert.Equal(t, []string{"foo", "bar"}, UniqueString([]string{"foo", "bar"}))
	assert.Equal(t, []string{"foo", "bar"}, UniqueString([]string{"foo", "bar", "bar"}))
}

func TestUniqueUint(t *testing.T) {
	assert.Equal(t, []uint{}, UniqueUint([]uint{}))
	assert.Equal(t, []uint{0, 1}, UniqueUint([]uint{0, 1}))
	assert.Equal(t, []uint{0, 1}, UniqueUint([]uint{0, 1, 1}))
}

func TestUniqueUint16(t *testing.T) {
	assert.Equal(t, []uint16{}, UniqueUint16([]uint16{}))
	assert.Equal(t, []uint16{0, 1}, UniqueUint16([]uint16{0, 1}))
	assert.Equal(t, []uint16{0, 1}, UniqueUint16([]uint16{0, 1, 1}))
}

func TestUniqueUint32(t *testing.T) {
	assert.Equal(t, []uint32{}, UniqueUint32([]uint32{}))
	assert.Equal(t, []uint32{0, 1}, UniqueUint32([]uint32{0, 1}))
	assert.Equal(t, []uint32{0, 1}, UniqueUint32([]uint32{0, 1, 1}))
}

func TestUniqueUint64(t *testing.T) {
	assert.Equal(t, []uint64{}, UniqueUint64([]uint64{}))
	assert.Equal(t, []uint64{0, 1}, UniqueUint64([]uint64{0, 1}))
	assert.Equal(t, []uint64{0, 1}, UniqueUint64([]uint64{0, 1, 1}))
}

func TestUniqueUint8(t *testing.T) {
	assert.Equal(t, []uint8{}, UniqueUint8([]uint8{}))
	assert.Equal(t, []uint8{0, 1}, UniqueUint8([]uint8{0, 1}))
	assert.Equal(t, []uint8{0, 1}, UniqueUint8([]uint8{0, 1, 1}))
}
