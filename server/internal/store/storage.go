package store

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Storage struct {
	Project interface {
		Insert(project *Project) error
		Get(id int64) (*Project, error)
		Update(project *Project) error
		Delete(id int64) error
		GetByName(url string) (*Project, error)
	}

	Users interface {
		Insert(user *User) error
		GetByEmail(email string) (*User, error)
		GetForToken(tokenScope, tokenPlaintext string) (*User, error)
	}

	Tokens TokenStore
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Project: &ProjectStore{db},
		Tokens:  TokenStore{db},
		Users:   &UsersStore{db},
	}
}
