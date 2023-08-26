package services

import (
	"internship_avito/pkg/model"
	"internship_avito/pkg/repository"
)

type Logics interface {
	CreateUser(user model.User) (int, error)
}

type Service struct {
	Logics
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Logics: NewAuthService(repos.Logics),
	}
}
