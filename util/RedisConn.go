package util

import (
	"github.com/go-redis/redis"
	"context"
	"fmt"
)

type RedisConn struct{
	Ctx  context.Context
}

func (conn *RedisConn) Conn() (*redis.Client, error){
	
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong,err := client.Ping(conn.Ctx).Result()

	if(err != nil){
		fmt.Println("Redis Connect Failed : ",err)
	}else{
		fmt.Println("Redis Connect Success : ",pong)
	}

	return client,err
}
