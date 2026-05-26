package database

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"

	"login-api/internal/config"
)

func Connect(cfg config.Config) (*sql.DB, error) {

	dsn :=
		fmt.Sprintf(
			`user="%s" password="%s" connectString="%s:%s/%s"`,

			cfg.DBUser,
			cfg.DBPassword,

			cfg.DBHost,
			cfg.DBPort,

			cfg.DBService,
		)

	db, err :=
		sql.Open(
			"godror",
			dsn,
		)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}