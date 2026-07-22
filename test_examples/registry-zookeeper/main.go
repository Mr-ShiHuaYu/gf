package main

import (
	"fmt"
	"github.com/Mr-ShiHuaYu/gf/contrib/registry/zookeeper/v2"
	"github.com/Mr-ShiHuaYu/gf/v2/net/gsvc"
)

func main() {
	fmt.Println("Zookeeper registry import test")
	
	var registry gsvc.Registry
	registry = zookeeper.New([]string{"127.0.0.1:2181"})
	fmt.Println("Zookeeper registry instance created:", registry != nil)
}
