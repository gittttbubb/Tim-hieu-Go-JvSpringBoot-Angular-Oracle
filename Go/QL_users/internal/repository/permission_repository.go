package repository

import (
	"database/sql"

	"go-user-management/internal/model"
)

type PermissionRepository interface {
	GetAll() ([]model.Permission, error)
	GetByID(id string) (*model.Permission, error)
	GetByFeatureCode(featureCode string) (*model.Permission, error)
	GetPermissionsByRole(roleID string) ([]string, error)
	GetPermissionsByRoleWithScope(roleID string,) ([]model.RolePermissionDetail, error)
}

type permissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}
func (r *permissionRepository) GetAll() ([]model.Permission, error) {
    query := `SELECT id, feature_group, feature_code, action, description  FROM permissions`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var permissions []model.Permission
    for rows.Next() {
        var p model.Permission
        err := rows.Scan(
			&p.ID, 
			&p.FeatureGroup,
			&p.FeatureCode,
			&p.Action,
			&p.Description,
		)
        if err != nil {
            return nil, err
        }
        permissions = append(permissions, p)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return permissions, nil
}

func (r *permissionRepository) GetByID(id string,) (*model.Permission, error) {
	query := `
	SELECT
		id,
		feature_group,
		feature_code,
		action,
		description
	FROM permissions
	WHERE id = :1`

	var permission model.Permission
	err := r.db.QueryRow(query, id,
	).Scan(
		&permission.ID,
		&permission.FeatureGroup,
		&permission.FeatureCode,
		&permission.Action,
		&permission.Description,
	)
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) GetByFeatureCode(featureCode string,) (*model.Permission, error) {
	query := `
	SELECT
		id,
		feature_group,
		feature_code,
		action,
		description
	FROM permissions
	WHERE feature_code = :1`
	var permission model.Permission
	err := r.db.QueryRow(query, featureCode,
	).Scan(
		&permission.ID,
		&permission.FeatureGroup,
		&permission.FeatureCode,
		&permission.Action,
		&permission.Description,
	)
	if err != nil {
		return nil, err
	}
	return &permission, nil
}
func (r *permissionRepository) GetPermissionsByRole(roleID string,) ([]string, error) {
	query := `
	SELECT p.feature_code
	FROM permissions p
	JOIN role_permissions rp
		ON rp.permission_id = p.id
	WHERE rp.role_id = :1
	AND rp.granted = 1`
	rows, err := r.db.Query(query, roleID,)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []string
	for rows.Next() {
		var code string
		if err := rows.Scan(&code); err != nil {
			return nil, err
		}
		permissions = append(
			permissions,
			code,
		)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *permissionRepository) GetPermissionsByRoleWithScope(roleID string,) ([]model.RolePermissionDetail, error) {
	query := `
	SELECT
		p.id,
		p.feature_code,
		rp.granted,
		rp.data_scope
	FROM role_permissions rp
	JOIN permissions p
	ON p.id = rp.permission_id
	WHERE rp.role_id = :1`
	rows, err := r.db.Query(query, roleID,)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []model.RolePermissionDetail
	for rows.Next() {
		var permission model.RolePermissionDetail
		var granted int
		err := rows.Scan(
			&permission.PermissionID,
			&permission.FeatureCode,
			&granted,
			&permission.DataScope,
		)
		if err != nil {
			return nil, err
		}
		permission.Granted = granted == 1
		permissions = append(permissions, permission,)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return permissions, nil
}