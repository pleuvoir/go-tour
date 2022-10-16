package model

import "fmt"

var AppName = "bot"
var appVersion = "1.0.0"

func Say() {
	fmt.Printf("%s", "hello")
}

func init() {
	fmt.Printf("%s,%s", AppName, appVersion)
}
