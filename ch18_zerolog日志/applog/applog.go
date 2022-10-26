package applog

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

var loggerCache = make(map[string]zerolog.Logger)

// Init 初始化日志
// 用于拆分多文件
func Init(filenames ...string) {
	//递归创建文件夹
	if err := os.MkdirAll("./logs", 0777); err != nil {
		panic(fmt.Sprint("创建日志文件夹错误：", err))
		return
	}
	for _, filename := range filenames {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		//文件夹要先创建，日志文件可以通过这种方式创建
		logFile, err := os.OpenFile(fmt.Sprintf("./logs/%s.log", filename), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
		if err != nil {
			fmt.Printf("打开或者创建日志文件错误 %v\n", err)
			continue
		}
		multi := zerolog.MultiLevelWriter(consoleWriter, logFile)
		logger := zerolog.New(multi).With().Timestamp().Logger()
		loggerCache[filename] = logger
	}
}

// Logger 获取日志
func Logger(fileName string) *zerolog.Logger {
	logger := loggerCache[fileName]
	return &logger
}
