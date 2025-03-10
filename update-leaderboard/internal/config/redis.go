package config

import (
	"context"
	"flag"
	"os"
	"strconv"
)

type RedisConfig struct {
	Host                     string
	Port                     int
	Number                   int
	User                     string
	Password                 string
	RedisChannelUpdate       string
	RedisChannelConsole      string
	RedisChannelNotification string
}

func newRedisConfig(_ context.Context) *RedisConfig {
	port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	number, _ := strconv.Atoi(os.Getenv("REDIS_NUMBER"))
	c := &RedisConfig{
		Host:                     os.Getenv("REDIS_HOST"),
		Port:                     port,
		Number:                   number,
		User:                     os.Getenv("REDIS_USER"),
		Password:                 os.Getenv("REDIS_PASSWORD"),
		RedisChannelUpdate:       os.Getenv("REDIS_CHANNEL_UPDATE"),
		RedisChannelNotification: os.Getenv("REDIS_CHANNEL_NOTIFICATION"),
		RedisChannelConsole:      os.Getenv("REDIS_CHANNEL_CONSOLE"),
	}

	flag.StringVar(&c.Host, "redis-host", c.Host, "redis host [REDIS_HOST]")
	flag.IntVar(&c.Port, "redis-port", c.Port, "redis port [REDIS_PORT]")
	flag.IntVar(&c.Number, "redis-number", c.Number, "redis name [REDIS_NUMBER]")
	flag.StringVar(&c.User, "redis-user", c.User, "redis user [REDIS_USER]")
	flag.StringVar(&c.Password, "redis-password", c.Password, "redis password [REDIS_PASSWORD]")
	flag.StringVar(&c.Host, "redis-channel-update", c.Host, "redis channel for get new updates  [REDIS_CHANNEL_UPDATE]")
	flag.StringVar(&c.Host, "redis-channel-console", c.Host, "redis channel for update leaderboard console  [REDIS_CHANNEL_CONSOLE]")
	flag.StringVar(&c.Host, "redis-channel-notification", c.Host, "redis channel for notification [REDIS_CHANNEL_NOTIFICATION]")
	return c
}
