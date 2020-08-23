package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/infra"
	"github.com/shoma-www/attend_manager/grpc/server"
	"github.com/shoma-www/attend_manager/grpc/service"

	_ "github.com/go-sql-driver/mysql"
	pb "github.com/shoma-www/attend_manager/grpc/proto"
)

const (
	port = ":50051"
)

func main() {
	logger := core.NewLogger(core.Debug)
	logger.Info("Start gRPC Server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}

	repof := infra.NewRepoFactory(logger)

	s := grpc.NewServer(grpc.UnaryInterceptor(LoggingInterceptor(logger)))
	Register(s, logger, repof)

	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Println(err)
	}

	logger.Info("Exit gRPC Server")
	s.Stop()
}

// Register サーバーの登録
func Register(s *grpc.Server, l core.Logger, repof *infra.RepoFactory) {
	pb.RegisterCheckServer(s, server.NewCheck(l))
	ur := repof.CreateUserRepository()
	us := service.NewUser(l, ur)
	pb.RegisterUserServer(s, server.NewUser(l, us))
}
