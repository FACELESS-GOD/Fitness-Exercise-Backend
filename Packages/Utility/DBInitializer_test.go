package Util

import (
	"reflect"
	"testing"
)

func Test_DBInitializer(test *testing.T) {
	// Test case 1 :

	result, err := DBInitializer("")

	if reflect.TypeOf(result).String() != "*sql.DB" {
		test.Error("Test Case 1: Failed - InValid Token", err)
	} else {
		test.Log("Test Case 1: Passed")
	}

	//

	// Test case 2 :

	result, err = DBInitializer("df")

	if reflect.TypeOf(result).String() != "*sql.DB" {
		test.Error("Test Case 2: Failed - InValid Token", err)
	} else {
		test.Log("Test Case 2: Passed")
	}

	//

	// Test case 3 :

	result, err = DBInitializer("root:Admin@123@tcp/AuthData_Test?charset=UTF8&parseTime=True&loc=Local")

	if reflect.TypeOf(result).String() != "*sql.DB" {

		test.Error("Test Case 3: Failed - InValid Token", err)
	} else {

		test.Log("Test Case 3: Passed")
	}

	//

}
