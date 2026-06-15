package repository

import (
	"database/sql"

	"go-user-management/internal/model"
)

type AuthRepository interface {
	FindByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}
func (r *authRepository) FindByUsername(username string) (*model.User, error) {

	query := `SELECT
		id,
		tenant_id,
		full_name,
		username,
		email,
		phone,
		password_hash,
		role_id,
		status,
		must_change_password,
		created_at,
		updated_at,
		created_by,
		password_changed_at
	FROM users WHERE username = :1
	`
	var user model.User
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.TenantID,
		&user.FullName,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.RoleID,
		&user.Status,
		&user.MustChangePassword,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.PasswordChangedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) GetByEmail(email string,) (*model.User, error) {
	query := `
	SELECT
		id,
		tenant_id,
		full_name,
		username,
		email,
		phone,
		password_hash,
		role_id,
		status,
		must_change_password,
		created_at,
		updated_at,
		created_by,
		password_changed_at
	FROM users
	WHERE email = :1
	`

	var user model.User
	err := r.db.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.TenantID,
		&user.FullName,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.RoleID,
		&user.Status,
		&user.MustChangePassword,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.PasswordChangedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}