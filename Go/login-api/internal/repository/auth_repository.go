package repository

import (
	"database/sql"

	"login-api/internal/model"
)
// public 
type AuthRepository interface {
	FindByUsername(username string) (*model.User, error)
	GetByID(id int) (*model.User, error)
	Register(user *model.User) error
	// USER
	UpdateProfile(id int, username string) error
	ChangePassword(id int, password string) error
	// ADMIN
	GetAllUsers() ([]model.User, error)
	UpdateRole(id int, role string) error
}
// private
type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}
// Cặp (r *authRepository) khai báo rằng hàm FindByUsername thuộc về struct authRepository, giúp hàm này có thể truy cập vào các thuộc tính bên trong struct (như biến db *sql.DB)
func (r *authRepository) FindByUsername(username string,) (*model.User, error) {
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

func (r *authRepository) Register(user *model.User) error {
	query := `INSERT INTO USERS (USERNAME, PASSWORD, ROLE) VALUES (:1, :2, :3)`
	_, err := r.db.Exec(query, user.Username, user.Password, user.Role,)
	return err
}
// USER
func (r *authRepository) UpdateProfile(id int, username string,) error {
    query := `UPDATE USERS SET USERNAME = :1 WHERE ID = :2`
    _, err := r.db.Exec(query, username, id)
    return err
}

func (r *authRepository) ChangePassword(id int, password string,) error {
    query := `UPDATE USERS SET PASSWORD = :1 WHERE ID = :2`
    _, err := r.db.Exec(query, password, id)
    return err
}

// ADMIN
func (r *authRepository) GetAllUsers() ([]model.User, error) {
	query := `SELECT ID, USERNAME, ROLE FROM USERS`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Role,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *authRepository) UpdateRole(id int, role string,) error {
	query := `UPDATE USERS SET ROLE = :1 WHERE ID = :2`
	_, err := r.db.Exec(query, role, id)
	return err
}