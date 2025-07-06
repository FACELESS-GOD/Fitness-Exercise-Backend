package Util

import (
	"database/sql"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/ConfigSetup"
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/FieldName"
	_ "github.com/go-sql-driver/mysql"
)

func DBInitializer() (*sql.DB, error) {
	dbinst, err := sql.Open(FieldName.DB_Type, ConfigSetup.DbConnString)

	if err != nil {
		return nil, err
	}
	return dbinst, nil
}
