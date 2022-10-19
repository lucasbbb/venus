package ptrs

import "time"

func IntPtr(v int) *int                      { return &v }
func Int8Ptr(v int8) *int8                   { return &v }
func Int16Ptr(v int16) *int16                { return &v }
func Int32Ptr(v int32) *int32                { return &v }
func Int64Ptr(v int64) *int64                { return &v }
func UintPtr(v uint) *uint                   { return &v }
func Uint8Ptr(v uint8) *uint8                { return &v }
func Uint16Ptr(v uint16) *uint16             { return &v }
func Uint32Ptr(v uint32) *uint32             { return &v }
func Uint64Ptr(v uint64) *uint64             { return &v }
func Float32Ptr(v float32) *float32          { return &v }
func Float64Ptr(v float64) *float64          { return &v }
func Complex64Ptr(v complex64) *complex64    { return &v }
func Complex128Ptr(v complex128) *complex128 { return &v }
func BoolPtr(v bool) *bool                   { return &v }
func BytePtr(v byte) *byte                   { return &v }
func RunePtr(v rune) *rune                   { return &v }
func StringPtr(v string) *string             { return &v }
func TimePtr(v time.Time) *time.Time         { return &v }

func Int(p *int, v int) int {
	if p != nil {
		return *p
	}
	return v
}

func Int8(p *int8, v int8) int8 {
	if p != nil {
		return *p
	}
	return v
}

func Int16(p *int16, v int16) int16 {
	if p != nil {
		return *p
	}
	return v
}

func Int32(p *int32, v int32) int32 {
	if p != nil {
		return *p
	}
	return v
}

func Int64(p *int64, v int64) int64 {
	if p != nil {
		return *p
	}
	return v
}

func Uint(p *uint, v uint) uint {
	if p != nil {
		return *p
	}
	return v
}

func Uint8(p *uint8, v uint8) uint8 {
	if p != nil {
		return *p
	}
	return v
}

func Uint16(p *uint16, v uint16) uint16 {
	if p != nil {
		return *p
	}
	return v
}

func Uint32(p *uint32, v uint32) uint32 {
	if p != nil {
		return *p
	}
	return v
}

func Uint64(p *uint64, v uint64) uint64 {
	if p != nil {
		return *p
	}
	return v
}

func Float32(p *float32, v float32) float32 {
	if p != nil {
		return *p
	}
	return v
}

func Float64(p *float64, v float64) float64 {
	if p != nil {
		return *p
	}
	return v
}

func Complex64(p *complex64, v complex64) complex64 {
	if p != nil {
		return *p
	}
	return v
}

func Complex128(p *complex128, v complex128) complex128 {
	if p != nil {
		return *p
	}
	return v
}

func Bool(p *bool, v bool) bool {
	if p != nil {
		return *p
	}
	return v
}

func Byte(p *byte, v byte) byte {
	if p != nil {
		return *p
	}
	return v
}

func Rune(p *rune, v rune) rune {
	if p != nil {
		return *p
	}
	return v
}

func String(p *string, v string) string {
	if p != nil {
		return *p
	}
	return v
}

func Time(p *time.Time, v time.Time) time.Time {
	if p != nil {
		return *p
	}
	return v
}
