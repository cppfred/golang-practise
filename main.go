package main

import (
	"context"
	"fmt"
	"im_test/data"
	"time"
)

func main() {
	var ctx, canclefunc = context.WithDeadline(context.TODO(), time.Now().Add(time.Second))
	rdc := data.NewRedisConn("localhost:6379", "", 0)
	err := data.PingRedis(ctx, rdc)
	if err != nil {
		fmt.Print(err)
		canclefunc()
		return
	}
	result, err := data.SetToken(ctx, rdc, "123456", "11WADE1TEST")
	if err != nil {
		fmt.Print(err)
		canclefunc()
		return
	}
	keys, err := data.Scan(ctx, rdc, "token_*", 10)
	if err != nil {
		fmt.Println(err)
		canclefunc()
		return
	}
	fmt.Println(keys)
	fmt.Println(result)
	canclefunc()
}
