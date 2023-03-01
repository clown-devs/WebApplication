package store

type Store interface {
	Employee() EmployeeRepository
	Token() TokenRepository
}
