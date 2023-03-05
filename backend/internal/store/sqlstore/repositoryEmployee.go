package sqlstore

import (
	"database/sql"
	"fmt"
	"sberapi/internal/model"
)

const (
	userSQLString = `SELECT
	e.id,fullname,username,e.encrypted_password,
	role_id, r.name,r.is_super, r.can_see_meetings,
	d.id, d.name
	FROM employees e
	INNER JOIN roles r ON e.role_id = r.id
	INNER JOIN directions d ON e.direction_id = d.id`
)

type EmployeeRepository struct {
	db *sql.DB
}

func (r *EmployeeRepository) Create(e *model.Employee) error {
	if err := e.Validate(); err != nil {
		return err
	}

	if err := e.BeforeCreate(); err != nil {
		return err
	}

	err := r.db.QueryRow(
		"INSERT INTO employees "+
			"(fullname, username, encrypted_password)"+
			"VALUES ($1, $2, $3) RETURNING id",
		e.Fullname, e.Username, e.EncryptedPassword).Scan(&e.ID)

	if err != nil {
		return err
	}

	if err := e.Validate(); err != nil {
		return err
	}

	return nil

}

func (r *EmployeeRepository) Find(id uint64) (*model.Employee, error) {
	u, err := r.parseUser(r.db.QueryRow(userSQLString+" WHERE e.id = $1", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Employee with that id is not found")
		}
		return nil, err
	}
	return u, nil
}

func (r *EmployeeRepository) FindByUsername(username string) (*model.Employee, error) {
	u, err := r.parseUser(r.db.QueryRow(userSQLString+" WHERE e.username = $1", username))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Employee with that username is not found")
		}
		return nil, err
	}
	return u, nil
}

func (r *EmployeeRepository) parseUser(row *sql.Row) (*model.Employee, error) {

	u := &model.Employee{Role: &model.Role{}, Direction: &model.Direction{}}
	err := row.Scan(
		&u.ID, &u.Fullname, &u.Username, &u.EncryptedPassword,
		&u.Role.ID, &u.Role.Name, &u.Role.IsSuper, &u.Role.CanSeeMeetings,
		&u.Direction.ID, &u.Direction.Name,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}
