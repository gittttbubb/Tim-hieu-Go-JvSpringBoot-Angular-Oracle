package database

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"

	"go-rbac-system/internal/config"
)

func NewOracle(cfg config.DatabaseConfig) (*sql.DB, error) {

	dsn := fmt.Sprintf(
		`user="%s" password="%s" connectString="%s:%d/%s"`,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.ServiceName,
	)

	db, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}