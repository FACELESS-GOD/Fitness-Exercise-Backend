package Util

import (
	"testing"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/StructStore"
)

func Test_CreateToken(test *testing.T) {

	testTokenProcessor := TokenProcessor{}

	// Test case 1 :
	testUserDT := StructStore.UserData{}
	result, err := testTokenProcessor.CreateToken(testUserDT, "")

	if err.Error() == "" {
		test.Error("Test Case 1: Failed", err)
	} else if result != "" {
		test.Error("Test Case 1: Failed - InValid Token", err)
	} else {
		test.Log("Test Case 1: Passed")
	}

	// Test case 2 :
	testUserDT.UserName = ""
	testUserDT.Designation = 0

	result, err = testTokenProcessor.CreateToken(testUserDT, "test_Secret")

	if err.Error() == "" {
		test.Error("Test Case 2: Failed", err)
	} else if result != "" {
		test.Error("Test Case 2: Failed - InValid Token", err)
	} else {
		test.Log("Test Case 2: Passed")
	}

	//

	// Test case 3 :
	testUserDT.UserName = "Test UserName"
	testUserDT.Designation = 0

	result, err = testTokenProcessor.CreateToken(testUserDT, "test_Secret")

	if err.Error() == "" {
		test.Error("Test Case 3: Failed", err)
	} else if result != "" {
		test.Error("Test Case 3: Failed - InValid Token", err)
	} else {
		test.Log("Test Case 4: Passed")
	}
	//

	// Test case 4 :
	testUserDT.UserName = "Test UserName"
	testUserDT.Designation = 1

	result, err = testTokenProcessor.CreateToken(testUserDT, "")

	if err.Error() == "" {
		test.Error("Test Case 4: Failed", err)
	} else if result != "" {
		test.Error("Test Case 4: Failed - InValid Token", err)
	} else {
		test.Log("Test Case 4: Passed")
	}

	//

	// Test case 4 :
	testUserDT.UserName = "Test UserName"
	testUserDT.Designation = 1

	result, err = testTokenProcessor.CreateToken(testUserDT, "Test_SecretKey")

	if err.Error() == "" {
		test.Error("Test Case 5: Failed", err)
	} else if result != "" {
		test.Error("Test Case 5: Failed - InValid Token", err)
	} else {
		test.Log("Test Case 5: Passed")
	}

	//

}
