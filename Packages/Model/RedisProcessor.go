package Model

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/StructStore"
	Util "github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Utility"
	"github.com/go-redis/redis/v8"
)

type RedisProcessorInterface interface {
	AddUser(*sync.WaitGroup, StructStore.UserData) (bool, error)
	ValidateUser(StructStore.UserData) (bool, error)
}

type RedisProcessor struct {
	RedisInst    *redis.Client
	Ctx          context.Context
	CallbackFunc context.CancelFunc
}

func NewRedisInstance(WG *sync.WaitGroup, RedisProc *RedisProcessor) (*RedisProcessor, error) {
	defer WG.Done()
	//Inst := RedisProcessor{}

	redisInst, err := Util.RedisInitializer()

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	RedisProc.RedisInst = redisInst
	RedisProc.Ctx = ctx
	RedisProc.CallbackFunc = cancel

	return RedisProc, nil
}

func (Red RedisProcessor) AddUser(WG *sync.WaitGroup, UserDT StructStore.UserData) (bool, error) {
	defer WG.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userJSON, err := json.Marshal(UserDT)

	if err != nil {
		return false, err
	}

	err = Red.RedisInst.Set(ctx, UserDT.UserName, userJSON, 0).Err()
	if err != nil {
		return false, err
	}

	return true, nil

}

func (Red RedisProcessor) ValidateUser(UserDT StructStore.UserData) (bool, error) {
	var currUser StructStore.UserData
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	val, err := Red.RedisInst.Get(ctx, UserDT.UserName).Result()

	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(val), currUser)

	if err != nil {
		return false, err
	}

	if currUser.Password != UserDT.Password {
		return false, nil
	} else {
		return true, nil
	}
}

func (Red RedisProcessor) GetUserData(UserName string) (string, bool) {
	val, err := Red.RedisInst.Get(Red.Ctx, UserName).Result()

	switch {
	case err == redis.Nil:
		return "", false
	case err != nil:
		log.Println(err)
		return "", false
	}
	return val, true

}
