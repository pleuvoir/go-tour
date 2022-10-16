package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int)
var sum = 0 //是线程安全的

func consumer(wg *sync.WaitGroup) {
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				wg.Done()
				return
			}
			sum = sum + num
		}
	}
}

func producer() {
	for i := 0; i < 10_0000; i++ {
		ch <- i
	}
	close(ch) //如果不关闭则会死锁
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)
	go producer()
	go consumer(&wg)

	wg.Wait()
	fmt.Printf("sum = %d \n", sum)
}
