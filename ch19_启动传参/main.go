package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {

	//第一种方式
	args := os.Args

	for i, arg := range args {
		println(i, arg)
	}

	//第二种方式
	config := struct {
		Debug bool
		Port  int
	}{}

	flag.BoolVar(&config.Debug, "debug", true, "是否开启debug模式")
	flag.IntVar(&config.Port, "port", 80, "端口")

	flag.Parse()

	json, _ := json.Marshal(config)

	fmt.Printf("json is %s\n", json)
}
