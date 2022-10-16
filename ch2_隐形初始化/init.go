package main

import "fmt"

func main() {
	load()
}

func load() {
	fmt.Printf("初始化..手动%s 不错\n", "1")
}

func init() {
	fmt.Printf("隐形初始化。。\n")
}
