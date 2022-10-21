package main

import (
	"fmt"
	"time"
)

func worker() {
	//defer func() {  //不能写在主函数，最外层catch没啥用
	//	if err := recover(); err != nil {
	//		fmt.Printf("%s", err)
	//	}
	//}()
	defer recovery()
	panic("严重错误")
}

func recovery() {
	if err := recover(); err != nil {
		fmt.Printf("死机了。%s\n", err)
	}
}

func main() {
	for true {
		worker()
		time.Sleep(1 * time.Second)
	}
}
