package Model

import (
	"database/sql"
	"log"
	"sync"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/StructStore"
	Util "github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Utility"
)

type DBProcessorInterface interface {
	AddUser(StructStore.UserData) (bool, error)
	ValidateUser(StructStore.UserAuth) (bool, error)
}

type DBProcessor struct {
	DBInstance *sql.DB
}

func NewDBProcessor(Wg *sync.WaitGroup, DbProc *DBProcessor) (*DBProcessor, error) {
	defer Wg.Done()

	//newDBProcessor := DBProcessor{}

	dbinst, err := Util.DBInitializer()

	if err != nil {
		return nil, err
	}

	DbProc.DBInstance = dbinst

	return DbProc, nil

}

func (DbProc DBProcessor) AddUser(UserDt StructStore.UserData) (bool, error) {

	transactinst, err := DbProc.DBInstance.Begin()

	if err != nil {
		log.Println(err)
		return false, err
	}

	stm, err := transactinst.Prepare("INSERT INTO SEC_USER VALUES (?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Println(err)
		transactinst.Rollback()
		return false, err
	}

	stm.Exec(UserDt.AuthorizationId, UserDt.Designation, UserDt.Email, UserDt.FirstName, UserDt.LastName, UserDt.MiddleName, UserDt.Password, UserDt.UserName)

	err_transactinst := transactinst.Commit()

	if err_transactinst != nil {
		log.Println(err)
		transactinst.Rollback()
		return false, err_transactinst
	}

	stm.Close()
	return true, nil
}

func (DbProc DBProcessor) ValidateUser(UserDt StructStore.UserAuth) (bool, error) {

	resp := make([]StructStore.ValidateUserResponse, 1)

	transactinst, err := DbProc.DBInstance.Begin()

	if err != nil {
		log.Println(err)
		return false, err
	}

	stm, err := transactinst.Prepare("SELECT UserName from SEC_USER WHERE UserName = ? AND Password = ?")
	if err != nil {
		log.Println(err)
		transactinst.Rollback()
		return false, err
	}

	//rst, is := stm.Exec(UserDt.UserName, UserDt.Password)

	rst, err := stm.Query(UserDt.UserName, UserDt.Password)

	if err != nil {
		return false, err
	}

	err = transactinst.Commit()

	if err != nil {
		transactinst.Rollback()
		return false, err
	}

	for rst.Next() {
		currres := StructStore.ValidateUserResponse{}

		err := rst.Scan(&currres.UserName)
		if err != nil {
			return false, err
		}

		resp = append(resp, currres)
	}

	if len(resp) != 1 {
		return false, nil
	}
	return true, nil
}

