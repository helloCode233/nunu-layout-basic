// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-basic/internal/handler"
	"github.com/go-nunu/nunu-layout-basic/internal/repository"
	"github.com/go-nunu/nunu-layout-basic/internal/server"
	"github.com/go-nunu/nunu-layout-basic/internal/service"
	"github.com/go-nunu/nunu-layout-basic/pkg/config"
	"github.com/go-nunu/nunu-layout-basic/pkg/log"
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewWire(config *config.Configuration, logger *log.Logger) (*gin.Engine, func(), error) {
	handlerHandler := handler.NewHandler(logger)
	serviceService := service.NewService(logger)
	db := repository.NewDb()
	repositoryRepository := repository.NewRepository(logger, db)
	userRepository := repository.NewUserRepository(repositoryRepository)
	userService := service.NewUserService(serviceService, userRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	engine := server.NewServerHTTP(logger, userHandler)
	return engine, func() {
	}, nil
}

// wire.go:

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(repository.NewDb, repository.NewRepository, repository.NewUserRepository)

var ServiceSet = wire.NewSet(service.NewService, service.NewUserService)

var HandlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler)
