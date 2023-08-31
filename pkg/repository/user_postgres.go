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

// CreateUser Функция создания пользователя
func (r *AuthPostgres) CreateUser(model.User) (int, error) {
	var id int

	query := "INSERT INTO users DEFAULT VALUES RETURNING user_id"
	row := r.db.QueryRow(query)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// AddUserToSegments Функция добавления пользователя в сегмент или нескольок сегментов, а так же автоматическое удаление при необходимости
func (r *AuthPostgres) AddUserToSegments(user model.UserSegments) (string, error) {
	var input []string
	input = user.SegmentsName
	var id int

	deleteTime := new(int64)          // создаем переменную формата int64
	now := float64(time.Now().Unix()) // фиксируем время с начала эпохи в формате float64

	query := "INSERT INTO segments_user (user_id, segments_name, time_create,time_delete) VALUES ($1, $2, $3, $4) RETURNING user_id"

	// для читаемой записи времени удаления в базе данных проверяем задано ли время удаления в запросе.
	// если задано, то переводим время из Unix формата в формат String по заданному шаблону и вставлем в бд
	// если не задано, то вставляем nil он же пустое поле
	if user.TTlTime > 0 {
		*deleteTime = int64(now + user.TTlTime*3600)

		typeDuration := time.Unix(*deleteTime, 0).Format("2 January 2006 15:04")

		for _, j := range input {
			row := r.db.QueryRow(query, user.Id, j, time.Now(), typeDuration)
			if err := row.Scan(&id); err != nil {
				return "error in db in func AddUserToSlug", err
			}
		}
	} else {
		deleteTime = nil

		for _, j := range input {
			row := r.db.QueryRow(query, user.Id, j, time.Now(), deleteTime)
			if err := row.Scan(&id); err != nil {
				return "error in db in func AddUserToSlug", err
			}
		}
	}

	go r.TTL() // запуск автоматического удаления
	return "successes", nil
}

// DeleteUserFromSegments Функция удаления пользователя из одного или нескольких сегментов
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

// GetUserSegments Функция получения получения списка сегментов пользователя
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

// TTL функция работает в горутине и проверяет записи на удаляет, но почему-то не выводит ничего в консоль
// в бд хранится формат Time, что позваляет проводить данные сравнения
// цикл отправляет запрос в бд 1 раз в минуту, частоту можно менять при необходимости
func (r *AuthPostgres) TTL() {
	for range time.Tick(time.Minute) {
		_, err := r.db.Exec(
			`DELETE FROM segments_user WHERE time_delete <= $1`,
			time.Now())
		if err != nil {
			logrus.Errorf("error in ttl, error : %s", err)
		}
	}
	logrus.Info("success delete")
}
