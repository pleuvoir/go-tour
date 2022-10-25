package main

import "fmt"

func main() {
	engine := Engine{}
	engine.Function = regular()

	function := engine.Function

	for i := 0; i < 3; i++ {
		s := function("pleuvoir")
		fmt.Printf("s is %s\n", s)
	}

}

type Engine struct {
	Function func(name string) string
}

func regular() (ret func(name string) string) {
	fmt.Printf("初始化一些东西。\n")
	return func(name string) string {
		fmt.Printf("我是worker。name is %s\n", name)
		return "我是匿名函数的返回值"
	}
}
