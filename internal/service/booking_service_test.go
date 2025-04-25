package service

import (
	"testing"
	"time"

	"github.com/sanjaykishor/Glofox/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestBookingService(t *testing.T) {
	bookingRepo := repository.NewBookingRepository()
	classRepo := repository.NewClassRepository()

	class := &repository.Class{
		ID:        "test-class-1",
		Name:      "Yoga",
		StartDate: time.Now().Add(24 * time.Hour), // Future date
		EndDate:   time.Now().Add(25 * time.Hour),
		Capacity:  20,
	}

	err := classRepo.Create(class)
	assert.NoError(t, err, "Should create test class without error")

	service := NewBookingService(bookingRepo, classRepo)

	createReq := &CreateBookingRequest{
		MemberName: "John Doe",
		Date:       time.Now().Format("2006-01-02"),
		ClassID:    "test-class-1",
	}

	booking, err := service.CreateBooking(createReq)
	assert.NoError(t, err, "Should create booking without error")
	assert.Equal(t, createReq.MemberName, booking.MemberName, "Booking member name should match request")
	assert.Equal(t, createReq.ClassID, booking.ClassID, "Booking class ID should match request")

	// Test getting all bookings
	allBookings := service.GetAllBookings()
	assert.Len(t, allBookings, 1, "Should return 1 booking")

	// Test getting booking by ID
	retrievedBooking, err := service.GetBookingByID(booking.ID)
	assert.NoError(t, err, "Should retrieve booking by ID without error")
	assert.Equal(t, booking.ID, retrievedBooking.ID, "Retrieved booking ID should match")

	// Test getting bookings by date
	bookingsByDate, err := service.GetBookingsByDate(time.Now().Format("2006-01-02"))
	assert.NoError(t, err, "Should retrieve bookings by date without error")
	assert.Len(t, bookingsByDate, 1, "Should return 1 booking for date")

	// Test error cases

	// Invalid date format
	_, err = service.CreateBooking(&CreateBookingRequest{
		MemberName: "USER A",
		Date:       "invalid-date",
		ClassID:    "test-class-1",
	})
	assert.Error(t, err, "Should return error for invalid date format")

	// Non-existent class
	_, err = service.CreateBooking(&CreateBookingRequest{
		MemberName: "USER A",
		Date:       time.Now().Format("2006-01-02"),
		ClassID:    "non-existent-class",
	})
	assert.Error(t, err, "Should return error for non-existent class")

	// Non-existent booking
	_, err = service.GetBookingByID("non-existent-booking")
	assert.Error(t, err, "Should return error for non-existent booking")

	// Invalid date format for GetBookingsByDate
	_, err = service.GetBookingsByDate("invalid-date")
	assert.Error(t, err, "Should return error for invalid date format in GetBookingsByDate")
}
