package redis

import (
	"context"
	"fmt"
	"log/slog"

	"main_service/internal/config"
	"main_service/internal/redis/model"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	*model.Model
}

func New(conf *config.RedisConfig, logger *slog.Logger) (*Redis, error) {
	db, err := NewDB(conf)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	_, err = db.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{
		Model: model.New(conf, logger.With(slog.String("module", "model")), db),
	}, nil
}

func NewDB(conf *config.RedisConfig) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	db := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password,
		DB:       conf.Number,
	})
	return db, nil
}
