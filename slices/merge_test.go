package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeBool(t *testing.T) {
	assert.Equal(t, []bool{}, MergeBool([]bool{}))
	assert.Equal(t, []bool{true, false}, MergeBool([]bool{true}, []bool{false}))
}

func TestMergeByte(t *testing.T) {
	assert.Equal(t, []byte{}, MergeByte([]byte{}))
	assert.Equal(t, []byte{'a', 'b'}, MergeByte([]byte{'a', 'b'}))
	assert.Equal(t, []byte{'a', 'b', 'b'}, MergeByte([]byte{'a', 'b'}, []byte{'b'}))
}

func TestMergeComplex128(t *testing.T) {
	assert.Equal(t, []complex128{}, MergeComplex128([]complex128{}))
	assert.Equal(
		t,
		[]complex128{complex(1, 0), complex(1, 1)},
		MergeComplex128([]complex128{complex(1, 0), complex(1, 1)}),
	)
	assert.Equal(t, []complex128{
		complex(1, 0),
		complex(1, 1),
		complex(1, 1),
	}, MergeComplex128([]complex128{complex(1, 0), complex(1, 1)}, []complex128{complex(1, 1)}))
}

func TestMergeComplex64(t *testing.T) {
	assert.Equal(t, []complex64{}, MergeComplex64([]complex64{}))
	assert.Equal(
		t,
		[]complex64{complex(1, 0), complex(1, 1)},
		MergeComplex64([]complex64{complex(1, 0), complex(1, 1)}),
	)
	assert.Equal(t, []complex64{
		complex(1, 0),
		complex(1, 1),
		complex(1, 1),
	}, MergeComplex64([]complex64{complex(1, 0), complex(1, 1)}, []complex64{complex(1, 1)}))
}

func TestMergeFloat32(t *testing.T) {
	assert.Equal(t, []float32{}, MergeFloat32([]float32{}))
	assert.Equal(t, []float32{0, 1}, MergeFloat32([]float32{0, 1}))
	assert.Equal(t, []float32{0, 1, 1}, MergeFloat32([]float32{0, 1}, []float32{1}))
}

func TestMergeFloat64(t *testing.T) {
	assert.Equal(t, []float64{}, MergeFloat64([]float64{}))
	assert.Equal(t, []float64{0, 1}, MergeFloat64([]float64{0, 1}))
	assert.Equal(t, []float64{0, 1, 1}, MergeFloat64([]float64{0, 1}, []float64{1}))
}

func TestMergeInt(t *testing.T) {
	assert.Equal(t, []int{}, MergeInt([]int{}))
	assert.Equal(t, []int{0, 1}, MergeInt([]int{0, 1}))
	assert.Equal(t, []int{0, 1, 1}, MergeInt([]int{0, 1}, []int{1}))
}

func TestMergeInt16(t *testing.T) {
	assert.Equal(t, []int16{}, MergeInt16([]int16{}))
	assert.Equal(t, []int16{0, 1}, MergeInt16([]int16{0, 1}))
	assert.Equal(t, []int16{0, 1, 1}, MergeInt16([]int16{0, 1}, []int16{1}))
}

func TestMergeInt32(t *testing.T) {
	assert.Equal(t, []int32{}, MergeInt32([]int32{}))
	assert.Equal(t, []int32{0, 1}, MergeInt32([]int32{0, 1}))
	assert.Equal(t, []int32{0, 1, 1}, MergeInt32([]int32{0, 1}, []int32{1}))
}

func TestMergeInt64(t *testing.T) {
	assert.Equal(t, []int64{}, MergeInt64([]int64{}))
	assert.Equal(t, []int64{0, 1}, MergeInt64([]int64{0, 1}))
	assert.Equal(t, []int64{0, 1, 1}, MergeInt64([]int64{0, 1}, []int64{1}))
}

func TestMergeInt8(t *testing.T) {
	assert.Equal(t, []int8{}, MergeInt8([]int8{}))
	assert.Equal(t, []int8{0, 1}, MergeInt8([]int8{0, 1}))
	assert.Equal(t, []int8{0, 1, 1}, MergeInt8([]int8{0, 1}, []int8{1}))
}

func TestMergeRune(t *testing.T) {
	assert.Equal(t, []rune{}, MergeRune([]rune{}))
	assert.Equal(t, []rune{'a', 'b'}, MergeRune([]rune{'a', 'b'}))
	assert.Equal(t, []rune{'a', 'b', 'b'}, MergeRune([]rune{'a', 'b'}, []rune{'b'}))
}

func TestMergeString(t *testing.T) {
	assert.Equal(t, []string{}, MergeString([]string{}))
	assert.Equal(t, []string{"foo", "bar"}, MergeString([]string{"foo", "bar"}))
	assert.Equal(t, []string{"foo", "bar", "bar"}, MergeString([]string{"foo", "bar"}, []string{"bar"}))
}

func TestMergeUint(t *testing.T) {
	assert.Equal(t, []uint{}, MergeUint([]uint{}))
	assert.Equal(t, []uint{0, 1}, MergeUint([]uint{0, 1}))
	assert.Equal(t, []uint{0, 1, 1}, MergeUint([]uint{0, 1}, []uint{1}))
}

func TestMergeUint16(t *testing.T) {
	assert.Equal(t, []uint16{}, MergeUint16([]uint16{}))
	assert.Equal(t, []uint16{0, 1}, MergeUint16([]uint16{0, 1}))
	assert.Equal(t, []uint16{0, 1, 1}, MergeUint16([]uint16{0, 1}, []uint16{1}))
}

func TestMergeUint32(t *testing.T) {
	assert.Equal(t, []uint32{}, MergeUint32([]uint32{}))
	assert.Equal(t, []uint32{0, 1}, MergeUint32([]uint32{0, 1}))
	assert.Equal(t, []uint32{0, 1, 1}, MergeUint32([]uint32{0, 1}, []uint32{1}))
}

func TestMergeUint64(t *testing.T) {
	assert.Equal(t, []uint64{}, MergeUint64([]uint64{}))
	assert.Equal(t, []uint64{0, 1}, MergeUint64([]uint64{0, 1}))
	assert.Equal(t, []uint64{0, 1, 1}, MergeUint64([]uint64{0, 1}, []uint64{1}))
}

func TestMergeUint8(t *testing.T) {
	assert.Equal(t, []uint8{}, MergeUint8([]uint8{}))
	assert.Equal(t, []uint8{0, 1}, MergeUint8([]uint8{0, 1}))
	assert.Equal(t, []uint8{0, 1, 1}, MergeUint8([]uint8{0, 1}, []uint8{1}))
}
