package service

import (
	"testing"
	"time"

	"github.com/sanjaykishor/Glofox/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestClassService(t *testing.T) {
	repo := repository.NewClassRepository()

	service := NewClassService(repo)

	createReq := &CreateClassRequest{
		Name:      "Yoga",
		StartDate: time.Now().Add(24 * time.Hour).Format("2006-01-02"), // Future date
		EndDate:   time.Now().Add(48 * time.Hour).Format("2006-01-02"), // Later date
		Capacity:  20,
	}

	class, err := service.CreateClass(createReq)
	assert.NoError(t, err, "Should create class without error")
	assert.Equal(t, createReq.Name, class.Name, "Class name should match request")
	assert.Equal(t, createReq.Capacity, class.Capacity, "Class capacity should match request")

	allClasses := service.GetAllClasses()
	assert.Len(t, allClasses, 1, "Should return 1 class")

	retrievedClass, err := service.GetClassByID(class.ID)
	assert.NoError(t, err, "Should retrieve class by ID without error")
	assert.Equal(t, class.ID, retrievedClass.ID, "Retrieved class ID should match")

	// Test error cases

	// Invalid start date format
	_, err = service.CreateClass(&CreateClassRequest{
		Name:      "Gym",
		StartDate: "invalid-date",
		EndDate:   time.Now().Add(48 * time.Hour).Format("2006-01-02"),
		Capacity:  15,
	})
	assert.Error(t, err, "Should return error for invalid start date format")

	// Invalid end date format
	_, err = service.CreateClass(&CreateClassRequest{
		Name:      "Gym",
		StartDate: time.Now().Add(24 * time.Hour).Format("2006-01-02"),
		EndDate:   "invalid-date",
		Capacity:  15,
	})
	assert.Error(t, err, "Should return error for invalid end date format")

	// End date before start date
	_, err = service.CreateClass(&CreateClassRequest{
		Name:      "Gym",
		StartDate: time.Now().Add(48 * time.Hour).Format("2006-01-02"), // Later date
		EndDate:   time.Now().Add(24 * time.Hour).Format("2006-01-02"), // Earlier date
		Capacity:  15,
	})
	assert.Error(t, err, "Should return error for end date before start date")

	// Non-existent class
	_, err = service.GetClassByID("non-existent-class")
	assert.Error(t, err, "Should return error for non-existent class")
}
