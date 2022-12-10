package main

import (
	"app/client/rpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	ServerClient := rpc.NewServerClient(conn)

	helloResponse, err := ServerClient.Hello(context.Background(), &rpc.Empty{})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(helloResponse, err)

	registerResponse, err := ServerClient.Register(context.Background(), &rpc.RegisterRequest{Name: "test", Password: "123456"})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(registerResponse, err)

}
