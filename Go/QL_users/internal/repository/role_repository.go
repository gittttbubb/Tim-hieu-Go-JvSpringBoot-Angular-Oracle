package repository


import (
	"database/sql"

	"go-user-management/internal/model"
)


type RoleRepository interface {
	GetByID(id string) (*model.Role, error)
	GetByName(name string) (*model.Role, error)
	Create(role *model.Role) error
	Update(role *model.Role) error
	Delete(id string) error
	List() ([]model.Role, error)
}

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) Create(role *model.Role,) error {
	query := `INSERT INTO roles(id, name, display_name,description) VALUES(:1,:2,:3,:4)`
	_, err := r.db.Exec(
		query,
		role.ID,
		role.Name,
		role.DisplayName,
		role.Description,
	)
	return err
}
func (r *roleRepository) Update(role *model.Role) error {
	query := `UPDATE roles SET id = :1, name = :2, display_name = :3, description = :4, WHERE id = :5`
	_, err := r.db.Exec(
		query,
		role.ID,
		role.Name,
		role.DisplayName,
		role.Description,
	)
	return err
}

func (r *roleRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM roles WHERE id = :1`, id)
	return err
}

func (r *roleRepository) List() ([]model.Role, error,) {
	rows, err := r.db.Query(`SELECT id, name, display_name, description FROM roles`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []model.Role
	for rows.Next() {
		var role model.Role
		rows.Scan(
			&role.ID,
			&role.Name,
			&role.DisplayName,
			&role.Description,
		)
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *roleRepository) GetByID(id string) (*model.Role, error) {
	query := `SELECT id, name, display_name, description FROM roles WHERE id = :1`
	var role model.Role
	err := r.db.QueryRow(query, id,).Scan(
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
func (r *roleRepository) GetByName(name string) (*model.Role, error) {
	query := `SELECT id, name, display_name, description FROM roles WHERE name = :1`
	var role model.Role
	err := r.db.QueryRow(query, name).Scan(
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