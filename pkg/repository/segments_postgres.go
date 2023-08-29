package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"internship_avito/pkg/model"
	"time"
)

type SegmentsPostgres struct {
	db *sqlx.DB
}

func NewSegmentsPostgres(db *sqlx.DB) *SegmentsPostgres {
	return &SegmentsPostgres{db: db}
}

func (s *SegmentsPostgres) CreateSegments(segments model.Segments) (string, error) {
	var input string
	query := "INSERT INTO segments (segments_name, time_create) VALUES ($1, $2) RETURNING segments_name"
	row := s.db.QueryRow(query, segments.SegmentsName, time.Now())
	if err := row.Scan(&input); err != nil {
		return "", err
	}
	return input, nil
}

func (s *SegmentsPostgres) DeleteSegments(segments model.Segments) (string, error) {
	var input string

	query := "DELETE FROM segments WHERE segments_name=$1 RETURNING segments_name"
	row := s.db.QueryRow(query, segments.SegmentsName)
	if err := row.Scan(&input); err != nil {
		return "wrong name segment", nil
	}
	return "successes", nil
}

func (s *SegmentsPostgres) UserCountInSegment(segments model.Segments) (int, error) {
	count := 0

	query := "SELECT user_id FROM segments_user WHERE segments_name=$1"
	row, err := s.db.Queryx(query, segments.SegmentsName)
	if err != nil {
		logrus.Errorf("error in r.db.Queryx, err: %s", err)
	}

	for row.Next() {
		count++
	}

	return count, nil
}
