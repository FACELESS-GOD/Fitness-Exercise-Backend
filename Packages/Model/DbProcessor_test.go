package Model

import (
	"testing"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/StructStore"
	Util "github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Utility"
)

func Test_AddDBUser(test *testing.T) {

	//wg := sync.WaitGroup{}

	test_DbProc := DBProcessor{}
	dbInst, err := Util.DBInitializer("root:Admin@123@tcp/AuthData_Test?charset=UTF8&parseTime=True&loc=Local")

	test_DbProc.DBInstance = dbInst

	//go NewDBProcessor(&wg, &test_DbProc)

	//wg.Add(1)
	//wg.Wait()
	test_UserData := StructStore.UserData{}
	// Test case 1 :

	result, err := test_DbProc.AddUser(test_UserData)

	if err == nil {
		test.Error("Test Case 1: Failed")
	} else if result == true {
		test.Error("Test Case 1: Failed")
	} else {
		test.Log("Test Case 1: Passed")
	}

	//

	// Test case 2 :
	test_UserData.UserName = "Admin"

	result, err = test_DbProc.AddUser(test_UserData)

	if err == nil {
		test.Error("Test Case 2: Failed")
	} else if result == true {
		test.Error("Test Case 2: Failed")
	} else {
		test.Log("Test Case 2: Passed")
	}

	//

	// Test case 3 :
	test_UserData.UserName = "Admin"
	test_UserData.Password = "Admin"

	result, err = test_DbProc.AddUser(test_UserData)

	if err == nil {
		test.Error("Test Case 3: Failed")
	} else if result == true {
		test.Error("Test Case 3: Failed")
	} else {
		test.Log("Test Case 3: Passed")
	}

	//

	// Test case 4 :
	test_UserData.UserName = "Admin"
	test_UserData.Password = "Admin"
	test_UserData.AuthorizationId = 1

	result, err = test_DbProc.AddUser(test_UserData)

	if err == nil {
		test.Error("Test Case 4: Failed")
	} else if result == true {
		test.Error("Test Case 4: Failed")
	} else {
		test.Log("Test Case 4: Passed")
	}

	//

	// Test case 5 :
	test_UserData.UserName = "Admin"
	test_UserData.Password = "Admin"
	test_UserData.AuthorizationId = 1
	test_UserData.Designation = 1

	result, err = test_DbProc.AddUser(test_UserData)

	if err != nil {
		test.Error("Test Case 5: Failed", err)
	} else if result != true {
		test.Error("Test Case 5: Failed", err)
	} else {
		test.Log("Test Case 5: Passed")
	}

	//

	// Test case 6 :
	test_UserData.UserName = "Admin"
	test_UserData.Password = "Admin"
	test_UserData.AuthorizationId = 1
	test_UserData.Designation = 1
	test_UserData.IsValid = true

	result, err = test_DbProc.AddUser(test_UserData)

	if err != nil {
		test.Error("Test Case 6: Failed", err)
	} else if result != true {
		test.Error("Test Case 6: Failed", err)
	} else {
		test.Log("Test Case 6: Passed")
	}

	//

	// Test case 7 :
	test_UserData.UserName = "Admin3"
	test_UserData.Password = "Admin3"
	test_UserData.AuthorizationId = 1
	test_UserData.Designation = 1
	test_UserData.IsValid = true
	test_UserData.Last_Modified_Date = "13/07/2025"

	result, err = test_DbProc.AddUser(test_UserData)

	if err != nil {
		test.Error("Test Case 7: Failed", err)
	} else if result != true {
		test.Error("Test Case 7: Failed", err)
	} else {
		test.Log("Test Case 7: Passed")
	}

	//

}

func Test_ValidateDBUser(test *testing.T) {

	test_DbProc := DBProcessor{}
	dbInst, err := Util.DBInitializer("root:Admin@123@tcp/AuthData_Test?charset=UTF8&parseTime=True&loc=Local")

	test_DbProc.DBInstance = dbInst

	//go NewDBProcessor(&wg, &test_DbProc)

	//wg.Add(1)
	//wg.Wait()
	test_UserData := StructStore.UserAuth{}
	// Test case 1 :

	result, err := test_DbProc.ValidateUser(test_UserData)

	if err == nil {
		test.Error("Test Case 1: Failed")
	} else if result == true {
		test.Error("Test Case 1: Failed")
	} else {
		test.Log("Test Case 1: Passed")
	}

	//

	test_UserData.UserName = "Admin"
	test_UserData.Password = "Admin"

	result, err = test_DbProc.ValidateUser(test_UserData)

	if err != nil {
		test.Error("Test Case 2: Failed")
	} else if result != true {
		test.Error("Test Case 2: Failed")
	} else {
		test.Log("Test Case 2: Passed")
	}

}
