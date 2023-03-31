package redis

import (
	"github.com/bav-demo/micro1/config"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Rdb *redis.Client
	c   *config.Config
}

func (r *Redis) Connect() {
	r.c = &config.Config{}

	c, _ := r.c.NewConfig()
	r.Rdb = redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password, // no password set
		DB:       int(c.Redis.DB),  // use default DB
	})
}
