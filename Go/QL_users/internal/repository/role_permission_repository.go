package repository

import (
	"database/sql"

	"go-user-management/internal/model"
)

type RolePermissionRepository interface {
	Create(rp *model.RolePermission,) error
	Delete(roleID string, permissionID string,) error
	GetByRoleID(roleID string,) ([]model.RolePermission, error)
	GetByRoleAndPermission(roleID string, permissionID string,) (*model.RolePermission, error)
}
type rolePermissionRepository struct {
	db *sql.DB
}

func NewRolePermissionRepository(db *sql.DB) RolePermissionRepository {
	return &rolePermissionRepository{
		db: db,
	}
}
func (r *rolePermissionRepository) Create(rp *model.RolePermission,) error {
	query := `INSERT INTO role_permissions(role_id, permission_id, granted, data_scope) VALUES(:1,:2,:3,:4)`
	_, err := r.db.Exec(
		query,
		rp.RoleID,
		rp.PermissionID,
		rp.Granted,
		rp.DataScope,
	)
	return err
}
func (r *rolePermissionRepository) Delete(roleID string, permissionID string,) error {
	query := `DELETE FROM role_permissions WHERE role_id = :1 AND permission_id = :2`
	_, err := r.db.Exec(
		query,
		roleID,
		permissionID,
	)
	return err
}
func (r *rolePermissionRepository) GetByRoleID(roleID string,) ([]model.RolePermission,error) {
	query := `SELECT role_id, permission_id, granted, data_scope FROM role_permissions WHERE role_id = :1`
	rows, err := r.db.Query(query, roleID,)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.RolePermission
	for rows.Next() {
		var rp model.RolePermission
		err := rows.Scan(
			&rp.RoleID,
			&rp.PermissionID,
			&rp.Granted,
			&rp.DataScope,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, rp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *rolePermissionRepository) GetByRoleAndPermission(roleID string, permissionID string,) (*model.RolePermission, error) {
	query := `
	SELECT
		role_id,
		permission_id,
		granted,
		data_scope
	FROM role_permissions
	WHERE role_id = :1
	AND permission_id = :2`
	var rp model.RolePermission
	err := r.db.QueryRow(query, roleID, permissionID,
	).Scan(
		&rp.RoleID,
		&rp.PermissionID,
		&rp.Granted,
		&rp.DataScope,
	)
	if err != nil {
		return nil, err
	}
	return &rp, nil
}