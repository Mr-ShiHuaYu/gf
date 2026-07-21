package main

import (
	"fmt"

	"github.com/Mr-ShiHuaYu/gf/v2/os/gbuild"
)

func main() {
	for k, v := range gbuild.Data() {
		fmt.Printf("%s: %v\n", k, v)
	}
}
