package Controller

import (
	"log"
	"net/http"
	"sync"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/ConfigSetup"
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/StructStore"
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Model"
	Util "github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Utility"
)

type ControllerStruct struct {
	DbInst    Model.DBProcessorInterface
	RedisInst Model.RedisProcessorInterface
	TokenProc Util.TokenProcessor
}

func NewController(DbInst Model.DBProcessor, RedisInst Model.RedisProcessor) *ControllerStruct {
	ctrlInst := ControllerStruct{}
	ctrlInst.DbInst = DbInst
	ctrlInst.RedisInst = RedisInst
	ctrlInst.TokenProc = Util.TokenProcessor{}
	return &ctrlInst
}

func (Ctrl *ControllerStruct) AddUser(Writer http.ResponseWriter, Req *http.Request) {
	newUser := &StructStore.UserData{}

	Util.ParseBody(Req, newUser)

	err := Ctrl.addUserToDB(*newUser)
	if err != nil {
		Ctrl.createErrorPayload(Writer)
		return
	}

	err = Ctrl.addUserToDataStore(*newUser)

	if err != nil {
		Ctrl.createErrorPayload(Writer)
		return
	}

	tkn, err := Ctrl.TokenProc.CreateToken(*newUser, ConfigSetup.JWTSecret)
	if err != nil {
		log.Println(err)
	}

	wg := sync.WaitGroup{}

	go Ctrl.DbInst.AddToken(&wg, newUser.UserName, tkn)
	wg.Add(1)

	go Ctrl.RedisInst.AddToken(&wg, newUser.UserName, tkn)
	wg.Add(1)

	Ctrl.createCorrectPayload(Writer)
	wg.Wait()
	return
}

func (Ctrl *ControllerStruct) ValidateUser(Writer http.ResponseWriter, Req *http.Request) {
	authDT := &StructStore.UserAuth{}
	Util.ParseBody(Req, authDT)
	isValid, err := Ctrl.validateUserfromDataStore(*authDT)

	if err != nil {
		log.Println(err)
	}

	if isValid == false {

		isValid, err = Ctrl.validateUserfromDB(*authDT)
		if err != nil {
			log.Println(err)
		}

		if isValid != true {
			Ctrl.createErrorPayload(Writer)
			return
		} else {
			Ctrl.createCorrectPayload(Writer)
			return
		}

	} else {
		Ctrl.createCorrectPayload(Writer)
		return
	}

}

func (Ctrl *ControllerStruct) addUserToDB(UserDT StructStore.UserData) error {
	isDone, err := Ctrl.DbInst.AddUser(UserDT)
	if err != nil {
		return err
	}
	if isDone != true {
		log.Default()
	}

	isDone, err = Ctrl.RedisInst.AddUser(UserDT)
	if err != nil {
		return err
	}

	return nil

}

func (Ctrl *ControllerStruct) addUserToDataStore(UserDT StructStore.UserData) error {

	isDone, err := Ctrl.RedisInst.AddUser(UserDT)
	if err != nil {
		return err
	}
	if isDone != true {
		log.Println()
	}
	return nil

}

func (Ctrl *ControllerStruct) validateUserfromDataStore(UserDT StructStore.UserAuth) (bool, error) {

	isValid, err := Ctrl.DbInst.ValidateUser(UserDT)
	if err != nil {
		return isValid, err
	}
	return isValid, nil

}

func (Ctrl *ControllerStruct) validateUserfromDB(UserDT StructStore.UserAuth) (bool, error) {
	isValid, err := Ctrl.RedisInst.ValidateUser(UserDT)
	if err != nil {
		return isValid, err
	}
	return isValid, nil

}

func (Ctrl *ControllerStruct) createErrorPayload(Writer http.ResponseWriter) {
	Writer.WriteHeader(http.StatusInternalServerError)
}

func (Ctrl *ControllerStruct) createCorrectPayload(Writer http.ResponseWriter) {
	Writer.WriteHeader(http.StatusOK)
}
