package Model

import (
	"database/sql"
	"sync"

	Util "github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Utility"
)

type DBProcessorInterface interface {
	AddUser()
	ValidateUser()
}

type DBProcessor struct {
	DBInstance *sql.DB
}

func NewDBProcessor(Wg *sync.WaitGroup) (*DBProcessor, error) {
	defer Wg.Done()

	newDBProcessor := DBProcessor{}

	dbinst, err := Util.DBInitializer()

	if err != nil {
		return nil, err
	}

	newDBProcessor.DBInstance = dbinst

	return &newDBProcessor, nil

}
