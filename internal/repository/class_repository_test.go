package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClassRepository(t *testing.T) {
	repo := NewClassRepository()

	// Test Create
	class := &Class{
		ID:        "test-class-1",
		Name:      "Yoga",
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Hour),
		Capacity:  20,
	}

	err := repo.Create(class)
	assert.NoError(t, err, "Should create class without error")

	// Test GetByID
	retrieved, err := repo.GetByID("test-class-1")
	assert.NoError(t, err, "Should retrieve class without error")
	assert.Equal(t, class.ID, retrieved.ID, "Retrieved class ID should match")
	assert.Equal(t, class.Name, retrieved.Name, "Retrieved class name should match")
	assert.Equal(t, class.Capacity, retrieved.Capacity, "Retrieved class capacity should match")

	// Test GetByID for non-existent class
	_, err = repo.GetByID("non-existent-id")
	assert.Error(t, err, "Should return error for non-existent class")

	// Test GetAll
	allClasses := repo.GetAll()
	assert.Len(t, allClasses, 1, "Should return 1 class")

	// Test duplicate ID
	duplicateClass := &Class{
		ID:        "test-class-1", // Same ID
		Name:      "Pilates",
		StartDate: time.Now().Add(time.Hour * 2),
		EndDate:   time.Now().Add(time.Hour * 3),
		Capacity:  15,
	}
	err = repo.Create(duplicateClass)
	assert.Error(t, err, "Should return error for duplicate class ID")
}
