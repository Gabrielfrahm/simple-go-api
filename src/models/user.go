package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.formate(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	var errorsMessage []string
	if user.Name == "" {
		errorsMessage = append(errorsMessage, "the name is required and cannot be empty.")
	}
	if user.Nick == "" {
		errorsMessage = append(errorsMessage, "the nick is required and cannot be empty.")
	}
	if user.Email == "" {
		errorsMessage = append(errorsMessage, "the email is required and cannot be empty.")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		errorsMessage = append(errorsMessage, "the email should be a valid format")
	}

	if step == "register" && user.Password == "" {
		errorsMessage = append(errorsMessage, "the password is required and cannot be empty.")
	}

	if len(errorsMessage) > 0 {
		return errors.New(strings.Join(errorsMessage, " "))
	}

	return nil
}

func (user *User) formate(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if step == "register" {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}

	return nil
}
