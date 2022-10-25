package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	workers := 1000

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker2(&wg)
	}
	wg.Wait()

	fmt.Printf("count = %d", count)
}

var count int64 = 0

func worker1(wg *sync.WaitGroup) {
	count++
	wg.Done()
}

func worker2(wg *sync.WaitGroup) {
	atomic.AddInt64(&count, 1) //特别简单
	wg.Done()
}
