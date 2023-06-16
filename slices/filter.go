package slices

func FilterInt(s []int, filter func(v int) bool) []int {
	res := make([]int, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterInt8(s []int8, filter func(v int8) bool) []int8 {
	res := make([]int8, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterInt16(s []int16, filter func(v int16) bool) []int16 {
	res := make([]int16, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterInt32(s []int32, filter func(v int32) bool) []int32 {
	res := make([]int32, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterInt64(s []int64, filter func(v int64) bool) []int64 {
	res := make([]int64, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterUint(s []uint, filter func(v uint) bool) []uint {
	res := make([]uint, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterUint8(s []uint8, filter func(v uint8) bool) []uint8 {
	res := make([]uint8, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterUint16(s []uint16, filter func(v uint16) bool) []uint16 {
	res := make([]uint16, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterUint32(s []uint32, filter func(v uint32) bool) []uint32 {
	res := make([]uint32, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterUint64(s []uint64, filter func(v uint64) bool) []uint64 {
	res := make([]uint64, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterBool(s []bool, filter func(v bool) bool) []bool {
	res := make([]bool, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterByte(s []byte, filter func(v byte) bool) []byte {
	res := make([]byte, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterRune(s []rune, filter func(v rune) bool) []rune {
	res := make([]rune, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterString(s []string, filter func(v string) bool) []string {
	res := make([]string, 0, len(s))
	for _, e := range s {
		if filter(e) {
			continue
		}
		res = append(res, e)
	}
	return res
}

func FilterStringEmpty(s []string) []string {
	return FilterString(s, func(v string) bool {
		return v == ""
	})
}
