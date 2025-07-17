package Model

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/ConfigSetup"
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/StructStore"
	Util "github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Utility"
)

func Test_NewRedisInstance(test *testing.T) {

	rdb := RedisProcessor{}
	ctx, call := context.WithTimeout(context.Background(), time.Second)
	rdb.Ctx = ctx
	rdb.CallbackFunc = call
	conn := ConfigSetup.RedisConn{}
	conn.Adder = "localhost:6379"
	conn.DB = 0
	conn.Password = ""
	redisInst, err := Util.RedisInitializer(conn)

	rdb.RedisInst = redisInst

	// Test Case 1:

	UserDT := StructStore.UserData{}

	IsValid, err := rdb.AddUser(UserDT)

	if err == nil {
		test.Error("Test Case 1: Failed ", err)
	} else if IsValid == true {
		test.Error("Test Case 1: Failed ", err)
	} else {
		test.Log("Test Case 1: Passed ")
	}

	//

	// Test Case 2
	UserDT.UserName = "Admin"
	UserDT.Password = "Admin"
	UserDT.AuthorizationId = 1
	UserDT.Designation = 1
	UserDT.IsValid = true
	UserDT.Last_Modified_Date = time.Now().String()

	IsValid, err = rdb.AddUser(UserDT)
	if err != nil {
		test.Error("Test Case 2: Failed ", err)
	} else if IsValid != true {
		test.Error("Test Case 2: Failed ", err)
	} else {
		test.Log("Test Case 2: Passed ")
	}

	//

}

func Test_AddUser(test *testing.T) {

	rdb := RedisProcessor{}
	ctx, call := context.WithTimeout(context.Background(), time.Second)
	rdb.Ctx = ctx
	rdb.CallbackFunc = call
	conn := ConfigSetup.RedisConn{}
	conn.Adder = "localhost:6379"
	conn.DB = 0
	conn.Password = ""
	redisInst, err := Util.RedisInitializer(conn)

	rdb.RedisInst = redisInst

	// Test Case 1:

	UserDT := StructStore.UserAuth{}

	IsValid, err := rdb.ValidateUser(UserDT)

	if err == nil {
		test.Error("Test Case 1: Failed ", err)
	} else if IsValid == true {
		test.Error("Test Case 1: Failed ", err)
	} else {
		test.Log("Test Case 1: Passed ")
	}

	//

	// Test Case 2
	UserDT.UserName = "Admin"
	UserDT.Password = "Admin"

	IsValid, err = rdb.ValidateUser(UserDT)
	if err != nil {
		test.Error("Test Case 2: Failed ", err)
	} else if IsValid != true {
		test.Error("Test Case 2: Failed ", err)
	} else {
		test.Log("Test Case 2: Passed ")
	}

	//
	// Test Case 3
	UserDT.UserName = "Admin"
	UserDT.Password = "Admin2safg"

	IsValid, err = rdb.ValidateUser(UserDT)
	if err != nil {
		test.Error("Test Case 3: Failed ", err)
	} else if IsValid == true {
		test.Error("Test Case 3: Failed ", err)
	} else {
		test.Log("Test Case 3: Passed ")
	}

	//

}

func Test_GetUserData(test *testing.T) {

	rdb := RedisProcessor{}
	ctx, call := context.WithTimeout(context.Background(), time.Second)
	rdb.Ctx = ctx
	rdb.CallbackFunc = call
	conn := ConfigSetup.RedisConn{}
	conn.Adder = "localhost:6379"
	conn.DB = 0
	conn.Password = ""
	redisInst, err := Util.RedisInitializer(conn)
	if err != nil {
		fmt.Print(err)
	}

	rdb.RedisInst = redisInst

	// Test Case 1:

	userDT, isValid := rdb.GetUserData("")

	if isValid == true {
		test.Error("Test Case 1: Failed ")
	} else if userDT != "" {
		test.Error("Test Case 1: Failed ")
	} else {
		test.Log("Test Case 1: Passed ")
	}

	//

	// Test Case 2:

	userDT, isValid = rdb.GetUserData("Admin")

	if isValid == false {
		test.Error("Test Case 2: Failed ")
	} else if userDT == "" {
		test.Error("Test Case 2: Failed ")
	} else {
		test.Log("Test Case 2: Passed ")
	}

	//

	// Test Case 3:

	userDT, isValid = rdb.GetUserData("Admindfg")

	if isValid == true {
		test.Error("Test Case 3: Failed ")
	} else if userDT != "" {
		test.Error("Test Case 3: Failed ")
	} else {
		test.Log("Test Case 3: Passed ")
	}

	//

}
