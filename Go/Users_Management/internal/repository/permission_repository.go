package repository

import (
	"database/sql"

	"go-rbac-system/internal/model"
)

type PermissionRepository interface {
	GetByID(id string) (*model.Permission, error)
	GetByFeatureCode(featureCode string) (*model.Permission, error)
	List() ([]model.Permission, error)
}

type permissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB,) PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}

func (r *permissionRepository) GetByID(id string,) (*model.Permission, error) {
	query := `SELECT id, feature_group, feature_code, action, description FROM permissions WHERE id = :1`
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
	query := `SELECT id, feature_group, feature_code, action, description FROM permissions WHERE feature_code = :1`
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

func (r *permissionRepository) List() ([]model.Permission, error) {
	query := `SELECT id, feature_group, feature_code, action, description FROM permissions ORDER BY feature_group, feature_code`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	permissions := make([]model.Permission, 0,)
	for rows.Next() {
		var item model.Permission
		err := rows.Scan(
			&item.ID,
			&item.FeatureGroup,
			&item.FeatureCode,
			&item.Action,
			&item.Description,
		)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, item,)
	}
	return permissions, rows.Err()
}