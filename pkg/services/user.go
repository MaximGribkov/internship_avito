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

func (s *AuthService) AddUserToSegment(user model.UserSegments) (string, error) {
	return s.repo.AddUserToSegments(user)
}

func (s *AuthService) DeleteUserFromSegment(user model.UserSegments) (string, error) {
	return s.repo.DeleteUserFromSegments(user)
}

func (s *AuthService) GetUserSegment(user model.User) ([]string, error) {
	return s.repo.GetUserSegments(user)
}
