package sqlstore

import (
	"database/sql"
	"sberapi/internal/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db                  *sql.DB
	employeeRepository  *EmployeeRepository
	tokenRepository     *TokenRepository
	directionRepository *DirectionRepository
	clientRepository    *ClientRepository
	contactRepository   *ContactRepository
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

func (s *Store) Token() store.TokenRepository {
	if s.tokenRepository == nil {
		s.tokenRepository = &TokenRepository{
			db: s.db,
		}
	}
	return s.tokenRepository
}

func (s *Store) Direction() store.DirectionRepository {
	if s.directionRepository == nil {
		s.directionRepository = &DirectionRepository{
			db: s.db,
		}
	}
	return s.directionRepository
}

func (s *Store) Client() store.ClientRepository {
	if s.clientRepository == nil {
		s.clientRepository = &ClientRepository{
			db: s.db,
		}
	}

	return s.clientRepository
}

func (s *Store) Contact() store.ContactRepository {
	if s.contactRepository == nil {
		s.contactRepository = &ContactRepository{
			db: s.db,
		}
	}
	return s.contactRepository

}
