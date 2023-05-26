//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"likyam.cn/src/common/client"
	"likyam.cn/src/common/config"
	"likyam.cn/src/common/server"
	"likyam.cn/src/order/service/v1/dao"

	serverV1 "likyam.cn/src/order/service/v1/server"
)

func InitServer(cfg string) (*Server, error) {

	wire.Build(
		// run server
		NewServer,
		// server
		serverV1.NewServer,
		// component
		dao.NewRepository,
		config.NewConfig,
		server.NewRedis,
		server.NewMysqlDB,
		NewHttpServer,
		NewGrpcServer,
		server.NewRunGroup,
		server.NewLogger,
		//server.NewTrace,
		// client
		client.NewUserClient,
	)

	return &Server{}, nil

}
