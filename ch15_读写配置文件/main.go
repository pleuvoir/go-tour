package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Preferences struct {
	Name    string  `json:"name"`
	Version float64 `json:"version"`
}

const configPath = "config.json"

func main() {

	preferences := Preferences{Name: "app", Version: 100.01}
	marshal, err := json.Marshal(preferences)

	err = os.WriteFile(configPath, marshal, 777)
	if err != nil {
		fmt.Printf("写入配置文件错误，%s\n", err)
		return
	}

	//读取配置文件
	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("读取文件错误，%s\n", err)
		return
	}
	fmt.Printf("%s\n", file) //{"name":"app","version":100.01}

	//构建一个对象用来序列化
	readConfig := Preferences{}

	//反序列化
	err = json.Unmarshal(file, &readConfig)
	if err != nil {
		fmt.Printf("配置文件转换为JSON错误，%s\n", err)
	}

	fmt.Printf("%v", readConfig) //{app 100.01}

}
