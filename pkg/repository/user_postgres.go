package repository

import (
	"github.com/jmoiron/sqlx"
	"internship_avito/pkg/model"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int

	query := "INSERT INTO users DEFAULT VALUES RETURNING user_id"
	row := r.db.QueryRow(query)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
