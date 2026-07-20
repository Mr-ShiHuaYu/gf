// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

// MapToMaps converts interface{} slice type variable `params` to another map slice type variable `pointer`.
// See doMapToMaps.
func MapToMaps(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return Scan(params, pointer, mapping...)
}
