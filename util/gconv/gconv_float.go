// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

// Float32 converts `interface{}` to float32.
func Float32(anyInput interface{}) float32 {
	v, _ := defaultConverter.Float32(anyInput)
	return v
}

// Float64 converts `interface{}` to float64.
func Float64(anyInput interface{}) float64 {
	v, _ := defaultConverter.Float64(anyInput)
	return v
}
