package slices

import "strings"

func HasInt(s []int, v int) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasInt8(s []int8, v int8) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasInt16(s []int16, v int16) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasInt32(s []int32, v int32) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasInt64(s []int64, v int64) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasUint(s []uint, v uint) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasUint8(s []uint8, v uint8) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasUint16(s []uint16, v uint16) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasUint32(s []uint32, v uint32) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasUint64(s []uint64, v uint64) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasBool(s []bool, v bool) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasByte(s []byte, v byte) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasRune(s []rune, v rune) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasString(s []string, v string) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func HasStringI(s []string, v string) bool {
	for _, e := range s {
		if strings.ToLower(e) == strings.ToLower(v) {
			return true
		}
	}
	return false
}
