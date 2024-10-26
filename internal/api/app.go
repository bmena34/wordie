package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "redis:6379" // Default to IPv4 loopback address
	}

	app := &App{
		router: loadRoutes(),
		rdb: redis.NewClient(&redis.Options{
			Addr: redisAddr,
			DB:   0,
		}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping().Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		err := a.rdb.Close()
		if err != nil {
			fmt.Println("error closing redis connection:", err)
		}
	}()

	fmt.Println("server Starting")

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("server failed to start: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
