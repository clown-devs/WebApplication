package sqlstore

import (
	"database/sql"
	"sberapi/internal/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db                 *sql.DB
	employeeRepository *EmployeeRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Employee() store.EmployeeRepository {
	if s.employeeRepository == nil {
		s.employeeRepository = &EmployeeRepository{
			db: s.db,
		}
	}
	return s.employeeRepository
}
