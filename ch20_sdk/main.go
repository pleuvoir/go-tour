package main

import "sdk/model"

func main() {

	user := model.User{}
	user.SetName("pleuvoir")
	name := user.GetName()
	println(name)

	model.Print()
}
