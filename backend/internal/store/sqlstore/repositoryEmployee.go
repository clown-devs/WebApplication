package sqlstore

import (
	"database/sql"
	"fmt"
	"sberapi/internal/model"
	"sberapi/internal/store"
)

const (
	userSQLString = `SELECT
	e.id,fullname,username,e.encrypted_password,
	role_id, r.name,r.is_super, r.can_see_meetings,
	d.id, d.name
	FROM employees e
	INNER JOIN roles r ON e.role_id = r.id
	LEFT JOIN directions d ON e.direction_id = d.id `
)

type EmployeeRepository struct {
	db               *sql.DB
	clientRepository store.ClientRepository
}

func (r *EmployeeRepository) All(filters *model.EmployeeFilters) ([]model.Employee, error) {
	employees := make([]model.Employee, 0, 10)
	query, values := r.filterQuery(userSQLString, filters)

	rows, err := r.db.Query(query, values...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		e, err := r.parseEmployee(rows)
		if err != nil {
			return nil, err
		}
		employees = append(employees, *e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *EmployeeRepository) Create(e *model.Employee) error {
	if err := e.Validate(); err != nil {
		return err
	}

	if err := e.BeforeCreate(); err != nil {
		return err
	}
	e.Role = &model.Role{}
	e.Clients = make([]*model.Client, 0)
	err := r.db.QueryRow(
		`WITH inserted_employee AS (
			INSERT INTO employees(fullname, username, encrypted_password)
			VALUES ($1, $2, $3)
			RETURNING employees.id, role_id)

			SELECT ie.id, roles.id, roles.name, roles.is_super, roles.can_see_meetings 
			FROM inserted_employee ie
				INNER JOIN roles on ie.role_id = roles.id`,
		e.Fullname, e.Username, e.EncryptedPassword).Scan(
		&e.ID,
		&e.Role.ID, &e.Role.Name, &e.Role.IsSuper, &e.Role.CanSeeMeetings)

	if err != nil {
		return err
	}

	if err := e.Validate(); err != nil {
		return err
	}

	return nil

}

func (r *EmployeeRepository) Delete(id uint64) error {
	_, err := r.db.Exec("DELETE FROM employees WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// Returns only public fields (without password and username)
func (r *EmployeeRepository) Find(id uint64) (*model.Employee, error) {
	e, err := r.parseEmployee(
		r.db.QueryRow(userSQLString+" WHERE e.id = $1", id),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Employee with that id is not found")
		}
		return nil, err
	}

	e.Username = ""
	return e, nil
}

func (r *EmployeeRepository) FindByUsername(username string) (*model.Employee, error) {
	u, err := r.parseEmployee(
		r.db.QueryRow(userSQLString+" WHERE e.username = $1", username),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Employee with that username is not found")
		}
		return nil, err
	}
	return u, nil
}

func (r *EmployeeRepository) FindByToken(token string) (*model.Employee, error) {
	u, err := r.parseEmployee(
		r.db.QueryRow(userSQLString+
			"WHERE e.id = "+
			"(select employee_id from tokens t where t.token = $1)", token),
	)
	if err != nil {
		return nil, err
	}

	return u, nil

}

type Scannable interface {
	Scan(dest ...any) error
}

// Parses sql row to employee model
func (r *EmployeeRepository) parseEmployee(row Scannable) (*model.Employee, error) {
	directionID, directionName := sql.NullInt64{}, sql.NullString{}
	u := &model.Employee{Role: &model.Role{}, Direction: nil}

	err := row.Scan(
		&u.ID, &u.Fullname, &u.Username, &u.EncryptedPassword,
		&u.Role.ID, &u.Role.Name, &u.Role.IsSuper, &u.Role.CanSeeMeetings,
		&directionID, &directionName,
	)

	if err != nil {
		return nil, err
	}

	if directionID.Valid {
		u.Direction = &model.Direction{ID: uint64(directionID.Int64), Name: directionName.String}
	}
	u.Clients, err = r.clientRepository.FindByEmployeeId(u.ID)

	return u, nil
}

// returns query with filters
func (r *EmployeeRepository) filterQuery(query string, filters *model.EmployeeFilters) (string, []interface{}) {
	values := []interface{}{}
	query += "WHERE 1=1 "
	if filters.ClientId != 0 {
		values = append(values, filters.ClientId)
		query += fmt.Sprintf("AND e.id in (SELECT employee_id from clients_employee where client_id = $%d) ", len(values))
	}
	if filters.DirectionId != 0 {
		values = append(values, filters.DirectionId)
		query += fmt.Sprintf("AND d.id = $%d ", len(values))
	}
	fmt.Println(query)
	return query, values
}

// cannot update username & password
func (r *EmployeeRepository) UpdatePublic(e *model.Employee) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var directionId *uint64
	if e.Direction != nil {
		directionId = &e.Direction.ID
	} else {
		directionId = nil
	}

	_, err = tx.Exec(
		`
			UPDATE employees
			SET fullname = $1, direction_id = $2
			WHERE id = $3
		`, e.Fullname, directionId, e.ID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM clients_employee WHERE employee_id = $1", e.ID)
	if err != nil {
		return err
	}

	for _, client := range e.Clients {
		_, err = tx.Exec("INSERT INTO clients_employee(client_id, employee_id) VALUES ($1, $2)", client.ID, e.ID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
