package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) GetName() (string, string) {
	return p.Name, "1.0.1"
}

func worker1() {
	p := Person{}
	p.SetName("pleuvoir")
	name, _ := p.GetName()
	fmt.Printf(name)
}

// 获取方法
func worker2() {
	p := Person{}
	rv := reflect.ValueOf(&p)
	value := []reflect.Value{reflect.ValueOf("peluvoir")}
	rv.MethodByName("SetName").Call(value)
	values := rv.MethodByName("GetName").Call(nil)
	for i, v := range values {
		fmt.Printf("\ni=%d,value=%s\n", i, v)
	}
}

func worker3() {
	s := Person{}
	rt := reflect.TypeOf(s)
	if field, ok := rt.FieldByName("Name"); ok {
		tag := field.Tag.Get("json")
		fmt.Printf("tag is %s \n", tag)
	}
}

func main() {
	//正常获取
	worker1()

	//获取方法
	worker2()

	//获取标签
	worker3()
}
