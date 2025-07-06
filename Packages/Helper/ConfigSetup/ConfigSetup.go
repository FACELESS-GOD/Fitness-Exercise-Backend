package ConfigSetup

import (
	"os"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/FieldName"
)

var DbConnString string
var RedisConnString string

func InitateSetup(EnvType string) {
	switch EnvType {

	case FieldName.Dev:
		dbConnString, isDbConnPresent := os.LookupEnv(FieldName.Dev_DB)
		if isDbConnPresent != true {
			panic("Database Connection String is not set as an environment variable.")
		}

		redisConnString, isRedisConnPresent := os.LookupEnv(FieldName.Dev_Redis)

		if isRedisConnPresent != true {
			panic("Redis Connection String is not set as an environment variable.")
		}

		DbConnString = dbConnString
		RedisConnString = redisConnString

	case FieldName.QA:

		dbConnString, isDbConnPresent := os.LookupEnv(FieldName.QA_DB)
		if isDbConnPresent != true {
			panic("Database Connection String is not set as an environment variable.")
		}

		redisConnString, isRedisConnPresent := os.LookupEnv(FieldName.QA_Redis)

		if isRedisConnPresent != true {
			panic("Redis Connection String is not set as an environment variable.")
		}

		DbConnString = dbConnString
		RedisConnString = redisConnString

	}

}
