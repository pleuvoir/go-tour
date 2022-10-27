package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	worker1()
}

func worker1() {

	//总共2秒超时
	value := context.WithValue(context.Background(), "token", "pleuvoir")
	timeout, cancelFunc := context.WithTimeout(value, 5*time.Second)
	defer cancelFunc()

	//模拟任务
	fmt.Println("开始任务")
	deep := 10
	go handler(timeout, deep)

	fmt.Println("开始阻塞", time.Now())
	//等待主线程超时，阻塞操作
	select {
	case <-timeout.Done():
		fmt.Println("阻塞结束", timeout.Err(), time.Now())
	}

}

// 模拟任务处理，循环下载图片等
func handler(timeout context.Context, deep int) {

	if deep > 0 {
		fmt.Printf("[begin]token is %s %s deep=%d\n", timeout.Value("token"), time.Now(), deep)
		time.Sleep(1 * time.Second)
		go handler(timeout, deep-1)
	}

	//下面的哪个先返回 先执行哪个
	//如果整体超时 或者 当前方法超过2秒 就结束
	select {

	//等待超时会返回
	case <-timeout.Done():
		fmt.Println("超时了。", timeout.Err())
		//等待这么久 然后会返回 这个函数可不是比较时间，这里其实是在模拟处理任务，固定执行一秒 和休息一秒效果一样
		//但是休息一秒的话就不会实时返回了，所以这里实际应用可以是一个带超时的回调？
	case <-time.After(time.Second):
		fmt.Printf("[ end ]执行完成耗时一秒     %s %d\n", time.Now(), deep)
	}

}
