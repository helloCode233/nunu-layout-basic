//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-basic/internal/handler"
	"github.com/go-nunu/nunu-layout-basic/internal/repository"
	"github.com/go-nunu/nunu-layout-basic/internal/router"
	"github.com/go-nunu/nunu-layout-basic/internal/service"
	"github.com/go-nunu/nunu-layout-basic/pkg/config"
	"github.com/go-nunu/nunu-layout-basic/pkg/log"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(
	router.NewServerHTTP,
)

var RepositorySet = wire.NewSet(

	repository.NewRepository,
	repository.NewUserRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

func NewWire(*config.Configuration, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServiceSet,
		RepositorySet,
		RouterSet,
		HandlerSet,
	))
}
