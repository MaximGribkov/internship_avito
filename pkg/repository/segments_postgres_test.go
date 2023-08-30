package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"internship_avito/pkg/model"
)

func TestSegmentsPostgres_CreateSegments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock DB: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	segmentsRepo := NewSegmentsPostgres(sqlxDB)

	segments := model.Segments{
		SegmentsName: "TestSegment",
	}

	now := time.Now()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO segments (segments_name, time_create) VALUES ($1, $2) RETURNING segments_name")).
		WithArgs(segments.SegmentsName, now).
		WillReturnRows(sqlmock.NewRows([]string{"segments_name"}).AddRow(segments.SegmentsName))

	result, err := segmentsRepo.CreateSegments(segments)

	assert.NoError(t, err)
	assert.Equal(t, segments.SegmentsName, result)
}

func TestSegmentsPostgres_DeleteSegments(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock DB: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	segmentsRepo := NewSegmentsPostgres(sqlxDB)

	segments := model.Segments{
		SegmentsName: "TestSegment",
	}

	mock.ExpectQuery(regexp.QuoteMeta("DELETE FROM segments WHERE segments_name=$1")).
		WithArgs(segments.SegmentsName).
		WillReturnRows(sqlmock.NewRows([]string{"segments_name"}).AddRow(segments.SegmentsName))

	result, err := segmentsRepo.DeleteSegments(segments)
	assert.NoError(t, err)
	assert.Equal(t, "successes", result)
}

func TestSegmentsPostgres_UserCountInSegment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock DB: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	segmentsRepo := NewSegmentsPostgres(sqlxDB)

	segments := model.Segments{
		SegmentsName: "TestSegment",
	}

	rows := sqlmock.NewRows([]string{"user_id"}).AddRow(1).AddRow(2).AddRow(3)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT user_id FROM segments_user WHERE segments_name=$1")).
		WithArgs(segments.SegmentsName).
		WillReturnRows(rows)

	count, err := segmentsRepo.UserCountInSegment(segments)
	assert.NoError(t, err)
	assert.Equal(t, 3, count)
}
