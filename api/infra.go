package main

import (
	"fmt"

	pb "github.com/shoma-www/attend_manager/api/proto"
	"google.golang.org/grpc"
)

// Grpc Grpcサーバへのクライアントを生成する構造体
type Grpc struct {
	address string
	cc      *grpc.ClientConn
}

// NewGrpc コンストラクタ
func NewGrpc(c *Config) *Grpc {
	conf := c.Client.Grpc
	return &Grpc{
		address: fmt.Sprintf("%s:%d", conf.Addr, conf.Port),
	}
}

func (g *Grpc) connection() *grpc.ClientConn {
	if g.cc == nil {
		con, err := grpc.Dial(g.address, grpc.WithInsecure())
		if err != nil {
			panic("Grpc dial failed!")
		}
		g.cc = con
	}
	return g.cc
}

// Close ClientのConnectionを閉じる
func (g *Grpc) Close() {
	if g.cc != nil {
		g.cc.Close()
		g.cc = nil
	}
}

// CreateCheckClient CheckServerへのClietnを生成
func (g *Grpc) CreateCheckClient() pb.CheckClient {
	return pb.NewCheckClient(g.connection())
}
