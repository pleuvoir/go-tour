package main

import (
	app2 "app/app"
	"encoding/json"
	"fmt"
)

func main() {

	a := app2.App{}
	fmt.Printf("%s\n", a)

	app := app2.App{AppName: "bot", AppVersion: "1.0.1"}

	json, _ := json.Marshal(app)

	fmt.Printf("json is %s\n", json)

	app2.PrintHelloWorld()

}
