package server

import (
	"context"

	"github.com/shoma-www/attend_manager/core"
	pb "github.com/shoma-www/attend_manager/grpc/proto"
)

type check struct {
	logger core.Logger
}

func NewCheckServer(l core.Logger) *check {
	return &check{logger: l}
}

// HealthCheck ヘルスチェック用
func (c *check) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	c.logger.WithUUID(ctx).Info("Health check doing.")
	return &pb.HealthResponse{
		Status: "success",
	}, nil
}
