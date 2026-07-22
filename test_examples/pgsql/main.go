package main

import (
	"fmt"
	_ "github.com/Mr-ShiHuaYu/gf/contrib/drivers/pgsql/v2"
	"github.com/Mr-ShiHuaYu/gf/v2/frame/g"
)

func main() {
	fmt.Println("PostgreSQL driver import test")
	
	db := g.DB()
	fmt.Println("DB instance created:", db != nil)
}
