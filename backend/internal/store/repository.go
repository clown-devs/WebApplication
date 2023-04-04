package store

import "sberapi/internal/model"

type EmployeeRepository interface {
	Create(*model.Employee) error
	Find(uint64) (*model.Employee, error)
	FindByUsername(string) (*model.Employee, error)
	FindByToken(string) (*model.Employee, error)
	All(*model.EmployeeFilters) ([]model.Employee, error)
	UpdatePublic(*model.Employee) error
	//TODO:
	//Delete
}

type TokenRepository interface {
	Create(uint64, string) (*model.Token, error) // FIXME: provide token model
	// Not returning an error if not found
	// Just nil
	Find(uint64) (*model.Token, error)
	Exists(string) (bool, error)
	//TODO:
	//Destroy
}

type DirectionRepository interface {
	Create(*model.Direction) error
	Find(uint64) (*model.Direction, error)
	All() ([]model.Direction, error)
	//TODO:
	//Delete
}

type ClientRepository interface {
	Create(*model.Client) error
	Find(uint64) (*model.Client, error)
	All() ([]model.Client, error)
	FindByEmployeeId(uint64) ([]*model.Client, error)
}

type ContactRepository interface {
	Create(*model.Contact) error //FIXME:
	All(*model.ContactFilters) ([]model.Contact, error)
	//Find(uint64) (*model.Contact, error)
	//All
}
