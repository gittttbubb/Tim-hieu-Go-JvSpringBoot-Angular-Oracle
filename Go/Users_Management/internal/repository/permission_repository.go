package repository

import (
	"database/sql"

	"go-rbac-system/internal/model"
)

type PermissionRepository interface {
	GetByID(id string) (*model.Permission, error)
	GetByFeatureCode(featureCode string) (*model.Permission, error)
	List(keyword string, offset int, pageSize int,) ([]model.Permission, int64, error)
	ListAll() ([]model.Permission, error)
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

func (r *permissionRepository) List(keyword string, offset int, pageSize int,) ([]model.Permission, int64, error) {
    var (
        rows  *sql.Rows
        err   error
        total int64
    )
    if keyword != "" {
        searchKeyword := "%" + keyword + "%"
        countQuery := `
            SELECT COUNT(*)
            FROM permissions
            WHERE
                LOWER(feature_group) LIKE LOWER(:keyword)
                OR LOWER(feature_code) LIKE LOWER(:keyword)
                OR LOWER(action) LIKE LOWER(:keyword)
                OR LOWER(description) LIKE LOWER(:keyword)
        `
        err = r.db.QueryRow(
            countQuery,
            sql.Named("keyword", searchKeyword),
        ).Scan(&total)
        if err != nil {
            return nil, 0, err
        }
        query := `
            SELECT
                id,
                feature_group,
                feature_code,
                action,
                description
            FROM permissions
            WHERE
                LOWER(feature_group) LIKE LOWER(:keyword)
                OR LOWER(feature_code) LIKE LOWER(:keyword)
                OR LOWER(action) LIKE LOWER(:keyword)
                OR LOWER(description) LIKE LOWER(:keyword)
            ORDER BY feature_group, feature_code
            OFFSET :offset ROWS
            FETCH NEXT :pageSize ROWS ONLY
        `
        rows, err = r.db.Query(
            query,
            sql.Named("keyword", searchKeyword),
            sql.Named("offset", offset),
            sql.Named("pageSize", pageSize),
        )
    } else {
        err = r.db.QueryRow(
            `SELECT COUNT(*) FROM permissions`,
        ).Scan(&total)
        if err != nil {
            return nil, 0, err
        }
        query := `
            SELECT
                id,
                feature_group,
                feature_code,
                action,
                description
            FROM permissions
            ORDER BY feature_group, feature_code
            OFFSET :offset ROWS
            FETCH NEXT :pageSize ROWS ONLY
        `
        rows, err = r.db.Query(
            query,
            sql.Named("offset", offset),
            sql.Named("pageSize", pageSize),
        )
    }
    if err != nil {
        return nil, 0, err
    }
    defer rows.Close()
    permissions := make([]model.Permission, 0)
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
            return nil, 0, err
        }
        permissions = append(permissions, item)
    }
    return permissions, total, rows.Err()
}

func (r *permissionRepository) ListAll() ([]model.Permission, error) {
	query := `SELECT id, feature_group, feature_code, action, description
		FROM permissions ORDER BY feature_group, feature_code`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	permissions := make([]model.Permission, 0)
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
		permissions = append(
			permissions,
			item,
		)
	}
	return permissions, rows.Err()
}