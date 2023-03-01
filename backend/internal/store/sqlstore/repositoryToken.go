package sqlstore

import (
	"database/sql"
	"fmt"
	"sberapi/internal/model"
)

type TokenRepository struct {
	db *sql.DB
}

func (r *TokenRepository) Create(employeeId uint64, salt string) (*model.Token, error) {
	token := model.NewToken(employeeId, salt)
	fmt.Println(token.Token)
	_, err := r.db.Exec(
		"INSERT INTO tokens"+
			"(employee_id, token)"+
			"VALUES ($1, $2)",
		token.EmployeeId, token.Token,
	)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (r *TokenRepository) Find(employeeId uint64) (*model.Token, error) {
	token := &model.Token{}

	if err := r.db.QueryRow("SELECT employee_id, token "+
		"FROM tokens WHERE employee_id = $1", employeeId).Scan(
		&token.EmployeeId,
		&token.Token,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return token, nil
}
