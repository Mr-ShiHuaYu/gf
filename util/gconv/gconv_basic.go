// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

// Byte converts `interface{}` to byte.
func Byte(anyInput interface{}) byte {
	v, _ := defaultConverter.Uint8(anyInput)
	return v
}

// Bytes converts `interface{}` to []byte.
func Bytes(anyInput interface{}) []byte {
	v, _ := defaultConverter.Bytes(anyInput)
	return v
}

// Rune converts `interface{}` to rune.
func Rune(anyInput interface{}) rune {
	v, _ := defaultConverter.Rune(anyInput)
	return v
}

// Runes converts `interface{}` to []rune.
func Runes(anyInput interface{}) []rune {
	v, _ := defaultConverter.Runes(anyInput)
	return v
}

// String converts `interface{}` to string.
// It's most commonly used converting function.
func String(anyInput interface{}) string {
	v, _ := defaultConverter.String(anyInput)
	return v
}

// Bool converts `interface{}` to bool.
// It returns false if `interface{}` is: false, "", 0, "false", "off", "no", empty slice/map.
func Bool(anyInput interface{}) bool {
	v, _ := defaultConverter.Bool(anyInput)
	return v
}
