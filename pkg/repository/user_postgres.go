package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"internship_avito/pkg/model"
	"time"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(model.User) (int, error) {
	var id int

	query := "INSERT INTO users DEFAULT VALUES RETURNING user_id"
	row := r.db.QueryRow(query)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) AddUserToSegments(user model.UserSegments) (string, error) {
	var input []string
	input = user.SegmentsName
	var id int
	query := "INSERT INTO segments_user (user_id, segments_name, time_create) VALUES ($1, $2, $3) RETURNING user_id"
	for _, j := range input {
		row := r.db.QueryRow(query, user.Id, j, time.Now())
		if err := row.Scan(&id); err != nil {
			return "error in db in func AddUserToSlug", err
		}
	}
	return "successes", nil
}

func (r *AuthPostgres) DeleteUserFromSegments(user model.UserSegments) (string, error) {
	var input []string
	input = user.SegmentsName
	var id int
	query := "DELETE FROM segments_user WHERE segments_name=$1 and user_id=$2 RETURNING user_id"
	for _, j := range input {
		row := r.db.QueryRow(query, j, user.Id)
		if err := row.Scan(&id); err != nil {
			return "error in db in func DeleteUserFromSlug", err
		}
	}
	return "successes", nil
}

func (r *AuthPostgres) GetUserSegments(user model.User) ([]string, error) {
	var segmentName string
	result := make([]string, 0)
	query := "SELECT segments_name FROM segments_user WHERE user_id=$1"
	row, err := r.db.Queryx(query, user.Id)
	if err != nil {
		logrus.Errorf("error in r.db.Queryx, err: %s", err)
	}

	for row.Next() {
		err := row.Scan(&segmentName)
		if err != nil {
			logrus.Errorf("error in row.Next, err: %s", err)
		}
		result = append(result, segmentName)
	}

	return result, nil
}
