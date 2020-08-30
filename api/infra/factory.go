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

// CreateCheckRepository Repositoryつくるで
func (rf *Factory) CreateCheckRepository() service.CheckRepository {
	return &checkGrpc{address: rf.address}
}
