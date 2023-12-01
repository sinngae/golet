package redis

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func ExampleConn() {
	option := &redis.Options{
		Addr:        "127.0.0.1:6379",
		DB:          0,
		PoolSize:    40, /// connection pool
		DialTimeout: time.Hour,
		Password:    "",
	}

	client := redis.NewClient(option)

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("ping redis fail,client=%+v, err=%s", client, err)
		return
	}

	resp := client.Set("abc", 124, time.Minute)
	err = resp.Err()
	fmt.Println(resp, err)

	respGet := client.Get("abc")
	err = respGet.Err()
	fmt.Println(resp, err)

	respGet = client.Get("abcd")
	err = respGet.Err()
	isNil := errors.Is(err, redis.Nil)
	fmt.Println(resp, err, isNil)
	// output:

}
