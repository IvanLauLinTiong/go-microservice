package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router 	http.Handler
	rdb 	*redis.Client
	config 	Config
}

func New(config Config) *App {
	app := &App {
		rdb: redis.NewClient(&redis.Options{
			Addr: config.RedisAddress,
		}),
		config: config,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}

	// := means initializayion + assignment
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis server: %w", err) //error unwrapping with %w
	}

	defer func()  {
		// fmt.Println("closing redis client...")
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis: ", err)
		}
	}()

	fmt.Println("Starting server...")

	ch := make(chan error, 1)

	// goroutine
	go func() {
		err := server.ListenAndServe() // blocking call
		if err != nil {
			ch <- fmt.Errorf("failed to listen server: %w", err) //error wrapping %w
		}
		close(ch)
	}()

	//blocking call
	//  ctx.Done()

	// blocking main execution
	// err, channelOpen := <-ch
	// if !channelOpen {
	// 	// channel was closed
	// }

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		// fmt.Println("done...")
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		// defer func() { fmt.Println("cancel 2..."); cancel()}()
		defer cancel()
		return server.Shutdown(timeout)
	}
}