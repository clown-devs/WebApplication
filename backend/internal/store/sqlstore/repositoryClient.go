package sqlstore

import (
	"database/sql"
	"fmt"
	"sberapi/internal/model"
)

type ClientRepository struct {
	db *sql.DB
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
