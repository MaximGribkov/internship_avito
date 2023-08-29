package services

import (
	"internship_avito/pkg/model"
	"internship_avito/pkg/repository"
)

type SegmentService struct {
	repo repository.LogicSegments
}

func NewSegmentService(repo repository.LogicSegments) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegments(segments model.Segments) (string, error) {
	return s.repo.CreateSegments(segments)
}

func (s *SegmentService) DeleteSegments(segments model.Segments) (string, error) {
	return s.repo.DeleteSegments(segments)
}

func (s *SegmentService) UserCountInSegment(segments model.Segments) (int, error) {
	return s.repo.UserCountInSegment(segments)
}
