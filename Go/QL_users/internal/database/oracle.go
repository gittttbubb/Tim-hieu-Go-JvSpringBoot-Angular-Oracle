package database

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"

	"go-user-management/internal/config"
)

func Connect(cfg *config.Config) (*sql.DB, error) {

	dsn := fmt.Sprintf(
		`user="%s" password="%s" connectString="%s:%d/%s"`,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Service,
	)

	db, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}