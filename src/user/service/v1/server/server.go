package serverV1

import (
	"context"
	"github.com/go-kit/log"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
	userPBV1 "likyam.cn/api/gen/go/proto/user/v1"
	"likyam.cn/src/user/service/v1/dao"
)

// Server Server struct
type Server struct {
	userPBV1.UnimplementedUserServiceServer
	repo   *dao.Repository
	logger log.Logger
	trace  *sdktrace.TracerProvider
}

// NewServer New service grpc server
func NewServer(
	repo *dao.Repository,
	logger log.Logger,
	trace *sdktrace.TracerProvider,
) userPBV1.UserServiceServer {
	return &Server{
		repo:   repo,
		logger: logger,
		trace:  trace,
	}
}

func (s *Server) GetUser(ctx context.Context, re *userPBV1.GetUserRequest) (*userPBV1.GetUserResponse, error) {
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
