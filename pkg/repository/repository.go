package repository

import (
	"github.com/jmoiron/sqlx"
	"internship_avito/pkg/model"
)

type Logics interface {
	CreateUser(user model.User) (int, error)
}

type Repository struct {
	Logics
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Logics: NewAuthPostgres(db),
	}
}
