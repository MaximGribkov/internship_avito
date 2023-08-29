package repository

import (
	"github.com/jmoiron/sqlx"
	"internship_avito/pkg/model"
)

type LogicsUser interface {
	CreateUser(user model.User) (int, error)
	AddUserToSegments(user model.UserSegments) (string, error)
	DeleteUserFromSegments(user model.UserSegments) (string, error)
	GetUserSegments(user model.User) ([]string, error)
}

type LogicSegments interface {
	CreateSegments(segments model.Segments) (string, error)
	DeleteSegments(segments model.Segments) (string, error)
	UserCountInSegment(segments model.Segments) (int, error)
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
