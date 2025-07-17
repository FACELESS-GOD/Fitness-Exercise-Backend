package Model

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/ConfigSetup"
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/StructStore"
	Util "github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Utility"
	"github.com/go-redis/redis/v8"
)

type RedisProcessorInterface interface {
	AddUser(StructStore.UserData) (bool, error)
	ValidateUser(StructStore.UserAuth) (bool, error)
	AddToken(*sync.WaitGroup, string, string)
}

type RedisProcessor struct {
	RedisInst    *redis.Client
	Ctx          context.Context
	CallbackFunc context.CancelFunc
}

func NewRedisInstance(Wg *sync.WaitGroup, RedisProc *RedisProcessor) (*RedisProcessor, error) {
	defer Wg.Done()
	//Inst := RedisProcessor{}

	redisInst, err := Util.RedisInitializer(ConfigSetup.RedisConenction)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	RedisProc.RedisInst = redisInst
	RedisProc.Ctx = ctx
	RedisProc.CallbackFunc = cancel

	return RedisProc, nil
}

func (Red RedisProcessor) AddUser(UserDT StructStore.UserData) (bool, error) {
	if UserDT.AuthorizationId == 0 || UserDT.Designation == 0 || UserDT.UserName == "" || UserDT.Password == "" {
		return false, errors.New("Invalid Data")
	}

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

func (Red RedisProcessor) ValidateUser(UserDT StructStore.UserAuth) (bool, error) {
	if UserDT.UserName == "" || UserDT.Password == "" {
		return false, errors.New("Invalid Data.")
	}
	var currUser StructStore.UserData
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	val, err := Red.RedisInst.Get(ctx, UserDT.UserName).Result()

	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(val), &currUser)

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
	if UserName == "" {
		return "", false
	}
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

func (Red RedisProcessor) AddToken(Wg *sync.WaitGroup, UserName string, Token string) {
	defer Wg.Done()
	err := Red.RedisInst.Set(Red.Ctx, Token, UserName, time.Hour).Err()
	if err != nil {
		log.Println(err)
		return
	}
	return

}
