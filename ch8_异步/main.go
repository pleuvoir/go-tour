package main

import (
	"bufio"
	"fmt"
	"os"
)

func worker() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d\n", i)
	}
}
func main() {

	go worker()
	go worker()

	//阻塞 获取控制台的输出
	reader := bufio.NewReader(os.Stdin)
	read, err := reader.ReadBytes('\n') //注意是单引号 回车后结束控制台输出
	if err != nil {
		fmt.Printf("err is =%s\n", err)
		return
	}
	fmt.Printf("read is %s \n", read)
}
