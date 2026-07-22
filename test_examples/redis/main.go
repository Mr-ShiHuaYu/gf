package main

import (
	"fmt"
	_ "github.com/Mr-ShiHuaYu/gf/contrib/nosql/redis/v2"
	"github.com/Mr-ShiHuaYu/gf/v2/frame/g"
)

func main() {
	fmt.Println("Redis driver import test")
	
	rdb := g.Redis()
	fmt.Println("Redis instance created:", rdb != nil)
}
