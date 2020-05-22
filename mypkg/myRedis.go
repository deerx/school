package mypkg

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

var (
	//MyClient 公共redis
	MyClient *redis.Client
)

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	MyClient = client
}
