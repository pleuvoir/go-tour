package main

import (
	"fmt"
	"os"
)

func _readFile() (int, error) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("error is = %s\n", err)
		return 0, err
	}
	fmt.Printf("file = %s \n", file)
	return len(file), err
}

func readFile() (int, error) {
	fileLength, err := _readFile()
	if err != nil {
		fmt.Printf("异常，存在错误 %s\n", err)
	}
	return fileLength, err
}

func main() {
	fileLength, _ := readFile()
	fmt.Printf("%d\n", fileLength)

}
