package sqlstore

import (
	"database/sql"
	"fmt"
	"sberapi/internal/model"
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
	u := &model.Employee{}

	if err := r.db.QueryRow("SELECT * FROM employees WHERE id = $1", id).Scan(
		&u.ID,
		&u.Fullname,
		&u.Username,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Employee with that id is not found")
		}
		return nil, err
	}
	return u, nil
}

func (r *EmployeeRepository) FindByUsername(username string) (*model.Employee, error) {
	u := &model.Employee{}
	if err := r.db.QueryRow("SELECT * FROM employees WHERE username = $1", username).Scan(
		&u.ID,
		&u.Fullname,
		&u.Username,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Employee with that username is not found")
		}
		return nil, err
	}
	return u, nil
}
