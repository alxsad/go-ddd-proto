package user

import (
	"database/sql/driver"
	"fmt"
	"github.com/google/uuid"
)

type Users []UserEntity

type UserEntity struct {
	ID    uuid.UUID        `db:"id"`
	Name  string           `db:"name"`
	Email EmailValueObject `db:"email"`
}

func NewUser(name string, email string) (user UserEntity, err error) {
	if name == "" {
		return user, fmt.Errorf("invalid name: %s", name)
	}
	if user.Email, err = NewEmail(email); err != nil {
		return user, err
	}
	user.ID = uuid.New()
	user.Name = name
	return
}

type EmailValueObject struct {
	Email string
}

func NewEmail(value string) (email EmailValueObject, err error) {
	if value == "" {
		return email, fmt.Errorf("invalid email: %s", value)
	}
	email.Email = value
	return
}

func (e EmailValueObject) String() string {
	return e.Email
}

func (e EmailValueObject) Value() (driver.Value, error) {
	return e.Email, nil
}

func (e *EmailValueObject) Scan(src interface{}) error {
	if src == nil {
		*e = EmailValueObject{}
	} else {
		*e = EmailValueObject{src.(string)}
	}
	return nil
}
