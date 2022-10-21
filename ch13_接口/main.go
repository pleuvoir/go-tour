package main

import "fmt"

type Person interface {
	Say()
	SetName(name string)
}

type ZhangSan struct {
	Value string
}

func (z *ZhangSan) Say() {
	fmt.Printf("name=%s", z.Value)
}

func (z *ZhangSan) SetName(name string) {
	z.Value = name + ":hehe"
}

func main() {
	zhangSan := ZhangSan{}
	zhangSan.SetName("pleuvoir")
	zhangSan.Say()

}
