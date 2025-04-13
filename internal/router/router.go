package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-basic/internal/handler"
	"github.com/go-nunu/nunu-layout-basic/internal/middleware"
	"github.com/go-nunu/nunu-layout-basic/pkg/config"
	"github.com/go-nunu/nunu-layout-basic/pkg/helper/path"
	"github.com/go-nunu/nunu-layout-basic/pkg/log"
	"path/filepath"
)

// @wire:Server
func NewServerHTTP(
	logger *log.Logger,
	configuration config.Configuration,
	userHandler *handler.UserHandler,
) *gin.Engine {
	switch configuration.App.Env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	}
	logger.Info("yaml start with" + configuration.App.Env)
	router := gin.Default()
	router.Use(
		middleware.CORS(),
	)
	rootDir := path.RootPath()
	// 前端项目静态资源
	router.StaticFile("/", filepath.Join(rootDir, "static/dist/index.html"))
	router.Static("/assets", filepath.Join(rootDir, "static/dist/assets"))
	router.StaticFile("/favicon.ico", filepath.Join(rootDir, "static/dist/favicon.ico"))
	// 其他静态资源
	router.Static("/public", filepath.Join(rootDir, "static"))
	router.Static("/storage", filepath.Join(rootDir, "storage/app/public"))

	return router
}
