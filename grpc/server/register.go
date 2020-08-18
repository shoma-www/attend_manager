package server

import (
	"github.com/shoma-www/attend_manager/core"
	pb "github.com/shoma-www/attend_manager/grpc/proto"
	"google.golang.org/grpc"
)

// Register サーバーの登録
func Register(s *grpc.Server, l core.Logger) {
	pb.RegisterCheckServer(s, NewCheckServer(l))
}
