package services

import "internship_avito/pkg/repository"

type Logics interface {
	//CreateUser
}

type Service struct {
	Logics
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
