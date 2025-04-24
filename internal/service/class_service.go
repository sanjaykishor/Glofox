package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sanjaykishor/Glofox/internal/repository"
)

type ClassService struct {
	repo *repository.ClassRepository
}

func NewClassService(repo *repository.ClassRepository) *ClassService {
	return &ClassService{
		repo: repo,
	}
}

// CreateClassRequest represents the data needed to create a class
type CreateClassRequest struct {
	Name      string `json:"name" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	Capacity  int    `json:"capacity" binding:"required,min=1"`
}

func (s *ClassService) CreateClass(req *CreateClassRequest) (*repository.Class, error) {
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, errors.New("invalid start date format, use YYYY-MM-DD")
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, errors.New("invalid end date format, use YYYY-MM-DD")
	}

	if endDate.Before(startDate) {
		return nil, errors.New("end date cannot be before start date")
	}

	class := &repository.Class{
		ID:        uuid.New().String(),
		Name:      req.Name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  req.Capacity,
	}

	if err := s.repo.Create(class); err != nil {
		return nil, err
	}

	return class, nil
}

// GetAllClasses returns all classes
func (s *ClassService) GetAllClasses() []*repository.Class {
	return s.repo.GetAll()
}

// GetClassByID retrieves a class by its ID
func (s *ClassService) GetClassByID(id string) (*repository.Class, error) {
	return s.repo.GetByID(id)
}
