package rpc

import (
	"context"
	"encoding/json"

	"ffly-plus/internal/code"
	pb "ffly-plus/internal/proto"
	"ffly-plus/service"

	"github.com/colinrs/pkgx/logger"
)

type server struct{}

func (s *server) UserRegister(ctx context.Context, req *pb.UserRegisterRequest) (*pb.RequestReplay, error) {

	logger.Info("####### get client UserRegister requests: %s", req.Nickname)
	return &pb.RequestReplay{Code: 0, Message: "OK", Data: []byte("UserRegister")}, nil
}

func (s *server) DeletetUser(ctx context.Context, req *pb.DeletetUserRequest) (*pb.RequestReplay, error) {

	logger.Info("####### get client DeletetUser requests: %s", req.UserName)
	return &pb.RequestReplay{Code: 0, Message: "OK", Data: []byte("DeletetUser")}, nil
}

func (s *server) UserUpdate(ctx context.Context, req *pb.UserUpdateRequest) (*pb.RequestReplay, error) {

	logger.Info("####### get client UserUpdate requests: %s", req.UserName)
	return &pb.RequestReplay{Code: 0, Message: "OK", Data: []byte("UserUpdate")}, nil
}

func (s *server) SelectUser(ctx context.Context, req *pb.SelectUserRequest) (*pb.RequestReplay, error) {
	query := map[string]interface{}{
		"id":        req.UserId,
		"user_name": req.UserName,
	}
	user, err := service.SelectUser(query)
	if err != nil {
		return &pb.RequestReplay{Code: 1, Message: err.Error(), Data: nil}, nil
	}
	logger.Info("####### get client SelectUser requests: %s", req.UserName)
	bytes, err := json.Marshal(user)
	if err != nil {
		return &pb.RequestReplay{Code: 1, Message: err.Error(), Data: nil}, nil
	}
	return &pb.RequestReplay{Code: int64(code.OK.Code), Message: code.OK.Message, Data: bytes}, nil
}
