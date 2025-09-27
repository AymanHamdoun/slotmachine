package database

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend/internal/app/config"

	_ "github.com/go-sql-driver/mysql"
)

var dbConnection *sql.DB

func GetDB(ctx context.Context) (*sql.DB, error) {
	if dbConnection != nil {
		return dbConnection, nil
	}

	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Get().DatabaseUser,
		config.Get().DatabasePass,
		config.Get().DatabaseHost,
		config.Get().DatabasePort,
		config.Get().DatabaseName,
	)

	var err error
	dbConnection, err = sql.Open("mysql", dbString)
	if err != nil {
		return nil, err
	}

	return dbConnection, nil
}
