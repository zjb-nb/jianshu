package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.lua
var lua string

func main() {
	fmt.Println(lua)
}
