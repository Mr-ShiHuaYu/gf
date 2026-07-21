// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import (
	"github.com/Mr-ShiHuaYu/gf/v2/internal/json"
	"github.com/Mr-ShiHuaYu/gf/v2/util/gconv/internal/converter"
)

// SliceMap is alias of Maps.
func SliceMap(anyInput interface{}, option ...MapOption) []map[string]interface{} {
	return Maps(anyInput, option...)
}

// SliceMapDeep is alias of MapsDeep.
// Deprecated: used SliceMap instead.
func SliceMapDeep(anyInput interface{}) []map[string]interface{} {
	return MapsDeep(anyInput)
}

// Maps converts `value` to []map[string]interface{}.
// Note that it automatically checks and converts json string to []map if `value` is string/[]byte.
func Maps(value interface{}, option ...MapOption) []map[string]interface{} {
	mapOption := MapOption{
		ContinueOnError: true,
	}
	if len(option) > 0 {
		mapOption = option[0]
	}
	result, _ := defaultConverter.SliceMap(value, SliceMapOption{
		MapOption: mapOption,
		SliceOption: converter.SliceOption{
			ContinueOnError: true,
		},
	})
	return result
}

// MapsDeep converts `value` to []map[string]interface{} recursively.
//
// TODO completely implement the recursive converting for all types.
// Deprecated: used Maps instead.
func MapsDeep(value interface{}, tags ...string) []map[string]interface{} {
	if value == nil {
		return nil
	}
	switch r := value.(type) {
	case string:
		list := make([]map[string]interface{}, 0)
		if len(r) > 0 && r[0] == '[' && r[len(r)-1] == ']' {
			if err := json.UnmarshalUseNumber([]byte(r), &list); err != nil {
				return nil
			}
			return list
		} else {
			return nil
		}

	case []byte:
		list := make([]map[string]interface{}, 0)
		if len(r) > 0 && r[0] == '[' && r[len(r)-1] == ']' {
			if err := json.UnmarshalUseNumber(r, &list); err != nil {
				return nil
			}
			return list
		} else {
			return nil
		}

	case []map[string]interface{}:
		list := make([]map[string]interface{}, len(r))
		for k, v := range r {
			list[k] = MapDeep(v, tags...)
		}
		return list

	default:
		array := Interfaces(value)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]interface{}, len(array))
		for k, v := range array {
			list[k] = MapDeep(v, tags...)
		}
		return list
	}
}
