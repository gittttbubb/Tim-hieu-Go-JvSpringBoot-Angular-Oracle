package repository

import (
	"database/sql"

	"go-user-management/internal/model"
)

type UserPermissionOverrideRepository interface {
	Create(override *model.UserPermissionOverride,) error
	Update(override *model.UserPermissionOverride,) error
	Delete(userID string, permissionID string,) error
	GetByUserID(userID string,) ([]model.UserPermissionOverride,error)
	GetByUserAndPermission(userID string, permissionID string,) (*model.UserPermissionOverride, error)
}
type userPermissionOverrideRepository struct {
	db *sql.DB
}

func NewUserPermissionOverrideRepository(db *sql.DB) UserPermissionOverrideRepository {
	return &userPermissionOverrideRepository{
		db: db,
	}
}
func (r *userPermissionOverrideRepository) Create(override *model.UserPermissionOverride,) error {
	query := `INSERT INTO user_permission_overrides(
		id,
		user_id,
		permission_id,
		granted,
		reason,
		created_by,
		data_scope
	) VALUES(:1,:2,:3,:4,:5,:6,:7)`

	_, err := r.db.Exec(
		query,
		override.ID,
		override.UserID,
		override.PermissionID,
		override.Granted,
		override.Reason,
		override.CreatedBy,
		override.DataScope,
	)
	return err
}

func (r *userPermissionOverrideRepository) Update(override *model.UserPermissionOverride,) error {
	query := `
	UPDATE user_permission_overrides
	SET
		granted = :1,
		reason = :2,
		data_scope = :3
	WHERE id = :4
	`
	_, err := r.db.Exec(
		query,
		override.Granted,
		override.Reason,
		override.DataScope,
		override.ID,
	)
	return err
}

func (r *userPermissionOverrideRepository) Delete(userID string, permissionID string,) error {
	query := `DELETE FROM user_permission_overrides WHERE user_id = :1 AND permission_id = :2`
	_, err := r.db.Exec(query, userID, permissionID,)
	return err
}

func (r *userPermissionOverrideRepository) GetByUserID(userID string,) ([]model.UserPermissionOverride,error) {
	query := `
	SELECT
		id,
		user_id,
		permission_id,
		granted,
		reason,
		created_by,
		created_at,
		data_scope
	FROM user_permission_overrides WHERE user_id = :1
	`
	rows, err := r.db.Query(query, userID,)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.UserPermissionOverride
	for rows.Next() {
		var item model.UserPermissionOverride
		err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.PermissionID,
			&item.Granted,
			&item.Reason,
			&item.CreatedBy,
			&item.CreatedAt,
			&item.DataScope,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result,nil
}

func (r *userPermissionOverrideRepository) GetByUserAndPermission(userID string, permissionID string,) (*model.UserPermissionOverride, error) {
	query := `
	SELECT
		id,
		user_id,
		permission_id,
		granted,
		reason,
		created_by,
		created_at,
		data_scope
	FROM user_permission_overrides
	WHERE user_id = :1
	AND permission_id = :2
	`
	var item model.UserPermissionOverride
	err := r.db.QueryRow(query, userID, permissionID,
	).Scan(
		&item.ID,
		&item.UserID,
		&item.PermissionID,
		&item.Granted,
		&item.Reason,
		&item.CreatedBy,
		&item.CreatedAt,
		&item.DataScope,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}