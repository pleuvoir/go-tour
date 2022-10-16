package main

import (
	"bufio"
	"fmt"
	"os"
)

func worker1(id int, in chan bool, out chan bool) {
	fmt.Printf("接收到消息 id= %d \n..", id)
	in <- false
	b, done := <-in
	fmt.Printf("", b, done)
	fmt.Printf("worker1 received id = %d \n", id)
	out <- true
	fmt.Printf("worker1 write to out id = %d \n", id)

}

func worker2(in <-chan int) {
	i, ok := <-in
	fmt.Printf("worker1 received %d %s \n", i, ok)
}

func producer(num int, in chan int) {
	in <- num
}

func main() {

	//in := make(chan bool)

	var ch = make(chan bool)

	out := make(chan bool)
	workers := 3

	for i := 0; i < workers; i++ {
		go worker1(i, ch, out)
	}

	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()

}
