package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
	"update-leaderboard/internal/config"
	redis "update-leaderboard/internal/redis"
	"update-leaderboard/internal/redis/model/leaderboard"
	"update-leaderboard/pkg/logger"
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
	//leaderboard.UpdateLeaderboard(ctx)
	if err := leaderboard.StartUpdate(ctx); err != nil {
		log.ErrorContext(ctx, "leaderboard service stopped with error", slog.Any("error", err))
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
	log.InfoContext(ctx, "leaderboard service shutdown complete")
}
