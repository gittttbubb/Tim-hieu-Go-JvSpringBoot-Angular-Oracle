package bootstrap

import (
	"database/sql"

	"go-rbac-system/internal/config"
	"go-rbac-system/internal/database"
)

func InitDatabase(cfg *config.Config) (*sql.DB, error) {
	db, err := database.NewOracle(cfg.Database)
	if err != nil {
		return nil, err
	}

	// optional: tuning connection pool (rất quan trọng production)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return db, nil
}