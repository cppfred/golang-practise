package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"time"
)

// this is go file for test whole process with redis

const TokenForward = "token_"

type Message struct {
}

type Token struct {
}

type RedisConn struct {
	rdb *redis.Client
}

func NewRedisConn(addr string, psw string, db int) *RedisConn {
	return &RedisConn{
		rdb: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: psw, // no password set
			DB:       db,  // use default DB
		}),
	}
}

func PingRedis(ctx context.Context, redisConn *RedisConn) (err error) {
	// Try connect redis
	_, err = redisConn.rdb.Ping(ctx).Result()
	if err != nil {
		errors.Wrap(err, "failed to ping redis")
	}
	return nil
}

// Scan function use to scan all the keys in redis, also can use to get value
func Scan(ctx context.Context, rdc *RedisConn, pattern string, count int64) (keys []string, err error) {
	var cursor uint64
	key_n := 0
	// Try scan keys
	for {
		keys, cursor, err = rdc.rdb.Scan(ctx, cursor, pattern, count).Result()
		if err != nil {
			return
		}
		key_n += len(keys)
		fmt.Printf("found %v keys\n", key_n)
		if cursor == 0 {
			break
		}
	}
	return
}

func SetToken(ctx context.Context, rdc *RedisConn, uid string, token string) (result string, err error) {
	result, err = rdc.rdb.Set(ctx, TokenForward+uid, token, 3*time.Hour).Result()
	if err != nil {
		return
	}
	return
}
