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

func testFileRename(path string) { //  T<60µs per call
	hex, key, err := data.ReadFileBytes(path)
	if err != nil {
		panic(err)
	}
	hash256, err := data.GetHash256(hex, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(hash256)
}

const testPath = "D:\\1.docx"

func main() {
	for i := 0; i < 10; i++ {
		st := time.Now()
		testFileRename(testPath)
		fmt.Println("once call used: ", time.Since(st)/time.Nanosecond)
		time.Sleep(1)
	}
}
