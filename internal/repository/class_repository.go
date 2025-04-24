package repository

import (
	"errors"
	"sync"
	"time"
)

// Class represents a fitness class
type Class struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Capacity  int       `json:"capacity"`
}

// ClassRepository handles class data storage
type ClassRepository struct {
	classes map[string]*Class
	mutex   sync.RWMutex
}

// NewClassRepository creates a new instance of ClassRepository
func NewClassRepository() *ClassRepository {
	return &ClassRepository{
		classes: make(map[string]*Class),
	}
}

// Create adds a new class to the repository
func (r *ClassRepository) Create(class *Class) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.classes[class.ID]; exists {
		return errors.New("class with this ID already exists")
	}

	r.classes[class.ID] = class
	return nil
}

// GetAll returns all classes
func (r *ClassRepository) GetAll() []*Class {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	classes := make([]*Class, 0, len(r.classes))
	for _, class := range r.classes {
		classes = append(classes, class)
	}
	return classes
}

// GetByID retrieves a class by its ID
func (r *ClassRepository) GetByID(id string) (*Class, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	class, exists := r.classes[id]
	if !exists {
		return nil, errors.New("class not found")
	}

	return class, nil
}
