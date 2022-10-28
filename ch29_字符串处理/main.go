package main

import (
	"fmt"
	"strings"
)

func main() {

	str := " pleuvoir  "

	trimSpace := strings.TrimSpace(str)

	fmt.Printf("去除空格 %s\n", trimSpace)

	subString := trimSpace[4:len(trimSpace)]
	fmt.Printf("subString after is %s\n", subString)

	prefix := strings.HasPrefix(subString, "vo")
	fmt.Printf("是否有前缀 vo : %v\n", prefix)

	suffix := strings.HasSuffix(subString, "ir")
	fmt.Printf("是否有后缀 ir : %v\n", suffix)

	builder := strings.Builder{}
	builder.WriteString("hello")
	builder.WriteString(" ")
	builder.WriteString("world")

	fmt.Printf("stringBuilder append is %s\n", builder.String())

	eles := []string{"1", "2"}

	join := strings.Join(eles, "@")
	fmt.Printf("join after is %s\n", join)

	//拼接格式化字符串，并且能返回
	sprintf := fmt.Sprintf("%s@%s", "1", "20")
	fmt.Printf("Sprintf after is %s\n", sprintf)

	//打印一个对象 比较清晰的方式

	person := struct {
		Name string
		Age  int
	}{"pleuvoir", 18}
	fmt.Printf("%v", person) // 输出 {Name:pleuvoir Age:18}

}
