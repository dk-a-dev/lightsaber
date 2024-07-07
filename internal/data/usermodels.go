package data

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"lightsabor.dkadev.net/internal/validator"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Activated bool      `json:"activated"`
	Version   int       `json:"-"`
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintextPass string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPass), 12)
	if err != nil {
		return err
	}
	p.plaintext = &plaintextPass
	p.hash = hash
	return nil
}

func (p *password) Matches(plaintextPass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPass))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePassword(v *validator.Validator, password string) {
	length := len(password)
	v.Check(password != "", "password", "must be provided")
	v.Check(length >= 8, "password", "must be at least 8 bytes long")
	v.Check(length <= 72, "password", "must not be more than 72 bytes long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")
	ValidateEmail(v, user.Email)
	if user.Password.plaintext != nil {
		ValidatePassword(v, *user.Password.plaintext)
	}
	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}
