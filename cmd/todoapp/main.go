package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/dyingvoid/todoapp/internal/core/logger"
	core_postgres_pool "github.com/dyingvoid/todoapp/internal/core/repository/postgres/conn"
	core_http_middleware "github.com/dyingvoid/todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/dyingvoid/todoapp/internal/core/transport/http/server"
	web_fs_repository "github.com/dyingvoid/todoapp/internal/features/repository/file_system"
	web_service "github.com/dyingvoid/todoapp/internal/features/service"
	web_transport_http "github.com/dyingvoid/todoapp/internal/features/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing postgres connection pool")
	pool, err := core_postgres_pool.NewConnectionPool(
		ctx,
		core_postgres_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "web"))
	webRepository := web_fs_repository.NewWebRepository()
	webService := web_service.NewWebService(webRepository)
	webTransportHTTP := web_transport_http.NewWebHTTPHandler(webService)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes( /* users */ )
	httpServer.RegisterAPIRoutes(apiVersionRouter)
	httpServer.RegisterRoutes(webTransportHTTP.Routes()...)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
