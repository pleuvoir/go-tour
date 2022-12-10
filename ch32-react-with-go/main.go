package main

import (
	"backend/pb"
	"context"
	"encoding/json"
	"fmt"
	socketIO "github.com/ambelovsky/gosf-socketio"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
	"time"
)

func main() {
	g := gin.Default()
	g.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"POST"},
		AllowHeaders:           []string{"*	"},
		AllowCredentials:       false,
		ExposeHeaders:          nil,
		MaxAge:                 12 * time.Hour,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))

	requestService := RequestService{}

	server := StartSocketIO(g, func(c *socketIO.Channel, message *pb.RequestMessage) *pb.ResultMessage {
		return requestEmitHandler(requestService, c, message)
	})

	pushService := PushService{server: server}

	go func() {
		var count int64 = 0
		for {
			count++
			pushService.UpdateCount(context.Background(), &pb.UpdateCountRequest{Count: count})
			time.Sleep(time.Second)
		}
	}()

	if err := g.Run(":8000"); err != nil {
		panic(err)
	}
	//pb.RegisterBackendServer(nil, &requestService)
	//pb.RegisterPushServer(nil, &pushService)
}

func requestEmitHandler(requestService RequestService, c *socketIO.Channel, message *pb.RequestMessage) *pb.ResultMessage {
	tp := reflect.TypeOf(requestService)
	methodName := message.GetMethod()
	method, ok := tp.MethodByName(methodName)
	if !ok {
		return &pb.ResultMessage{
			Method:  methodName,
			Data:    "",
			Code:    int32(codes.NotFound),
			Message: "method not find",
		}
	}

	parameter := method.Type.In(2)
	req := reflect.New(parameter.Elem()).Interface()
	if err := json.Unmarshal([]byte(message.GetParam()), req); err != nil {
		fmt.Println(fmt.Sprintf("请求参数非json格式，%s", message.GetParam()))
	}

	in := make([]reflect.Value, 0)
	ctx := context.Background()
	in = append(in, reflect.ValueOf(ctx))
	in = append(in, reflect.ValueOf(req))

	//接口返回有错误
	call := reflect.ValueOf(&requestService).MethodByName(methodName).Call(in)
	if call[1].Interface() != nil {
		e := call[1].Interface().(error)
		st, _ := status.FromError(e)
		return &pb.ResultMessage{
			Method:  methodName,
			Data:    "",
			Code:    int32(st.Code()),
			Message: st.Message(),
		}
	}

	//返回的结构体异常
	payload, err := json.Marshal(call[0].Interface())
	if err != nil {
		return &pb.ResultMessage{
			Method:  methodName,
			Data:    "",
			Code:    int32(codes.Aborted),
			Message: "rsp payload not json error",
		}
	}

	//业务正常返回
	return &pb.ResultMessage{
		Method:  methodName,
		Data:    string(payload),
		Code:    0,
		Message: "",
	}
}
