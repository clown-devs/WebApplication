package store

import (
	"database/sql"
	"sberapi/internal/model"
)

type EmployeeRepository struct {
	db *sql.DB
}

func (r *EmployeeRepository) Create(u *model.Employee) (*model.Employee, error) {
	err := r.db.QueryRow(
		"INSERT INTO employees "+
			"(first_name, second_name, third_name, username, encrypted_password)"+
			"VALUES ($1, $2, $3, $4, $5) RETURNING id",
		u.Firstname, u.Secondname, u.Thirdname,
		u.Username, u.Encrypted_password).Scan(&u.ID)

	if err != nil {
		return nil, err
	}

	return u, nil

}

func (r *EmployeeRepository) FindByUsername(username string) (*model.Employee, error) {
	return nil, nil
}
