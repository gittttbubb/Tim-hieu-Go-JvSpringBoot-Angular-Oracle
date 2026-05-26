package repository

import (
	"database/sql"

	"login-api/internal/model"
)

type AuthRepository interface {
	FindByUsername(username string) (*model.User, error)
	GetByID(id int) (*model.User, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(
	db *sql.DB,
) AuthRepository {

	return &authRepository{
		db: db,
	}
}

func (r *authRepository) FindByUsername(
	username string,
) (*model.User, error) {

	query := `SELECT ID, USERNAME, PASSWORD, ROLE FROM USERS WHERE USERNAME=:1`

	var user model.User
	err := r.db.QueryRow(query, username,).Scan(&user.ID, &user.Username, &user.Password, &user.Role,)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) GetByID(id int) (*model.User, error) {

	query := `SELECT ID, USERNAME, PASSWORD, ROLE FROM USERS WHERE ID = :1`

	var user model.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Role,)
	if err != nil {
		return nil, err
	}

	return &user, nil
}