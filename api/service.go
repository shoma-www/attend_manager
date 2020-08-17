package main

import (
	"context"

	pb "github.com/shoma-www/attend_manager/api/proto"
	"github.com/shoma-www/attend_manager/core"
)

// HealthCheckService 他システムへのhealthcheck
type HealthCheckService struct {
	logger *core.Logger
	client pb.CheckClient
}

// NewHealthCheckService constructor
func NewHealthCheckService(l *core.Logger, cl pb.CheckClient) HealthCheckService {
	return HealthCheckService{logger: l, client: cl}
}

// Grpc GrpcのサーバーにHealthCheckを実施
func (hs HealthCheckService) Grpc(ctx context.Context) error {
	res, err := hs.client.HealthCheck(ctx, &pb.HealthRequest{})
	hs.logger.Debug("health check grpc status: %s", res.GetStatus())
	return err
}
