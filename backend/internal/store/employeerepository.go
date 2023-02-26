package store

import (
	"database/sql"
	"sberapi/internal/model"
)

type EmployeeRepository struct {
	db *sql.DB
}

func (r *EmployeeRepository) Create(e *model.Employee) (*model.Employee, error) {
	if err := e.Validate(); err != nil {
		return nil, err
	}

	if err := e.BeforeCreate(); err != nil {
		return nil, err
	}

	err := r.db.QueryRow(
		"INSERT INTO employees "+
			"(first_name, second_name, third_name, username, encrypted_password)"+
			"VALUES ($1, $2, $3, $4, $5) RETURNING id",
		e.Firstname, e.Secondname, e.Thirdname,
		e.Username, e.EncryptedPassword).Scan(&e.ID)

	if err != nil {
		return nil, err
	}

	if err := e.Validate(); err != nil {
		return nil, err
	}

	return e, nil

}

func (r *EmployeeRepository) FindByUsername(username string) (*model.Employee, error) {
	u := &model.Employee{}
	if err := r.db.QueryRow("SELECT * FROM employees WHERE username = $1", username).Scan(
		&u.ID,
		&u.Firstname,
		&u.Secondname,
		&u.Thirdname,
		&u.Username,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}
	return u, nil
}
