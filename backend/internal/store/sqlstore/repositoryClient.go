package sqlstore

import (
	"database/sql"
	"fmt"
	"sberapi/internal/model"
)

type ClientRepository struct {
	db *sql.DB
}

func (r *ClientRepository) All() ([]model.Client, error) {
	clients := make([]model.Client, 0, 10)

	rows, err := r.db.Query("SELECT id, name, inn FROM clients")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		client := model.Client{}
		err := rows.Scan(&client.ID, &client.Name, &client.Inn)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil

}

func (r *ClientRepository) Create(client *model.Client) error {
	if err := client.Validate(); err != nil {
		return err
	}

	err := r.db.QueryRow("INSERT INTO clients(name, inn)"+
		"VALUES ($1, $2) RETURNING id", client.Name, client.Inn).Scan(&client.ID)

	if err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) Find(id uint64) (*model.Client, error) {
	client := &model.Client{}
	err := r.db.QueryRow("SELECT id, name, inn FROM clients WHERE id = $1", id).
		Scan(&client.ID, &client.Name, &client.Inn)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Client with that id is not found")
		}
		return nil, err
	}
	return client, nil
}

func (r *ClientRepository) FindByEmployeeId(employeeId uint64) ([]*model.Client, error) {
	clients := make([]*model.Client, 0, 10)

	rows, err := r.db.Query("SELECT c.id, c.name, c.inn FROM clients_employee ce "+
		"INNER JOIN clients c ON c.id = ce.client_id "+
		"WHERE ce.employee_id = $1", employeeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		client := &model.Client{}
		err := rows.Scan(&client.ID, &client.Name, &client.Inn)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}
