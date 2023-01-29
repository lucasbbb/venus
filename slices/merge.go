package slices

func MergeInt(parts ...[]int) []int {
	result := make([]int, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeInt8(parts ...[]int8) []int8 {
	result := make([]int8, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeInt16(parts ...[]int16) []int16 {
	result := make([]int16, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeInt32(parts ...[]int32) []int32 {
	result := make([]int32, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeInt64(parts ...[]int64) []int64 {
	result := make([]int64, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeUint(parts ...[]uint) []uint {
	result := make([]uint, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeUint8(parts ...[]uint8) []uint8 {
	result := make([]uint8, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeUint16(parts ...[]uint16) []uint16 {
	result := make([]uint16, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeUint32(parts ...[]uint32) []uint32 {
	result := make([]uint32, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeUint64(parts ...[]uint64) []uint64 {
	result := make([]uint64, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeFloat32(parts ...[]float32) []float32 {
	result := make([]float32, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeFloat64(parts ...[]float64) []float64 {
	result := make([]float64, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeComplex64(parts ...[]complex64) []complex64 {
	result := make([]complex64, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeComplex128(parts ...[]complex128) []complex128 {
	result := make([]complex128, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeBool(parts ...[]bool) []bool {
	result := make([]bool, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeByte(parts ...[]byte) []byte {
	result := make([]byte, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeRune(parts ...[]rune) []rune {
	result := make([]rune, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}

func MergeString(parts ...[]string) []string {
	result := make([]string, 0)
	for _, part := range parts {
		result = append(result, part...)
	}
	return result
}
