// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package cmd

import (
	_ "github.com/Mr-ShiHuaYu/gf/contrib/drivers/mssql/v2"
	_ "github.com/Mr-ShiHuaYu/gf/contrib/drivers/mysql/v2"
	_ "github.com/Mr-ShiHuaYu/gf/contrib/drivers/pgsql/v2"
	_ "github.com/Mr-ShiHuaYu/gf/contrib/drivers/sqlitecgo/v2"

	// do not add dm in cli pre-compilation,
	// the dm driver does not support certain target platforms.
	// _ "github.com/Mr-ShiHuaYu/gf/contrib/drivers/dm/v2"
	"github.com/Mr-ShiHuaYu/gf/cmd/gf/v2/internal/cmd/gendao"
)

type (
	cGenDao = gendao.CGenDao
)
