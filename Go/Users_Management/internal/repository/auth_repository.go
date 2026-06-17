package repository

import (
	"database/sql"

	"go-rbac-system/internal/model"
)

type AuthRepository interface {
	GetUserByUsername(username string) (*model.User, error)
	GetRolePermissions(roleID string) ([]model.RolePermission, error)
	GetUserPermissionOverrides(userID string) ([]model.UserPermissionOverride, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) GetUserByUsername(username string,) (*model.User, error) {
	query := `SELECT id, tenant_id, full_name, username, email, phone, password_hash,
			role_id, status, must_change_password, created_at, updated_at, created_by, password_changed_at
		FROM users WHERE username = :1`
	var user model.User
	var mustChange int
	err := r.db.QueryRow(query, username,
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

func (r *authRepository) GetRolePermissions(roleID string,) ([]model.RolePermission, error) {
	query := `SELECT role_id, permission_id, granted, data_scope FROM role_permissions WHERE role_id = :1`
	rows, err := r.db.Query(query, roleID,)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	permissions := make([]model.RolePermission, 0,)
	for rows.Next() {
		var item model.RolePermission
		var granted int
		err := rows.Scan(&item.RoleID, &item.PermissionID, &granted, &item.DataScope,)
		if err != nil {
			return nil, err
		}
		item.Granted = granted == 1
		permissions = append(permissions, item,)
	}
	return permissions, rows.Err()
}

func (r *authRepository) GetUserPermissionOverrides(userID string,) ([]model.UserPermissionOverride, error) {
	query := `SELECT id, user_id, permission_id, granted, reason, created_by, created_at, data_scope
		FROM user_permission_overrides WHERE user_id = :1`
	rows, err := r.db.Query(query, userID,)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	overrides := make([]model.UserPermissionOverride, 0,)
	for rows.Next() {
		var item model.UserPermissionOverride
		var granted int
		err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.PermissionID,
			&granted,
			&item.Reason,
			&item.CreatedBy,
			&item.CreatedAt,
			&item.DataScope,
		)
		if err != nil {
			return nil, err
		}
		item.Granted = granted == 1
		overrides = append(
			overrides,
			item,
		)
	}
	return overrides, rows.Err()
}
