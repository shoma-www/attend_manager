package infra

import (
	"errors"
	"fmt"

	"github.com/shoma-www/attend_manager/api/config"
	"github.com/shoma-www/attend_manager/api/service"
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
func NewRepoFactory(c *config.Config) *RepoFactory {
	conf := c.Client.Grpc
	return &RepoFactory{
		address: fmt.Sprintf("%s:%d", conf.Addr, conf.Port),
	}
}

// CreateCheckRepository Repositoryつくるで
func (rf *RepoFactory) CreateCheckRepository() service.CheckRepository {
	return &checkGrpc{address: rf.address}
}
