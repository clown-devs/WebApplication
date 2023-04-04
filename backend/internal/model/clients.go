package model

import validation "github.com/go-ozzo/ozzo-validation"

type Client struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Inn  string `json:"inn"` // max length: 12
}

func (c *Client) Validate() error {
	err := validation.ValidateStruct(c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Inn, validation.Required, validation.Length(1, 12)),
	)
	return err
}

type Contact struct {
	ID       uint64  `json:"id"`
	Fullname string  `json:"fullname"`
	Phone    *string `json:"phone"`
	Position *string `json:"position"`
	Email    *string `json:"email"`
	Client   *Client `json:"client"`
}

type ContactFilters struct {
	ClientId uint64
}

func (c *Contact) Validate() error {
	err := validation.ValidateStruct(c,
		validation.Field(&c.Fullname, validation.Required),
	)
	return err
}
