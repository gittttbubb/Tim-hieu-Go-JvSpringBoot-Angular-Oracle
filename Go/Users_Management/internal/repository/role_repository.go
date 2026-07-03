package repository

import (
	"database/sql"

	"go-rbac-system/internal/model"
)

type RoleRepository interface {
	GetByID(id string) (*model.Role, error)
	GetByName(name string) (*model.Role, error)
	List(keyword string, offset int, pageSize int,) ([]model.Role, int64, error)
	ListAll() ([]model.Role, error)
	Create(role *model.Role) error
	Update(role *model.Role) error
	Delete(id string) error
}

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB,) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) GetByID(id string,) (*model.Role, error) {
	query := `SELECT id, name, display_name, description FROM roles WHERE id = :1`
	var role model.Role
	err := r.db.QueryRow(query, id,
	).Scan(
		&role.ID,
		&role.Name,
		&role.DisplayName,
		&role.Description,
	)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetByName(name string,) (*model.Role, error) {
	query := `SELECT id, name, display_name, description FROM roles WHERE name = :1`
	var role model.Role
	err := r.db.QueryRow(query, name,
	).Scan(
		&role.ID,
		&role.Name,
		&role.DisplayName,
		&role.Description,
	)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) List(keyword string, offset int, pageSize int,) ([]model.Role, int64, error) {
	var (
		rows  *sql.Rows
		err   error
		total int64
	)
	if keyword != "" {
		searchKeyword := "%" + keyword + "%"
		countQuery := `
			SELECT COUNT(*)
			FROM roles
			WHERE
				LOWER(name) LIKE LOWER(:keyword)
				OR LOWER(display_name) LIKE LOWER(:keyword)
				OR LOWER(description) LIKE LOWER(:keyword)
		`
		err = r.db.QueryRow(countQuery, sql.Named("keyword", searchKeyword)).Scan(&total)
		if err != nil {
			return nil, 0, err
		}
		query := `SELECT id, name, display_name, description FROM roles
			WHERE LOWER(name) LIKE LOWER(:keyword)
			OR LOWER(display_name) LIKE LOWER(:keyword)
			OR LOWER(description) LIKE LOWER(:keyword)
			ORDER BY name
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
		err = r.db.QueryRow(`SELECT COUNT(*) FROM roles`).Scan(&total)
		if err != nil {
			return nil, 0, err
		}
		query := `SELECT id, name, display_name, description FROM roles 
		ORDER BY name OFFSET :offset ROWS FETCH NEXT :pageSize ROWS ONLY`
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
	roles := make([]model.Role, 0)
	for rows.Next() {
		var role model.Role
		err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.DisplayName,
			&role.Description,
		)
		if err != nil {
			return nil, 0, err
		}
		roles = append(roles, role)
	}
	return roles, total, rows.Err()
}

func (r *roleRepository) ListAll() ([]model.Role, error) {
	query := `SELECT id, name, display_name, description
		FROM roles ORDER BY name`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	roles := make([]model.Role, 0)
	for rows.Next() {
		var role model.Role
		err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.DisplayName,
			&role.Description,
		)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, rows.Err()
}

func (r *roleRepository) Create(role *model.Role,) error {
	query := `INSERT INTO roles (id, name, display_name, description)
		VALUES (:1, :2, :3, :4)`
	_, err := r.db.Exec(
		query,
		role.ID,
		role.Name,
		role.DisplayName,
		role.Description,
	)
	return err
}

func (r *roleRepository) Update(role *model.Role,) error {
	query := `UPDATE roles SET name = :1, display_name = :2, description = :3 WHERE id = :4`
	_, err := r.db.Exec(
		query,
		role.Name,
		role.DisplayName,
		role.Description,
		role.ID,
	)
	return err
}

func (r *roleRepository) Delete(id string,) error {
	query := `DELETE FROM roles WHERE id = :1`
	_, err := r.db.Exec(query, id,)
	return err
}