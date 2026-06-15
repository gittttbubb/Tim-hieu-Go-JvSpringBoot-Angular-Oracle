package repository

import (
	"database/sql"

	"go-user-management/internal/model"
)

type UserRepository interface {
	GetByID(id string) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id string) error
	List() ([]model.User, error)
	UpdatePassword(userID string, passwordHash string,) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByID(id string) (*model.User, error) {
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
	FROM users WHERE id = :1`
	var user model.User
	err := r.db.QueryRow(query, id,).Scan(
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
func (r *userRepository) GetByUsername(username string) (*model.User, error) {
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
	FROM users WHERE username = :1`
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

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
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
	FROM users WHERE email = :1`
	var user model.User
	err := r.db.QueryRow(query, email,).Scan(
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

func (r *userRepository) Create(user *model.User) error {
	query := `
	INSERT INTO users (
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
		created_by
	)
	VALUES (
		:1,:2,:3,:4,:5,
		:6,:7,:8,:9,:10,:11
	)`
	_, err := r.db.Exec(
		query,
		user.ID,
		user.TenantID,
		user.FullName,
		user.Username,
		user.Email,
		user.Phone,
		user.PasswordHash,
		user.RoleID,
		user.Status,
		user.MustChangePassword,
		user.CreatedBy,
	)
	return err
}

func (r *userRepository) Update(user *model.User,) error {
	query := `
	UPDATE users
	SET
		full_name = :1,
		email = :2,
		phone = :3,
		role_id = :4,
		status = :5,
		updated_at = SYSTIMESTAMP
	WHERE id = :6`
	_, err := r.db.Exec(
		query,
		user.FullName,
		user.Email,
		user.Phone,
		user.RoleID,
		user.Status,
		user.ID,
	)
	return err
}

func (r *userRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = :1`, id)
	return err
}

func (r *userRepository) List() ([]model.User, error,) {
	query := `
	SELECT
		id,
		tenant_id,
		full_name,
		username,
		email,
		phone,
		role_id,
		status,
		must_change_password,
		created_at,
		updated_at,
		password_changed_at
	FROM users ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	// đảm bảo việc đóng rows sẽ luôn luôn được thực thi ngay trước khi hàm List() kết thúc
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		err := rows.Scan(
			&u.ID,
			&u.TenantID,
			&u.FullName,
			&u.Username,
			&u.Email,
			&u.Phone,
			&u.RoleID,
			&u.Status,
			&u.MustChangePassword,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.PasswordChangedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) UpdatePassword(userID string, passwordHash string,) error {
	query := `
	UPDATE users
	SET
		password_hash = :1,
		password_changed_at = SYSTIMESTAMP,
		must_change_password = 0
	WHERE id = :2
	`
	_, err := r.db.Exec(query, passwordHash, userID,)
	return err
}