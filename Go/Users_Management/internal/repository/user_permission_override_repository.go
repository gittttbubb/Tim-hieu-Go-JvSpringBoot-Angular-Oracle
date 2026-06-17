package repository

import (
	"database/sql"

	"go-rbac-system/internal/model"
)

type UserPermissionOverrideRepository interface {
	GetByID(id string) (*model.UserPermissionOverride, error)
	GetByUserAndPermission(userID string,permissionID string) (*model.UserPermissionOverride, error)
	GetByUserID(userID string,) ([]model.UserPermissionOverride, error)
	Create(override *model.UserPermissionOverride,) error
	Update(override *model.UserPermissionOverride,) error
	Delete(id string) error
}

type userPermissionOverrideRepository struct {
	db *sql.DB
}

func NewUserPermissionOverrideRepository(db *sql.DB,) UserPermissionOverrideRepository {
	return &userPermissionOverrideRepository{
		db: db,
	}
}

func scanUserPermissionOverride(scanner interface {Scan(dest ...any) error},) (*model.UserPermissionOverride, error) {
	var item model.UserPermissionOverride
	var granted int
	err := scanner.Scan(
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
	return &item, nil
}

func (r *userPermissionOverrideRepository) GetByID(id string,) (*model.UserPermissionOverride, error) {
	query := `SELECT id, user_id, permission_id, granted, reason, created_by, created_at, data_scope
		FROM user_permission_overrides WHERE id = :1`
	row := r.db.QueryRow(query, id,)
	return scanUserPermissionOverride(row)
}

func (r *userPermissionOverrideRepository) GetByUserAndPermission(userID string, permissionID string,) (*model.UserPermissionOverride, error) {
	query := `SELECT id, user_id, permission_id, granted, reason, created_by, created_at, data_scope
		FROM user_permission_overrides WHERE user_id = :1 AND permission_id = :2`
	row := r.db.QueryRow(query, userID, permissionID,)
	return scanUserPermissionOverride(row)
}

func (r *userPermissionOverrideRepository) GetByUserID(userID string,) ([]model.UserPermissionOverride, error) {
	query := `SELECT id, user_id, permission_id, granted, reason, created_by, created_at, data_scope
		FROM user_permission_overrides WHERE user_id = :1 ORDER BY created_at DESC`
	rows, err := r.db.Query(query, userID,)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]model.UserPermissionOverride, 0,)
	for rows.Next() {
		item, err := scanUserPermissionOverride(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *item,)
	}
	return result, rows.Err()
}

func (r *userPermissionOverrideRepository) Create(override *model.UserPermissionOverride,) error {
	granted := 0
	if override.Granted {
		granted = 1
	}
	query := `INSERT INTO user_permission_overrides ( id, user_id, permission_id, granted, reason, created_by, created_at, data_scope)
		VALUES (:1, :2, :3, :4, :5, :6, :7, :8)`
	_, err := r.db.Exec(
		query,
		override.ID,
		override.UserID,
		override.PermissionID,
		granted,
		override.Reason,
		override.CreatedBy,
		override.CreatedAt,
		override.DataScope,
	)
	return err
}

func (r *userPermissionOverrideRepository) Update(override *model.UserPermissionOverride,) error {
	granted := 0
	if override.Granted {
		granted = 1
	}
	query := `UPDATE user_permission_overrides SET granted = :1, reason = :2, data_scope = :3 WHERE id = :4`
	_, err := r.db.Exec(
		query,
		granted,
		override.Reason,
		override.DataScope,
		override.ID,
	)
	return err
}

func (r *userPermissionOverrideRepository) Delete(id string,) error {
	query := `DELETE FROM user_permission_overrides WHERE id = :1`
	_, err := r.db.Exec(query, id,)
	return err
}