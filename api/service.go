package main

import (
	"context"

	pb "github.com/shoma-www/attend_manager/api/proto"
	"github.com/shoma-www/attend_manager/core"
	"google.golang.org/grpc"
)

// HealthCheckService 他システムへのhealthcheck
type HealthCheckService struct {
	logger *core.Logger
}

// NewHealthCheckService constructor
func NewHealthCheckService(l *core.Logger) HealthCheckService {
	return HealthCheckService{logger: l}
}

// Grpc GrpcのサーバーにHealthCheckを実施
func (hs HealthCheckService) Grpc(ctx context.Context) error {
	conn, err := grpc.Dial("grpc:50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewCheckClient(conn)
	res, err := client.HealthCheck(ctx, &pb.HealthRequest{})
	hs.logger.Debug("health check grpc status: %s", res.GetStatus())
	return err
}
