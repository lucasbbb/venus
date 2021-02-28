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
func ByteSlicePtr(v []byte) *[]byte          { return &v }
func TimePtr(v time.Time) *time.Time         { return &v }
