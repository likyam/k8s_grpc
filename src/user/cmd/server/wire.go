//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"likyam.cn/src/common/config"
	"likyam.cn/src/common/server"
	"likyam.cn/src/user/service/v1/dao"

	serverV1 "likyam.cn/src/user/service/v1/server"
)

func InitServer(cfg string) (*Server, error) {

	wire.Build(
		NewServer,
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
	)

	return &Server{}, nil

}
