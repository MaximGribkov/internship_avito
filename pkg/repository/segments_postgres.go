package repository

import (
	"fmt"
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
	var inputCreate string
	var count int

	// создание сегмента
	queryCreateSegment := "INSERT INTO segments (segments_name, time_create) VALUES ($1, $2) RETURNING segments_name"
	rowCreateSegment := s.db.QueryRow(queryCreateSegment, segments.SegmentsName, time.Now())
	if err := rowCreateSegment.Scan(&inputCreate); err != nil {
		logrus.Errorf("error in func CreateSegments in rowCreateSegment.Scan(), err: %s", err)
		return "", err
	}

	err := s.db.Get(&count, "SELECT COUNT(*) FROM users")
	if err != nil {
		logrus.Errorf("error in func CreateSegments in s.db.Get(), err: %s", err)
	}

	// Вычисление процента
	percent := segments.Percent
	percentUser := (count * percent) / 100

	answer := fmt.Sprintf("Success, segment %v create. Count added user to segemnt = %v. Total user = %v",
		inputCreate, percentUser, count)

	// Добавление необходимого числа пользователей в сегмент
	queryAddUser := "INSERT INTO segments_user (user_id, segments_name, time_create) VALUES ($1, $2, $3) RETURNING user_id"
	rowIdUser, err := s.db.Queryx("SELECT * FROM users")

	if err != nil {
		logrus.Errorf("error in func CreateSegments in rowAddUser, err: %s", err)
		return "", err
	}
	if percentUser <= 0 { // для случая если хотим просто создать сегмент. запись в бд не будет
		return answer, nil
	}
	for rowIdUser.Next() {
		i := 0

		sliceId, err := rowIdUser.SliceScan() // возвращает слайс с id пользоватлей
		if err != nil {
			logrus.Errorf("error in func CreateSegments in rowIdUser.SliceScan(), err: %s", err)
			return "", err
		}

		_, err = s.db.Queryx(queryAddUser, sliceId[i], segments.SegmentsName, time.Now())
		if err != nil {
			logrus.Errorf("error in func CreateSegments in s.db.Queryx, err: %s", err)
			return "", err
		}
		percentUser--
		if percentUser == 0 { //данная конструкция использована чтобы избежать паники
			return answer, nil
		}
		i++
	}

	return answer, nil
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
