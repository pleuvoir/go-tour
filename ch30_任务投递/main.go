package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	chunks := 10 //文件分成n份
	workers := 5 //个线程处理

	wg := sync.WaitGroup{}
	wg.Add(chunks)

	jobs := make(chan int, chunks) //带缓冲的管道 等于任务数

	for i := 0; i < workers; i++ {
		go handler1(i, jobs, &wg)
	}

	//将任务全部投递给worker
	scheduler(jobs, chunks)

	wg.Wait()

	fmt.Println("download finished .")
}

// 分成 chunks 份任务 里分发
// 将 n 份下载任务都到管道中去，这里管道数量等于 任务数量n 管道不会阻塞
func scheduler(jobs chan int, chunks int) {
	for i := 0; i < chunks; i++ {
		//time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		jobs <- i
	}
}

// 写法2
// 注意这里的是直接接受管道，这也是一种固定写法，下面的 range jobs 可以认为是阻塞去抢这个任务，多个线程都在抢任务
func handler2(workerId int, jobs <-chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		//	fmt.Printf("workerId[%d] job[%d] start download .\n", workerId, job)
		time.Sleep(1 * time.Second)
		fmt.Printf("workerId[%d] job[%d] download ok.\n", workerId, job)
		wg.Done() //这里不要break，这样执行完当前的线程就能继续抢了
	}
}

// 写法1，select case 多路复用
func handler1(workerId int, jobs chan int, wg *sync.WaitGroup) {
	for {
		select {
		case job, _ := <-jobs:
			//	fmt.Printf("workerId[%d] job[%d] start download .\n", workerId, job)
			time.Sleep(3 * time.Second)
			fmt.Printf("workerId[%d] job[%d] download ok.\n", workerId, job)
			wg.Done() //这里不要break，这样执行完当前的线程就能继续抢了
		}
	}
}
