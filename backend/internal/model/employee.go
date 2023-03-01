package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type Employee struct {
	ID                uint64 `json:"id"`
	Firstname         string `json:"first_name"`
	Secondname        string `json:"second_name"`
	Thirdname         string `json:"third_name"`
	Username          string `json:"username"`
	Password          string `json:"password,omitempty"` //Don't store in DB
	EncryptedPassword string `json:"-"`
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
