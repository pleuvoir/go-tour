package main

import (
	"fmt"
	"sync"
	"time"
)

func upload(waitGroup *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Printf("正在上传 i=%d \n", i)
	}
	time.Sleep(5 * time.Second)
	waitGroup.Done()
}

func saveToDb() {
	fmt.Printf("保存到数据库中\n")
	time.Sleep(3 * time.Second)
}

func main() {

	begin := time.Now()
	fmt.Printf("程序开始 %s \n", begin.Format(time.RFC850))

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	go upload(&waitGroup)
	go saveToDb()
	waitGroup.Wait()

	fmt.Printf("程序结束 耗时 %d ms ", time.Now().UnixMilli()-begin.UnixMilli())
}
