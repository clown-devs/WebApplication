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

	err := r.db.QueryRow(
		"WITH inserted_contact AS ("+
			"INSERT INTO contacts(fullname, phone, email, position, client_id)"+
			"VALUES ($1, $2, $3, $4, $5)"+
			"RETURNING contacts.id, contacts.client_id) "+
			"SELECT ic.id, cl.id, cl.name, cl.inn  FROM inserted_contact ic "+
			"INNER JOIN clients cl on ic.client_id = cl.id",
		contact.Fullname, contact.Phone, contact.Email, contact.Position, contact.Client.ID).
		Scan(&contact.ID, &contact.Client.ID, &contact.Client.Name, &contact.Client.Inn)

	if err != nil {
		return err
	}
	return nil
}
