package main

import (
	"context"
	"log"

	pb "github.com/shoma-www/attend_manager/api/proto"
	"google.golang.org/grpc"
)

// HealthCheckService 他システムへのhealthcheck
type HealthCheckService struct {
}

// NewHealthCheckService constructor
func NewHealthCheckService() HealthCheckService {
	return HealthCheckService{}
}

// Grpc GrpcのサーバーにHealthCheckを実施
func (hs HealthCheckService) Grpc(ctx context.Context) error {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewCheckClient(conn)
	res, err := client.HealthCheck(ctx, &pb.HealthRequest{})
	log.Printf("health check grpc status: %s", res.GetStatus())
	return err
}
