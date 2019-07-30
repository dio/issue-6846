package example

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/dio/grpctranscoding/api/example/v1"
)

type service struct{}

// NewService ...
func NewService() pb.ExampleServer {
	return &service{}
}

// Sample ...
func (s *service) Sample(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	fmt.Println("ok")
	return &empty.Empty{}, nil
}
