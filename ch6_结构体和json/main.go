package main

import (
	"app/app"
	"encoding/json"
	"fmt"
)

func main() {

	a := app.App{}
	fmt.Printf("%s\n", a)

	app := app.App{AppName: "bot", AppVersion: "1.0.1"}

	json, _ := json.Marshal(app)

	fmt.Printf("json is %s\n", json)

}
