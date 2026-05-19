package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
	config Config
}

func New(config Config) *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{Addr: config.RedisAddress}),
		config: config,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}

	log.Printf("listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil

}
