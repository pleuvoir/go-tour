package main

import (
	"backend/pb"
	"fmt"
	socketIO "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"github.com/gin-gonic/gin"
)

const requestEmit = "request"

type requestEmitFunc func(c *socketIO.Channel, message *pb.RequestMessage) *pb.ResultMessage

func StartSocketIO(g *gin.Engine, emitFunc requestEmitFunc) *socketIO.Server {
	//使用HTTP代理socketIO
	server := socketIO.NewServer(transport.GetDefaultWebsocketTransport())
	g.Any("/socket.io/*any", gin.WrapH(server))
	_ = server.On(socketIO.OnConnection, func(c *socketIO.Channel) {
		fmt.Println(fmt.Sprintf("socketIO连接成功。cid is %s", c.Id()))
		//	c.Emit("hello", "im pleuvoir")
	})

	//监听request emit
	if err := server.On(requestEmit, emitFunc); err != nil {
		panic(err)
	}

	return server
}
