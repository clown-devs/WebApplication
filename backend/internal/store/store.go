package store

import (
	"database/sql"
	. "sberapi/internal/config"

	_ "github.com/lib/pq"
)

type Store struct {
	config             *Config
	db                 *sql.DB
	employeeRepository *EmployeeRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DbConnStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Employee() *EmployeeRepository {
	if s.employeeRepository != nil {
		return s.employeeRepository
	}

	s.employeeRepository = &EmployeeRepository{
		store: s,
	}
	return s.employeeRepository
}
