package services

import (
	"internship_avito/pkg/model"
	"internship_avito/pkg/repository"
)

type AuthService struct {
	repo repository.LogicsUser
}

func NewAuthService(repo repository.LogicsUser) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	return s.repo.CreateUser(user)
}
