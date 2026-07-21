// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

// MapToMap converts interface{} map type variable `params` to another map type variable `pointer`
// using reflect.
// See doMapToMap.
func MapToMap(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return Scan(params, pointer, mapping...)
}
