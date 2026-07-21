// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

// SliceInt is alias of Ints.
func SliceInt(anyInput interface{}) []int {
	return Ints(anyInput)
}

// SliceInt32 is alias of Int32s.
func SliceInt32(anyInput interface{}) []int32 {
	return Int32s(anyInput)
}

// SliceInt64 is alias of Int64s.
func SliceInt64(anyInput interface{}) []int64 {
	return Int64s(anyInput)
}

// Ints converts `interface{}` to []int.
func Ints(anyInput interface{}) []int {
	result, _ := defaultConverter.SliceInt(anyInput, SliceOption{
		ContinueOnError: true,
	})
	return result
}

// Int32s converts `interface{}` to []int32.
func Int32s(anyInput interface{}) []int32 {
	result, _ := defaultConverter.SliceInt32(anyInput, SliceOption{
		ContinueOnError: true,
	})
	return result
}

// Int64s converts `interface{}` to []int64.
func Int64s(anyInput interface{}) []int64 {
	result, _ := defaultConverter.SliceInt64(anyInput, SliceOption{
		ContinueOnError: true,
	})
	return result
}
