// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"c1/internal/biz"
	"c1/internal/conf"
	"c1/internal/data"
	"c1/internal/server"
	"c1/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	c2Repo := data.NewC2Repo(dataData, logger)
	c1Usecase := biz.NewC1Usecase(c2Repo, logger)
	c1Service := service.NewC1Service(c1Usecase, logger)
	grpcServer := server.NewGRPCServer(confServer, c1Service, logger)
	httpServer := server.NewHTTPServer(confServer, c1Service, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
