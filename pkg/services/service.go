package services

import (
	"internship_avito/pkg/model"
	"internship_avito/pkg/repository"
)

type LogicsUser interface {
	CreateUser(user model.User) (int, error)
	AddUserToSegment(user model.UserSegments) (string, error)
	DeleteUserFromSegment(user model.UserSegments) (string, error)
	GetUserSegment(user model.User) ([]string, error)
}

type LogicSegment interface {
	CreateSegments(segments model.Segments) (string, error)
	DeleteSegments(segments model.Segments) (string, error)
	UserCountInSegment(segments model.Segments) (int, error)
}

type Service struct {
	LogicsUser
	LogicSegment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		LogicsUser:   NewAuthService(repos.LogicsUser),
		LogicSegment: NewSegmentService(repos.LogicSegments),
	}
}
