package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/ent"
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
		logger.Error("%s", err.Error())
		os.Exit(1)
		return
	}

	cl, err := ent.Open("mysql", "root:root@tcp(mysql:3306)/attend")
	if err != nil {
		logger.Error("%s", err.Error())
		os.Exit(1)
		return
	}
	defer cl.Close()
	repof := infra.NewFactory(logger, cl)

	s := grpc.NewServer(grpc.UnaryInterceptor(LoggingInterceptor(logger)))
	Register(s, logger, repof)

	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Println(err)
	}

	logger.Info("Exit gRPC Server")
	s.GracefulStop()
}

// Register サーバーの登録
func Register(s *grpc.Server, l core.Logger, factory *infra.Factory) {
	pb.RegisterCheckServer(s, server.NewCheck(l))
	tr := factory.CreateTransaction()
	ur := factory.CreateUserRepository()
	us := service.NewUser(l, tr, ur)
	pb.RegisterUserServer(s, server.NewUser(l, us))
}
