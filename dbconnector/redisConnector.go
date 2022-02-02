package dbconnector

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisConn struct {
	Rdb *redis.Client
	Ctx context.Context
}

func (r *RedisConn) Connect() {
	r.Ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Ping(r.Ctx).Err()
	if err != nil {
		log.Fatal(fmt.Errorf("error in Redis ping: %v", err))
	}

	r.Rdb = rdb
}
