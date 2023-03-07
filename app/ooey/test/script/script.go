package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Go 语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go 语言编程之旅", "帮助信息")
	flag.Parse()

	fmt.Printf("name: %s", name)
}

//go run main.go -name=eddycjy -n= 煎鱼
