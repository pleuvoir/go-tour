package main

import (
	"backend/pb"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RequestService struct{}

func (s RequestService) Hello(ctx context.Context, empty *emptypb.Empty) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "hello response"}, nil
}
