package Util

import (
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/ConfigSetup"
	"github.com/go-redis/redis/v8"
)

func RedisInitializer() (*redis.Client, error) {

	//redis://<user>:<pass>@localhost:6379/<db>"
	newRedisOption, err := redis.ParseURL(ConfigSetup.RedisConnString)
	if err != nil {
		return nil, err
	}

	newClient := redis.NewClient(newRedisOption)

	return newClient, nil

}
