package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type Employee struct {
	ID                uint64
	Firstname         string
	Secondname        string
	Thirdname         string
	Username          string
	Password          string //Don't store in DB
	EncryptedPassword string
}

func (e *Employee) Validate() error {
	return validation.ValidateStruct(e,
		validation.Field(&e.Username, validation.Required),
		validation.Field(&e.Password, validation.Required),
	)
}

func (e *Employee) BeforeCreate() error {
	if len(e.Password) > 0 {
		enc, err := encryptPassword(e.Password)
		if err != nil {
			return err
		}
		e.EncryptedPassword = enc
	}
	return nil
}

func encryptPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}
