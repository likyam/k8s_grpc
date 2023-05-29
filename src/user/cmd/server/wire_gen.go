// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"likyam.cn/src/common/config"
	"likyam.cn/src/common/server"
	"likyam.cn/src/user/service/v1/dao"
	"likyam.cn/src/user/service/v1/server"
)

// Injectors from wire.go:

func InitServer(cfg string) (*Server, error) {
	configConfig := config.NewConfig(cfg)
	db := server.NewMysqlDB(configConfig)
	client := server.NewRedis(configConfig)
	repository := dao.NewRepository(db, client)
	group := server.NewRunGroup()
	logger := server.NewLogger()
	httpServer := NewHttpServer(configConfig)
	userServiceServer := serverV1.NewServer(repository, logger)
	grpcServer := NewGrpcServer(logger, userServiceServer)
	serverServer := NewServer(repository, configConfig, group, logger, httpServer, grpcServer, db)
	return serverServer, nil
}