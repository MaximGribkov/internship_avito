package repository

import (
	"github.com/jmoiron/sqlx"
	"internship_avito/pkg/model"
)

type SegmentsPostgres struct {
	db *sqlx.DB
}

func NewSegmentsPostgres(db *sqlx.DB) *SegmentsPostgres {
	return &SegmentsPostgres{db: db}
}

func (s *SegmentsPostgres) CreateSegments(segments model.Segments) (int, error) {
	var input int
	query := "INSERT INTO segments (segments_name) VALUES ($1) RETURNING segments_id"
	row := s.db.QueryRow(query, segments.SegmentsName)
	if err := row.Scan(&input); err != nil {
		return 0, err
	}
	return input, nil
}

func (s *SegmentsPostgres) DeleteSegments(segments model.Segments) (string, error) {
	var input string

	//_, err := s.db.Exec("DELETE FROM segments WHERE segments_name=$1", segments)
	//if err != nil {
	//	return "bad", err
	//}
	//fmt.Println(segments)
	//return "ok", nil

	query := "DELETE FROM segments WHERE segments_name=$1"
	row := s.db.QueryRow(query, segments.SegmentsName)
	if err := row.Scan(&input); err != nil {
		return "successfully", nil
	}
	return "wrong name segment", nil
}
