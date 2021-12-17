package grpcserver

import (
	"context"
	"project/pkg/fibonacci"
	"project/proto"
	"strconv"
)

type GRPCServer struct {
}

func (s *GRPCServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	var str string
	for _, value := range fibonacci.Culc(int(req.Start), int(req.End)) {
		str = str + strconv.Itoa(value) + " ; "
	}

	return &proto.AddResponse{Res: str}, nil
}
