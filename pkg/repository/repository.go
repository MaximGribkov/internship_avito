package repository

import (
	"github.com/jmoiron/sqlx"
	"internship_avito/pkg/model"
)

type LogicsUser interface {
	CreateUser(user model.User) (int, error)
}

type LogicSegments interface {
	CreateSegments(segments model.Segments) (int, error)
	DeleteSegments(segments model.Segments) (string, error)
}

type Repository struct {
	LogicsUser
	LogicSegments
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		LogicsUser:    NewAuthPostgres(db),
		LogicSegments: NewSegmentsPostgres(db),
	}
}
