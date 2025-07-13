package Util

import (
	"reflect"
	"testing"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/ConfigSetup"
)

func Test_RedisInitializer(test *testing.T) {
	// Test case 1 :
	testConn := ConfigSetup.RedisConn{}

	result, err := RedisInitializer(testConn)

	if reflect.TypeOf(result).String() != "*redis.Client" {
		test.Error("Test Case 1: Failed - InValid Token", err)
	} else {
		test.Log("Test Case 1: Passed")
	}

	//

	// Test case 2 :

	result, err = RedisInitializer(testConn)
	testConn.Adder = ""

	if reflect.TypeOf(result).String() != "*redis.Client" {
		test.Error("Test Case 2: Failed - InValid Token", err)
	} else {
		test.Log("Test Case 2: Passed")
	}

	//

	// Test case 3 :

	testConn.Adder = "localhost:6379"
	testConn.Password = ""
	result, err = RedisInitializer(testConn)

	if reflect.TypeOf(result).String() != "*redis.Client" {

		test.Error("Test Case 3: Failed - InValid Token", err)
	} else {

		test.Log("Test Case 3: Passed")
	}

	//

	// Test case 4 :

	testConn.Adder = "localhost:6379"
	testConn.Password = ""
	testConn.DB = 0
	result, err = RedisInitializer(testConn)
	//test.Log("Hello")
	if reflect.TypeOf(result).String() != "*redis.Client" {

		test.Error("Test Case 4: Failed - InValid Token", err)
	} else {

		test.Log("Test Case 4: Passed")
	}

	//

}
