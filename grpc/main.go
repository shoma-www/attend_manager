package main

import (
	"log"
	"net"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/config"
	"github.com/shoma-www/attend_manager/grpc/ent"
	"github.com/shoma-www/attend_manager/grpc/infra"
	pb "github.com/shoma-www/attend_manager/grpc/proto"
	"github.com/shoma-www/attend_manager/grpc/server"
	"github.com/shoma-www/attend_manager/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	path := "./config/config.yaml"
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c, err := config.LoadConfig(absPath)
	if err != nil {
		log.Fatalln(err)
		return
	}

	logger := core.NewLogger(core.Debug)
	logger.Info("Start gRPC Server")
	lis, err := net.Listen("tcp", c.Server.Addr)
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
	rFactory := infra.NewFactory(logger, cl)

	s := grpc.NewServer(grpc.UnaryInterceptor(LoggingInterceptor(logger)))
	Register(s, logger, rFactory)

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

	gr := factory.CreateAttendanceGroupRepository()
	gs := service.NewAttendanceGroup(l, tr, gr, ur)
	pb.RegisterAttendanceGroupServer(s, server.NewAttendanceGroup(l, gs))
}
