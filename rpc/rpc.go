package rpc

import (
	"log"
	"net"

	pb "github.com/colinrs/ffly-plus/internal/proto"

	"github.com/colinrs/pkgx/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// InitRPCService ...
func InitRPCService() error {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("fail to listen")
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	reflection.Register(s)
	logger.Info("start listen:%s", port)
	if err := s.Serve(lis); err != nil {
		logger.Error("fail to server:%#v", err)
		return err
	}
	return nil
}
