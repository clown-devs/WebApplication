package sqlstore

import (
	"database/sql"
	"sberapi/internal/model"
)

type ContactRepository struct {
	db *sql.DB
}

func (r *ContactRepository) Create(contact *model.Contact) error {
	if err := contact.Validate(); err != nil {
		return err
	}

	err := r.db.QueryRow("INSERT INTO contacts(fullname, phone, email, position, client_id)"+
		"VALUES ($1, $2, $3, $4, $5) RETURNING id",
		contact.Fullname, contact.Phone, contact.Email, contact.Position, contact.Client.ID).
		Scan(&contact.ID)

	if err != nil {
		return err
	}
	return nil
}
