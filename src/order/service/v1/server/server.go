package serverV1

import (
	"context"
	"errors"
	"github.com/go-kit/log"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
	orderPBV1 "likyam.cn/api/gen/go/proto/order/v1"
	userPBV1 "likyam.cn/api/gen/go/proto/user/v1"
	"likyam.cn/src/common/config"
	"likyam.cn/src/order/service/v1/dao"
)

// Server Server struct
type Server struct {
	orderPBV1.UnimplementedOrderServiceServer
	conf   *config.Config
	repo   *dao.Repository
	logger log.Logger
	trace  *sdktrace.TracerProvider

	userClient userPBV1.UserServiceClient
}

// NewServer New service grpc server
func NewServer(
	repo *dao.Repository,
	logger log.Logger,
	conf *config.Config,
	trace *sdktrace.TracerProvider,

	userClient userPBV1.UserServiceClient,
) orderPBV1.OrderServiceServer {
	return &Server{
		conf:   conf,
		repo:   repo,
		logger: logger,
		trace:  trace,

		userClient: userClient,
	}
}

func (s *Server) GetOrderById(ctx context.Context, re *orderPBV1.GetOrderByIdRequest) (*orderPBV1.GetOrderByIdResponse, error) {
	orderData, err := s.repo.Info(re.GetOrderId())
	if err != nil {
		return nil, err
	}
	user, err := s.userClient.GetUser(ctx, &userPBV1.GetUserRequest{UserId: re.GetUserId()})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	order := &orderPBV1.Order{
		OrderNo:    orderData.OrderNo,
		UserName:   orderData.UserName,
		UserId:     orderData.UserId,
		Remark:     orderData.Remark,
		PricePay:   orderData.PricePay,
		PriceTotal: orderData.PriceTotal,
		PricePaid:  orderData.PricePaid,
		IsInvoice:  orderData.IsInvoice,
		PayTime:    timestamppb.New(orderData.PayTime),
		AddTime:    timestamppb.New(orderData.AddTime),
	}

	return &orderPBV1.GetOrderByIdResponse{
		Order: order,
	}, nil
}
