package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type Employee struct {
	ID                uint64 `json:"id"`
	Fullname          string `json:"fullname"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"` //Don't store in DB
	EncryptedPassword string `json:"-"`

	Clients   []*Client  `json:"client_list"`
	Role      *Role      `json:"role"`
	Direction *Direction `json:"direction"`
}

type EmployeeFilters struct {
	DirectionId uint64 `db:"direction_id"`
	ClientId    uint64 `db:"client_id"`
}

func (e *Employee) Validate() error {
	err := validation.ValidateStruct(e,
		validation.Field(&e.Username, validation.Required),
		validation.Field(&e.Password, validation.Required),
		validation.Field(&e.Password, validation.By(requiredIf(e.EncryptedPassword == ""))),
	)
	return err
}

func (e *Employee) Sanitize() {
	e.Password = ""
}

func (e *Employee) BeforeCreate() error {
	if len(e.Password) > 0 {
		enc, err := EncryptPassword(e.Password)
		if err != nil {
			return err
		}
		e.EncryptedPassword = enc
	}
	return nil
}

func (e *Employee) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(e.EncryptedPassword), []byte(password)) == nil
}

func EncryptPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

type Role struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	IsSuper        bool   `json:"is_super"`
	CanSeeMeetings bool   `json:"can_see_meetings"`
}

type Direction struct {
	ID   uint64 `json:"id"` //needed because sql can return null
	Name string `json:"name"`
}
