package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func quit() {
	println("执行一些清理工作。。")
}

//正常的退出
//终端 CTRL+C退出
//异常退出

func main() {

	defer quit()
	println("进来了")

	//读取信号，没有一直会阻塞住
	exitChan := make(chan os.Signal)

	//监听信号
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		//有可能一次接收到多个
		for s := range signals {
			switch s {
			case syscall.SIGINT, syscall.SIGQUIT:
				println("\n监听到操作系统信号。。")
				quit() //如果监听到这个信号没处理，那么程序就不会退出了
				if i, ok := s.(syscall.Signal); ok {
					value := int(i)
					fmt.Printf("是信号类型，准备退出 %d", value)
				} else {
					println("不知道是啥，0退出")
					os.Exit(0)
				}
				//	os.Exit(value)
				exitChan <- s
			}
		}
	}()

	println("\n程序在这里被阻塞了。")
	<-exitChan
	//panic("heh")
	println("\n阻塞被终止了。")
}
