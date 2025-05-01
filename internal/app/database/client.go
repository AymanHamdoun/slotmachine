package database

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend/internal/app/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB(ctx context.Context) (*sql.DB, error) {
	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Get().DatabaseUser,
		config.Get().DatabasePass,
		config.Get().DatabaseHost,
		config.Get().DatabasePort,
		config.Get().DatabaseName,
	)

	return sql.Open("mysql", dbString)
}
