package main

import (
	"fmt"
	"sync"
)

type Number struct {
	Value int
	mutex sync.Mutex //加锁
}

func (receiver *Number) Add() {
	receiver.mutex.Lock()
	defer receiver.mutex.Unlock() //退出时会执行
	receiver.Value = receiver.Value + 1
	//fmt.Printf("add\n")
}

func (receiver *Number) Get() int {
	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()
	return receiver.Value
}

func main() {
	number := Number{Value: 0}

	wg := sync.WaitGroup{}

	n := 100_0000
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(wg *sync.WaitGroup) {
			number.Add()
			wg.Done()
		}(&wg)
	}

	wg.Wait()
	fmt.Printf("count=%d", number.Get())
}
