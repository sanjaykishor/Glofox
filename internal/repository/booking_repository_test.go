package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBookingRepository(t *testing.T) {
	repo := NewBookingRepository()

	booking := &Booking{
		ID:         "test-booking-1",
		MemberName: "John Doe",
		ClassID:    "test-class-1",
		Date:       time.Now(),
		CreatedAt:  time.Now(),
	}

	err := repo.Create(booking)
	assert.NoError(t, err, "Should create booking without error")

	// Test GetByID
	retrieved, err := repo.GetByID("test-booking-1")
	assert.NoError(t, err, "Should retrieve booking without error")
	assert.Equal(t, booking.ID, retrieved.ID, "Retrieved booking ID should match")
	assert.Equal(t, booking.MemberName, retrieved.MemberName, "Retrieved booking member name should match")

	// Test GetByID for non-existent booking
	_, err = repo.GetByID("non-existent-id")
	assert.Error(t, err, "Should return error for non-existent booking")

	// Test GetAll
	allBookings := repo.GetAll()
	assert.Len(t, allBookings, 1, "Should return 1 booking")

	// Test GetBookingsByDate
	bookingsByDate := repo.GetBookingsByDate(booking.Date)
	assert.Len(t, bookingsByDate, 1, "Should return 1 booking for date")

	// Test GetByClassID
	bookingsByClass, err := repo.GetByClassID(booking.ClassID)
	assert.NoError(t, err, "Should retrieve bookings by class ID without error")
	assert.Len(t, bookingsByClass, 1, "Should return 1 booking for class ID")

	// Test duplicate ID
	duplicateBooking := &Booking{
		ID:         "test-booking-1", // Same ID
		MemberName: "USER A",
		ClassID:    "test-class-2",
		Date:       time.Now(),
		CreatedAt:  time.Now(),
	}
	err = repo.Create(duplicateBooking)
	assert.Error(t, err, "Should return error for duplicate booking ID")
}
