package main

import (
	"backend/pb"
	"context"
	socketIO "github.com/ambelovsky/gosf-socketio"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PushService struct {
	server *socketIO.Server
}

type PushMessage struct {
	Method string
	Data   any
}

func (s PushService) Push(method string, data any) {
	s.server.BroadcastToAll("push", &PushMessage{Method: method, Data: data})
}

func (s PushService) UpdateCount(ctx context.Context, request *pb.UpdateCountRequest) (*emptypb.Empty, error) {
	s.Push("UpdateCount", request)
	return &emptypb.Empty{}, nil
}
