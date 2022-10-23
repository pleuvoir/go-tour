package main

import "app/applog"

func main() {

	applog.Init("hehe")

	applog.Logger("hehe").Info().Msg("我是测试")
}
