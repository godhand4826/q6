package main

import (
	"net/http"
	"q6/lib/logger"
	"q6/lib/serve"
	"q6/pkg/config"
	"q6/pkg/router"
	"q6/pkg/service"
	"time"

	"go.uber.org/zap"
)

// @title		Q6 API
// @version	1.0
// @host		localhost:8080
// @BasePath	/
func main() {
	cfg := config.NewConfig()

	logger := logger.NewLogger()
	defer func() {
		_ = logger.Sync()
	}()

	matchingSystem := service.NewMatchingSystem(logger)
	engine := serve.NewServer(logger)
	router.NewMatchingSystemRouter(matchingSystem).BindOn(engine)
	serve.NewSwaggerRouter().BindOn(engine)

	server := &http.Server{
		ReadTimeout: time.Second,
		Addr:        cfg.Addr,
		Handler:     engine.Handler(),
	}

	logger.Info("Starting server at", zap.String("addr", cfg.Addr))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("Server unexpected stopped", zap.Error(err))
	}
}
