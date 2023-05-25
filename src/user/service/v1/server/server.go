package serverV1

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	userPBV1 "likyam.cn/api/gen/go/proto/user/v1"
	"likyam.cn/src/user/service/v1/dao"
)

// Server Server struct
type Server struct {
	userPBV1.UnimplementedUserServiceServer
	repo   *dao.Repository
	logger log.Logger
}

// NewServer New service grpc server
func NewServer(
	repo *dao.Repository,
	logger log.Logger,
) userPBV1.UserServiceServer {
	return &Server{
		repo:   repo,
		logger: logger,
	}
}

func (s *Server) GetUser(ctx context.Context, re *userPBV1.GetUserRequest) (*userPBV1.GetUserResponse, error) {
	fmt.Println(re.GetUserId())
	userData, err := s.repo.Info(re.GetUserId())
	if err != nil {
		return nil, err
	}

	userPb := &userPBV1.User{
		UserId:        userData.UserID,
		UserNo:        userData.UserNo,
		UserName:      userData.UserName,
		Email:         userData.Email,
		PhoneNumber:   userData.PhoneNumber,
		Avatar:        userData.Avatar,
		Unionid:       userData.UnionID,
		Openid:        userData.OpenID,
		MiniappOpenid: userData.MiniappOpenID,
		AddTime:       timestamppb.New(userData.AddTime),
	}

	return &userPBV1.GetUserResponse{
		User: userPb,
	}, nil
}
