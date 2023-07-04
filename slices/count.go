package slices

import "strings"

func CountInt(s []int, v int) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountInt8(s []int8, v int8) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountInt16(s []int16, v int16) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountInt32(s []int32, v int32) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountInt64(s []int64, v int64) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountUint(s []uint, v uint) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountUint8(s []uint8, v uint8) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountUint16(s []uint16, v uint16) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountUint32(s []uint32, v uint32) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountUint64(s []uint64, v uint64) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountBool(s []bool, v bool) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountByte(s []byte, v byte) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountRune(s []rune, v rune) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountString(s []string, v string) int {
	var count int
	for _, e := range s {
		if e == v {
			count++
		}
	}
	return count
}

func CountStringI(s []string, v string) int {
	var count int
	for _, e := range s {
		if strings.EqualFold(e, v) {
			count++
		}
	}
	return count
}
