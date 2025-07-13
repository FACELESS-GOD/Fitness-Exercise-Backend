package Util

import (
	"context"
	"fmt"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/ConfigSetup"
	"github.com/go-redis/redis/v8"
)

func RedisInitializer(Conn ConfigSetup.RedisConn) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     Conn.Adder,
		Password: Conn.Password,
		DB:       Conn.DB,
	})

	res, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(res)
		return nil, err
	}

	return client, nil
}
