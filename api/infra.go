package main

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/shoma-www/attend_manager/api/proto"
	"google.golang.org/grpc"
)

// createGrpcConn Connectioを生成する
func createGrpcConn(address string) (*grpc.ClientConn, error) {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("grpc dial failed")
	}
	return con, nil
}

// RepoFactory Repositoryのファクトリ
type RepoFactory struct {
	address string
}

// NewRepoFactory コンストラクタ
func NewRepoFactory(c *Config) *RepoFactory {
	conf := c.Client.Grpc
	return &RepoFactory{
		address: fmt.Sprintf("%s:%d", conf.Addr, conf.Port),
	}
}

// CreateCheckRepository Repositoryつくるで
func (rf *RepoFactory) CreateCheckRepository() CheckRepository {
	return &checkGrpc{address: rf.address}
}

// checkGrpc Check系のGrpc通信するやつ
type checkGrpc struct {
	address string
}

// HealthCheck CheckServerへのClietnを生成
func (cg *checkGrpc) HealthCheck(ctx context.Context) (*HealthCheckStatus, error) {
	con, err := createGrpcConn(cg.address)
	if err != nil {
		return nil, err
	}
	defer con.Close()
	client := pb.NewCheckClient(con)
	pbst, err := client.HealthCheck(ctx, &pb.HealthRequest{})
	if err != nil {
		return nil, err
	}
	st := &HealthCheckStatus{
		Status: pbst.GetStatus(),
	}
	return st, nil
}
