package Util

import (
	"database/sql"
	"errors"

	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Helper/FieldName"
	_ "github.com/go-sql-driver/mysql"
)

func DBInitializer(Conn string) (*sql.DB, error) {

	if Conn == "" {
		return nil, errors.New("Invalid Config")
	}

	dbinst, err := sql.Open(FieldName.DB_Type, Conn)

	if err != nil {
		return nil, err
	}

	return dbinst, nil
}
