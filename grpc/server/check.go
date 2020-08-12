package server

import (
	"context"

	pb "github.com/shoma-www/attend_manager/grpc/proto"
	"google.golang.org/grpc"
)

type check struct {
}

// Register チェック用サーバーの登録
func Register(s *grpc.Server) {
	pb.RegisterCheckServer(s, &check{})
}

// HealthCheck ヘルスチェック用
func (c *check) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{
		Status: "success",
	}, nil
}
