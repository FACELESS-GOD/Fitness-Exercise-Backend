package ConfigSetup

import (
	"os"
	"strconv"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/FieldName"
)

type RedisConn struct {
	Adder    string
	DB       int
	Password string
}

var DbConnString string
var RedisAdder string
var JWTSecret string
var RedisConenction RedisConn

func InitateSetup(EnvType string) {
	switch EnvType {

	case FieldName.Dev:
		newRedis := RedisConn{}

		dbConnString, isDbConnPresent := os.LookupEnv(FieldName.Dev_DB)
		if isDbConnPresent != true {
			panic("Database Connection String is not set as an environment variable.")
		}

		jwtSecret, isJwtSecret := os.LookupEnv(FieldName.Dev_Jwt_Secret)

		if isJwtSecret != true {
			panic("JWT Secret String is not set as an environment variable.")
		}

		redisAdder, isRedisConnPresent := os.LookupEnv(FieldName.Dev_RedisAdder)

		if isRedisConnPresent != true {
			panic("Redis Connection String is not set as an environment variable.")
		}

		redisPassword, isRedisConnPresent := os.LookupEnv(FieldName.Dev_RedisPassword)

		if isRedisConnPresent != true {
			panic("Redis Connection String is not set as an environment variable.")
		}

		redisDB, isRedisConnPresent := os.LookupEnv(FieldName.Dev_RedisDB)

		if isRedisConnPresent != true {
			panic("Redis Connection String is not set as an environment variable.")
		}

		DbConnString = dbConnString
		JWTSecret = jwtSecret
		newRedis.Adder = redisAdder
		newRedis.Password = redisPassword
		db, err := strconv.Atoi(redisDB)

		if err != nil {
			panic(db)
		}

		newRedis.DB = db
		RedisConenction = newRedis

	case FieldName.QA:
		newRedis := RedisConn{}

		dbConnString, isDbConnPresent := os.LookupEnv(FieldName.QA_DB)
		if isDbConnPresent != true {
			panic("Database Connection String is not set as an environment variable.")
		}

		jwtSecret, isJwtSecret := os.LookupEnv(FieldName.QA_Jwt_Secret)

		if isJwtSecret != true {
			panic("JWT Secret String is not set as an environment variable.")
		}

		redisAdder, isRedisConnPresent := os.LookupEnv(FieldName.QA_RedisAdder)

		if isRedisConnPresent != true {
			panic("Redis Connection String is not set as an environment variable.")
		}

		redisPassword, isRedisConnPresent := os.LookupEnv(FieldName.QA_RedisPassword)

		if isRedisConnPresent != true {
			panic("Redis Connection String is not set as an environment variable.")
		}

		redisDB, isRedisConnPresent := os.LookupEnv(FieldName.QA_RedisDB)

		if isRedisConnPresent != true {
			panic("Redis Connection String is not set as an environment variable.")
		}

		DbConnString = dbConnString

		JWTSecret = jwtSecret

		newRedis.Adder = redisAdder
		newRedis.Password = redisPassword
		db, err := strconv.Atoi(redisDB)

		if err != nil {
			panic(db)
		}

		newRedis.DB = db
		RedisConenction = newRedis

	}

}
