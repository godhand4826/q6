package main

import (
	"net/http"
	"os"

	"cmp"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

func main() {
	addr := cmp.Or(os.Getenv("ADDR"), ":8080")

	logger := NewLogger()
	defer logger.Sync()

	matchingSystem := NewMatchingSystem(logger)
	engine := NewServer(logger)
	NewMatchingSystemRouter(matchingSystem).BindOn(engine)

	server := &http.Server{
		Addr:    addr,
		Handler: engine.Handler(),
	}

	logger.Info("Starting server at", zap.String("addr", addr))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("Server unexpected stopped", zap.Error(err))
	}
}
