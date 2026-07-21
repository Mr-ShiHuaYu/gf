// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import (
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

// Time converts `interface{}` to time.Time.
func Time(anyInput interface{}, format ...string) time.Time {
	t, _ := defaultConverter.Time(anyInput, format...)
	return t
}

// Duration converts `interface{}` to time.Duration.
// If `interface{}` is string, then it uses time.ParseDuration to convert it.
// If `interface{}` is numeric, then it converts `interface{}` as nanoseconds.
func Duration(anyInput interface{}) time.Duration {
	d, _ := defaultConverter.Duration(anyInput)
	return d
}

// GTime converts `interface{}` to *gtime.Time.
// The parameter `format` can be used to specify the format of `interface{}`.
// It returns the converted value that matched the first format of the formats slice.
// If no `format` given, it converts `interface{}` using gtime.NewFromTimeStamp if `interface{}` is numeric,
// or using gtime.StrToTime if `interface{}` is string.
func GTime(anyInput interface{}, format ...string) *gtime.Time {
	t, _ := defaultConverter.GTime(anyInput, format...)
	return t
}
