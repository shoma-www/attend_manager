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

// Factory Repositoryのファクトリ
type Factory struct {
	address string
}

// NewFactory コンストラクタ
func NewFactory(c *config.Config) *Factory {
	conf := c.Client.Grpc
	return &Factory{
		address: fmt.Sprintf("%s:%d", conf.Addr, conf.Port),
	}
}

// CreateCheck Repositoryつくるで
func (rf *Factory) CreateCheck() service.CheckRepository {
	return &checkGrpc{address: rf.address}
}

// CreateUser Repositoryつくるで
func (rf *Factory) CreateUser() service.UserRepository {
	return &userGrpc{address: rf.address}
}

// CreateGroup Repositoryつくるで
func (rf *Factory) CreateGroup() service.GroupRepository {
	return &groupGrpc{address: rf.address}
}

// CreateSession Sessionつくるで
func (rf *Factory) CreateSession() service.Session {
	return service.NewRedisSession("redis:6379", "", 0)
}
