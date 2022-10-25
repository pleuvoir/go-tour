package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"sync"
)

var sessions = sync.Map{}

func main() {

	engine := gin.Default()

	engine.GET("/store", func(context *gin.Context) {
		uuid := uuid2.NewString()
		sessions.Store(uuid, uuid)
		context.JSON(200, struct {
			Token string `json:"token"`
		}{uuid})
	})

	engine.GET("/load", func(context *gin.Context) {

		value := struct {
			Token string `json:"token"`
		}{}

		value.Token = context.Query("token")

		load, ok := sessions.Load(value.Token)
		if ok {
			context.JSON(200, struct {
				Message string `json:"message"`
			}{fmt.Sprint("Load from sync map，result is ", load)})
		} else {
			context.JSON(404, struct {
				Message string `json:"message"`
			}{fmt.Sprint("not found key is ", value.Token)})
		}

	})

	engine.Run(":8080")
}
