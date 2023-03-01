package store

import "sberapi/internal/model"

type EmployeeRepository interface {
	Create(*model.Employee) error
	Find(uint64) (*model.Employee, error)
	FindByUsername(string) (*model.Employee, error)
	//TODO:
	//Delete
}

type TokenRepository interface {
	Create(uint64, string) (*model.Token, error) // employee id, salt
	// Not returning an error if not found
	Find(uint64) (*model.Token, error)
	//TODO:
	//Destroy
}
