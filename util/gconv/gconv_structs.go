// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import "github.com/gogf/gf/v2/util/gconv/internal/converter"

// Structs converts interface{} slice to given struct slice.
// Also see Scan, Struct.
func Structs(params interface{}, pointer interface{}, paramKeyToAttrMap ...map[string]string) (err error) {
	return Scan(params, pointer, paramKeyToAttrMap...)
}

// SliceStruct is alias of Structs.
func SliceStruct(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return Structs(params, pointer, mapping...)
}

// StructsTag acts as Structs but also with support for priority tag feature, which retrieves the
// specified priorityTagAndFieldName for `params` key-value items to struct attribute names mapping.
// The parameter `priorityTag` supports multiple priorityTagAndFieldName that can be joined with char ','.
func StructsTag(params interface{}, pointer interface{}, priorityTag string) (err error) {
	return defaultConverter.Structs(params, pointer, StructsOption{
		SliceOption: converter.SliceOption{
			ContinueOnError: true,
		},
		StructOption: converter.StructOption{
			PriorityTag:     priorityTag,
			ContinueOnError: true,
		},
	})
}
