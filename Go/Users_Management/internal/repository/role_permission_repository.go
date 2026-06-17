package repository

import (
	"database/sql"

	"go-rbac-system/internal/model"
)

type RolePermissionRepository interface {
	GetByRoleAndPermission(roleID string, permissionID string,) (*model.RolePermission, error)
	ListByRoleID(roleID string) ([]model.RolePermission, error)
	Assign(permission *model.RolePermission) error
	Update(permission *model.RolePermission) error
	Delete(roleID string, permissionID string,) error
}

type rolePermissionRepository struct {
	db *sql.DB
}

func NewRolePermissionRepository(db *sql.DB,) RolePermissionRepository {
	return &rolePermissionRepository{
		db: db,
	}
}

func scanRolePermission(scanner interface {Scan(dest ...any) error},) (*model.RolePermission, error) {
	var item model.RolePermission
	var granted int
	err := scanner.Scan(&item.RoleID, &item.PermissionID, &granted, &item.DataScope,)
	if err != nil {
		return nil, err
	}
	item.Granted = granted == 1
	return &item, nil
}

func (r *rolePermissionRepository) GetByRoleAndPermission(roleID string,permissionID string,) (*model.RolePermission, error) {
	query := `SELECT role_id, permission_id, granted, data_scope FROM role_permissions WHERE role_id = :1 AND permission_id = :2`
	row := r.db.QueryRow(query, roleID, permissionID,)
	return scanRolePermission(row)
}

func (r *rolePermissionRepository) ListByRoleID(roleID string,) ([]model.RolePermission, error) {
	query := `SELECT role_id, permission_id, granted, data_scope FROM role_permissions WHERE role_id = :1 ORDER BY permission_id`
	rows, err := r.db.Query(query, roleID,)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]model.RolePermission, 0,)
	for rows.Next() {
		item, err := scanRolePermission(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *item,)
	}
	return result, rows.Err()
}

func (r *rolePermissionRepository) Assign(permission *model.RolePermission,) error {
	granted := 0
	if permission.Granted {
		granted = 1
	}
	query := `INSERT INTO role_permissions (role_id, permission_id, granted, data_scope) VALUES (:1, :2, :3, :4)`
	_, err := r.db.Exec(
		query,
		permission.RoleID,
		permission.PermissionID,
		granted,
		permission.DataScope,
	)
	return err
}

func (r *rolePermissionRepository) Update(permission *model.RolePermission,) error {
	granted := 0
	if permission.Granted {
		granted = 1
	}
	query := `UPDATE role_permissions SET granted = :1, data_scope = :2 WHERE role_id = :3AND permission_id = :4`
	_, err := r.db.Exec(
		query,
		granted,
		permission.DataScope,
		permission.RoleID,
		permission.PermissionID,
	)
	return err
}

func (r *rolePermissionRepository) Delete(roleID string,permissionID string,) error {
	query := `DELETE FROM role_permissions WHERE role_id = :1 AND permission_id = :2 `
	_, err := r.db.Exec(query, roleID, permissionID,)
	return err
}