package store

import "sberapi/internal/model"

type EmployeeRepository interface {
	Create(*model.Employee) error
	Find(uint64) (*model.Employee, error)
	FindByUsername(string) (*model.Employee, error)
}
