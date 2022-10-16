package main

import "fmt"

var item []int
var m = map[int]int{
	100: 1000,
}
var m2 = make(map[int]int)

func main() {

	for i := 0; i < 10; i++ {
		item = append(item, i)
		m[i] = i
		m2[i] = i
	}

	for i := range item {
		fmt.Printf("item vlaue=%d\n", i)
	}

	for key, value := range m {
		fmt.Printf("m:key=%d,value=%d\n", key, value)
	}

	for _, value := range m2 {
		fmt.Printf("m2:value=%d\n", value)
	}
}
