package slices

func UniqueInt(s []int) []int {
	result := make([]int, 0, len(s))
	m := make(map[int]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueInt8(s []int8) []int8 {
	result := make([]int8, 0, len(s))
	m := make(map[int8]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueInt16(s []int16) []int16 {
	result := make([]int16, 0, len(s))
	m := make(map[int16]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueInt32(s []int32) []int32 {
	result := make([]int32, 0, len(s))
	m := make(map[int32]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueInt64(s []int64) []int64 {
	result := make([]int64, 0, len(s))
	m := make(map[int64]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueUint(s []uint) []uint {
	result := make([]uint, 0, len(s))
	m := make(map[uint]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueUint8(s []uint8) []uint8 {
	result := make([]uint8, 0, len(s))
	m := make(map[uint8]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueUint16(s []uint16) []uint16 {
	result := make([]uint16, 0, len(s))
	m := make(map[uint16]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueUint32(s []uint32) []uint32 {
	result := make([]uint32, 0, len(s))
	m := make(map[uint32]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueUint64(s []uint64) []uint64 {
	result := make([]uint64, 0, len(s))
	m := make(map[uint64]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueBool(s []bool) []bool {
	result := make([]bool, 0, len(s))
	m := make(map[bool]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueByte(s []byte) []byte {
	result := make([]byte, 0, len(s))
	m := make(map[byte]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueRune(s []rune) []rune {
	result := make([]rune, 0, len(s))
	m := make(map[rune]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}

func UniqueString(s []string) []string {
	result := make([]string, 0, len(s))
	m := make(map[string]struct{}, len(s))
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
		result = append(result, e)
	}
	return result
}
