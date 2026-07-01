package repository

import (
	"database/sql"
	"time"

	"go-rbac-system/internal/model"
)

type UserRepository interface {
	GetByID(id string) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	List() ([]model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	UpdatePassword(userID string, passwordHash string, changedAt time.Time, mustChangePassword bool,) error
	Delete(id string) error
	UpdateStatus(id string, status string) error
	UpdateMustChangePassword(userID string, value bool) error
	UpdateRole(userID string, roleID string,) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB,) UserRepository {
	return &userRepository{
		db: db,
	}
}
// Dùng chung cho user repository
func scanUser(scanner interface {Scan(dest ...any) error},) (*model.User, error) {
	var user model.User
	var mustChange int
	err := scanner.Scan(
		&user.ID,
		&user.TenantID,
		&user.FullName,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.RoleID,
		&user.Status,
		&mustChange,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.PasswordChangedAt,
	)
	if err != nil {
		return nil, err
	}
	user.MustChangePassword = mustChange == 1
	return &user, nil
}

const userSelectQuery = ` SELECT id, tenant_id, full_name, username, email, phone, password_hash, role_id, status, 
	must_change_password, created_at, updated_at, created_by, password_changed_at FROM users`

func (r *userRepository) GetByID(id string,) (*model.User, error) {
	row := r.db.QueryRow(userSelectQuery +  ` WHERE id = :1`,id,)
	return scanUser(row)
}

func (r *userRepository) GetByUsername(username string,) (*model.User, error) {
	row := r.db.QueryRow(userSelectQuery +  ` WHERE username = :1`,username,)
	return scanUser(row)
}

func (r *userRepository) GetByEmail(email string,) (*model.User, error) {
	row := r.db.QueryRow(userSelectQuery +  ` WHERE email = :1`, email,)
	return scanUser(row)
}

func (r *userRepository) List() ([]model.User, error) {
	query := userSelectQuery +  ` ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]model.User, 0)
	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, *user,)
	}
	return users, rows.Err()
}

func (r *userRepository) Create(user *model.User,) error {
	mustChange := 0
	if user.MustChangePassword {
		mustChange = 1
	}
	query := `INSERT INTO users (id, tenant_id, full_name, username, email, phone, password_hash, role_id,
			status, must_change_password, created_at, updated_at, created_by, password_changed_at)
		VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14)`
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
		mustChange,
		user.CreatedAt,
		user.UpdatedAt,
		user.CreatedBy,
		user.PasswordChangedAt,
	)
	return err
}

func (r *userRepository) Update(user *model.User,) error {
	mustChange := 0
	if user.MustChangePassword {
		mustChange = 1
	}
	query := `UPDATE users SET full_name = :1, username = :2, email = :3, phone = :4, role_id = :5,
			status = :6, must_change_password = :7, updated_at = :8 WHERE id = :9`
	_, err := r.db.Exec(
		query,
		user.FullName,
		user.Username,
		user.Email,
		user.Phone,
		user.RoleID,
		user.Status,
		mustChange,
		user.UpdatedAt,
		user.ID,
	)
	return err
}

func (r *userRepository) Delete(id string,) error {
	query := `DELETE FROM users WHERE id = :1`
	_, err := r.db.Exec(query, id,)
	return err
}

func (r *userRepository) UpdatePassword(userID string, passwordHash string, changedAt time.Time, mustChangePassword bool,) error {
	mustChange := 0
	if mustChangePassword {
		mustChange = 1
	}
	query := `UPDATE users SET password_hash = :1, password_changed_at = :2, must_change_password = :3, updated_at = :4 WHERE id = :5`
	_, err := r.db.Exec(
		query,
		passwordHash,
		changedAt,
		mustChange,
		changedAt,
		userID,
	)
	return err
}

func (r *userRepository) UpdateStatus(id string, status string,) error {
	query := `UPDATE users SET status = :1, updated_at = :2 WHERE id = :3`
	result, err := r.db.Exec(query, status, time.Now(), id,)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *userRepository) UpdateMustChangePassword(userID string, value bool) error {
	query := `UPDATE users SET must_change_password = :1, updated_at = :2 WHERE id = :3`
	mustChange := 0
	if value {
		mustChange = 1
	}
	_, err := r.db.Exec(
		query,
		mustChange,
		time.Now(),
		userID,
	)
	return err
}

func (r *userRepository) UpdateRole(userID string, roleID string,) error {
    query := `UPDATE users SET role_id = :1, updated_at = :2 WHERE id = :3`
    _, err := r.db.Exec(query, roleID, time.Now(), userID,
    )
    return err
}