package main

import (
	"context"
	"fmt"
	"im_test/data"
	"time"
)

func testRedis() {
	var ctx, canclefunc = context.WithDeadline(context.TODO(), time.Now().Add(time.Second))
	rdc := data.NewRedisConn("localhost:6379", "", 0)
	// ping redis db
	err := rdc.PingRedis(ctx)
	if err != nil {
		fmt.Print(err)
		canclefunc()
		return
	}
	// test "set" command
	result, err := rdc.SetToken(ctx, "123456", "11WADE1TEST")
	if err != nil {
		fmt.Print(err)
		canclefunc()
		return
	}
	// test "scan" command
	keys, err := rdc.Scan(ctx, "token_*", 10)
	if err != nil {
		fmt.Println(err)
		canclefunc()
		return
	}
	// print result
	fmt.Println(keys)
	fmt.Println(result)
	canclefunc()
}

func testMysql() {

}

func main() {
	testRedis()
	testMysql()
}
