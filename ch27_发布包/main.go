package main

import (
	"github.com/pleuvoir/applog"
)

const logFile = "default"

func main() {

	applog.Init(logFile)

	applog.Logger(logFile).Info().Str("name", "pleuvoir").Msg("引入外部包")
}
