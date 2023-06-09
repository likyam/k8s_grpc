package client

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	userPBV1 "likyam.cn/api/gen/go/proto/user/v1"
	"likyam.cn/src/common/config"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewUserClient 实例化 user 客户端
func NewUserClient(conf *config.Config) (userPBV1.UserServiceClient, error) {

	serverAddress := conf.Client.UserHost + conf.Client.UserPort
	fmt.Println(serverAddress)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		return nil, err
	}

	client := userPBV1.NewUserServiceClient(conn)
	return client, nil

}
