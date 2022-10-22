package main

import "fmt"

type PersonHelper struct {
	Name string
}

func (p *PersonHelper) init(name string) {
	p.Name = name
}

func main() {
	p := PersonHelper{}
	p.init("pleuvoir")

	fmt.Printf("p's Name=%s", p.Name)
}
