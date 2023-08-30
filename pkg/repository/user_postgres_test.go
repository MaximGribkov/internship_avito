package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"internship_avito/pkg/model"
)

func TestAuthPostgres_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock DB: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	authRepo := NewAuthPostgres(sqlxDB)

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users DEFAULT VALUES RETURNING user_id")).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(1))

	id, err := authRepo.CreateUser(model.User{})
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestAuthPostgres_AddUserToSegments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock DB: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	authRepo := NewAuthPostgres(sqlxDB)

	userSegments := model.UserSegments{
		Id:           1,
		SegmentsName: []string{"Segment1", "Segment2"},
	}

	for i := range userSegments.SegmentsName {
		mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO segments_user (user_id, segments_name, time_create) VALUES ($1, $2, $3) RETURNING user_id")).
			WithArgs(userSegments.Id, userSegments.SegmentsName[i], time.Now()).
			WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(1))
	}

	result, err := authRepo.AddUserToSegments(userSegments)
	assert.NoError(t, err)
	assert.Equal(t, "successes", result)
}

func TestAuthPostgres_DeleteUserFromSegments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock DB: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	authRepo := NewAuthPostgres(sqlxDB)

	userSegments := model.UserSegments{
		Id:           1,
		SegmentsName: []string{"Segment1", "Segment2"},
	}

	for i := range userSegments.SegmentsName {
		mock.ExpectQuery(regexp.QuoteMeta("DELETE FROM segments_user WHERE segments_name=$1 and user_id=$2 RETURNING user_id")).
			WithArgs(userSegments.SegmentsName[i], userSegments.Id).
			WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(1))
	}

	result, err := authRepo.DeleteUserFromSegments(userSegments)
	assert.NoError(t, err)
	assert.Equal(t, "successes", result)
}

func TestAuthPostgres_GetUserSegments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock DB: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	authRepo := NewAuthPostgres(sqlxDB)

	user := model.User{
		Id: 1,
	}

	rows := sqlmock.NewRows([]string{"segments_name"}).AddRow("Segment1").AddRow("Segment2")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT segments_name FROM segments_user WHERE user_id=$1")).
		WithArgs(user.Id).
		WillReturnRows(rows)

	result, err := authRepo.GetUserSegments(user)
	assert.NoError(t, err)
	assert.Equal(t, []string{"Segment1", "Segment2"}, result)
}
