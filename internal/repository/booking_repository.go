package repository

import (
	"errors"
	"sync"
	"time"
)

// Booking represents a class booking by a studio member
type Booking struct {
	ID         string    `json:"id"`
	MemberName string    `json:"member_name"`
	ClassID    string    `json:"class_id,omitempty"`
	Date       time.Time `json:"date"`
	CreatedAt  time.Time `json:"created_at"`
}

// BookingRepository handles booking data storage
type BookingRepository struct {
	bookings map[string]*Booking
	mutex    sync.RWMutex
}

// NewBookingRepository creates a new instance of BookingRepository
func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		bookings: make(map[string]*Booking),
	}
}

// Create adds a new booking to the repository
func (r *BookingRepository) Create(booking *Booking) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.bookings[booking.ID]; exists {
		return errors.New("booking with this ID already exists")
	}

	r.bookings[booking.ID] = booking
	return nil
}

// GetAll returns all bookings
func (r *BookingRepository) GetAll() []*Booking {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	bookings := make([]*Booking, 0, len(r.bookings))
	for _, booking := range r.bookings {
		bookings = append(bookings, booking)
	}
	return bookings
}

// GetByID retrieves a booking by its ID
func (r *BookingRepository) GetByID(id string) (*Booking, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	booking, exists := r.bookings[id]
	if !exists {
		return nil, errors.New("booking not found")
	}

	return booking, nil
}

// GetBookingsByDate retrieves all bookings for a specific date
func (r *BookingRepository) GetBookingsByDate(date time.Time) []*Booking {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// Format date to compare only year, month, day
	targetDate := date.Format("2006-01-02")

	bookings := make([]*Booking, 0)
	for _, booking := range r.bookings {
		if booking.Date.Format("2006-01-02") == targetDate {
			bookings = append(bookings, booking)
		}
	}

	return bookings
}
