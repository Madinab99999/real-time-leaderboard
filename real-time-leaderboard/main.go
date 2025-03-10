package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"real-time-leaderboard/internal/config"
	"real-time-leaderboard/internal/redis"
	"real-time-leaderboard/internal/redis/model/leaderboard"
	"real-time-leaderboard/pkg/logger"
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

	leaderboard := leaderboard.New(conf.REDIS, log, redis.Client)
	log.InfoContext(ctx, "Real Time Global Leaderboard service started")
	leaderboard.PrintLeaderboard()
	go leaderboard.SubscribeLeaderboardUpdates()
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
