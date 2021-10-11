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
const TokenExpiredTime = time.Hour * 24 // expired time: 24 hr

type Message struct {
}

type Token struct {
	token string
	uid   string
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

func (rdc *RedisConn) PingRedis(ctx context.Context) (err error) {
	// Try connect redis
	_, err = rdc.rdb.Ping(ctx).Result()
	if err != nil {
		return errors.Wrapf(err, "failed to ping redis")
	}
	return nil
}

// Scan function use to scan all the keys in redis, also can use to get value
func (rdc *RedisConn) Scan(ctx context.Context, pattern string, count int64) (keys []string, err error) {
	var cursor uint64
	keyNum := 0
	// Try scan keys
	for {
		keys, cursor, err = rdc.rdb.Scan(ctx, cursor, pattern, count).Result()
		if err != nil {
			return nil, errors.Wrapf(err, "redis scan keys error")
		}
		keyNum += len(keys)
		fmt.Printf("found %v keys\n", keyNum)
		if cursor == 0 {
			// TODO: make log after production
			break
		}
	}
	return keys, nil
}

func (rdc *RedisConn) SetToken(ctx context.Context, uid string, token string) (result string, err error) {
	result, err = rdc.rdb.Set(ctx, TokenForward+token, uid, TokenExpiredTime).Result()
	if err != nil {
		return result, errors.Wrapf(err, "Set token error")
	}
	return result, nil
}

// ValidateToken validate user's token from redis cache
func (rdc *RedisConn) ValidateToken(ctx context.Context, uid string, token string) (isValid bool, err error) {
	uid_, err := rdc.rdb.Get(ctx, TokenForward+token).Result()
	if err != nil {
		return false, errors.Wrapf(err, "Set token error")
	}
	if uid_ != uid { // invalid token
		return false, nil
	}
	return true, nil // valid token
}

// GetGroups search groups relation is in redis
func (rdc *RedisConn) GetGroups() {

}

// SetGroups cache groups relation from sql
func (rdc *RedisConn) SetGroups() {

}
