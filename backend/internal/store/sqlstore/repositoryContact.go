package sqlstore

import (
	"database/sql"
	"sberapi/internal/model"
)

const (
	contactSQLString = `
	SELECT con.id, con.fullname, con.phone, con.position, con.email,
	cl.id, cl.name, cl.inn
	FROM CONTACTS AS con
	INNER JOIN clients AS cl ON cl.id = con.client_id
	`
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

// TODO: implement
func (r *ContactRepository) All(filters *model.ContactFilters) ([]model.Contact, error) {
	contacts := make([]model.Contact, 0, 10)
	rows, err := r.db.Query(contactSQLString)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		contact, err := r.parseContact(rows)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, *contact)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	//FIXME:
	//query, values := r.filterQuery(userSQLString, filters)
	return contacts, nil
}

// Parses sql row to employee model
func (r *ContactRepository) parseContact(row Scannable) (*model.Contact, error) {
	contact := &model.Contact{}
	contact.Client = &model.Client{}

	err := row.Scan(&contact.ID, &contact.Fullname, &contact.Phone, &contact.Position, &contact.Email,
		&contact.Client.ID, &contact.Client.Name, &contact.Client.Inn)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (r *ContactRepository) filterQuery(query string, filters *model.ContactFilters) (string, []interface{}) {
	values := make([]interface{}, 0, 1)

	return query, values
}
