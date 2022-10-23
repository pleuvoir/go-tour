package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {

	Two()
}

func Two() {

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

	//文件夹要先创建，日志文件可以通过这种方式创建
	logFile, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)

	if err != nil {
		fmt.Printf("open log file err: %v\n", err)
		return
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)

	logger := zerolog.New(multi).With().Timestamp().Logger()

	logger.Info().Msg("Hello World!")
}

func One() {

	p := struct {
		Name string
		age  int
	}{"pleuovir", 18}

	//没有访问权限的不会打印出来
	log.Info().Interface("person", p).Msgf("人物信息")
	//{"level":"info","person":{"Name":"pleuovir"},"time":"2022-10-23T17:57:22+08:00","message":"人物信息"}

}
