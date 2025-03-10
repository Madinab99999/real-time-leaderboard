package main

import (
	"context"
	"log/slog"
	"notification/internal/config"
	redis "notification/internal/redis"
	"notification/internal/redis/model/leaderboard"
	"notification/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	conf := config.New(ctx)

	// logger
	log := logger.New(logger.Options{
		Format:     logger.TextFormat,
		Level:      slog.LevelDebug,
		AddSource:  false,
		TimeFormat: time.DateTime,
	})

	//redis
	redis, err := redis.New(conf.REDIS, log.With(slog.String("service", "redis")))
	if err != nil {
		log.ErrorContext(ctx, "failed to start redis", slog.Any("error", err))
		panic(err)
	}
	defer redis.Close()
	log.InfoContext(ctx, "successfully connected to Redis")

	leaderboard := leaderboard.New(conf.REDIS, log, redis.Client)
	if err := leaderboard.StartNotification(ctx); err != nil {
		log.ErrorContext(ctx, "notification service stopped with error", slog.Any("error", err))
		os.Exit(1)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		slog.Info("received shutdown signal", "signal", sig)
		cancel()
	}()
	<-ctx.Done()
	log.InfoContext(ctx, "notification service shutdown complete")
}
