package redis

import (
	"context"
	"example/env"

	"github.com/go-redis/redis/v8"
)

type handler struct {
	c *redis.Client
}

var ctx = context.Background()

func NewRedis(cfg *env.RedisConfig) Manager {
	h := &handler{}
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,     //"localhost:6379", //config.Addr,
		Password: cfg.Password, //setconfig.Password,//"",               // no password setconfig.Password,
		DB:       cfg.DB,       //0,                // use default DBconfig.DB,
	})
	h.c = cli
	return h
}

func (h handler) SetToken(playerID string, token string) error {
	err := h.c.Set(ctx, playerID, token, 0).Err()
	return err

}
func (h handler) GetToken(playerID string) (string, error) {
	val, err := h.c.Get(ctx, playerID).Result()
	return val, err
}
