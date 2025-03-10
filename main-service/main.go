package main

import (
	"context"
	"fmt"
	"log/slog"
	"main_service/internal/config"
	"main_service/internal/governor"
	"main_service/internal/redis"
	"main_service/internal/rest"
	"main_service/pkg/logger"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// config
	conf := config.New(ctx)

	// logger
	log := logger.New(logger.Options{
		Format:     logger.TextFormat,
		Level:      slog.LevelDebug,
		AddSource:  false,
		TimeFormat: time.DateTime,
	})
	//governor
	gov := governor.New(conf)

	//redis
	redis, err := redis.New(conf.REDIS, log.With(slog.String("service", "redis")))
	if err != nil {
		log.ErrorContext(ctx, "failed to start redis", slog.Any("error", err))
		panic(err)
	}

	//rest
	r := rest.New(conf.API.Rest, log.With(slog.String("service", "rest")), gov)
	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		if err := r.Start(ctx); err != nil {
			log.ErrorContext(ctx, "failed to start rest", slog.Any("error", err))
		}
		cancelFunc()
	}(ctx, cancel)

	gov.Config(ctx, conf, log.With(slog.String("service", "governor")), redis)

	go func(cancelFunc context.CancelFunc) {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)

		sig := <-shutdown
		log.WarnContext(ctx, "signal received - shutting down...", slog.Any("signal", sig))

		cancelFunc()
	}(cancel)

	<-ctx.Done()

	fmt.Println("shutting down gracefully")
}
