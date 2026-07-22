package main

import (
	"fmt"
	"github.com/Mr-ShiHuaYu/gf/contrib/registry/file/v2"
	"github.com/Mr-ShiHuaYu/gf/v2/net/gsvc"
)

func main() {
	fmt.Println("File registry import test")
	
	var registry gsvc.Registry
	registry = file.New("/tmp/registry")
	fmt.Println("File registry instance created:", registry != nil)
}
