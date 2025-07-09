package Controller

import (
	"net/http"

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

func (Ctrl *ControllerStruct) AddUser(Writer http.ResponseWriter, Req *http.Request) {}

func (Ctrl *ControllerStruct) ValidateUser(Writer http.ResponseWriter, Req *http.Request) {}

func (Ctrl *ControllerStruct) addUserToDataStore(UserDT StructStore.UserData) error {

}

func (Ctrl *ControllerStruct) validateUserfromDataStore(UserDT StructStore.UserData) (bool, error) {

}
