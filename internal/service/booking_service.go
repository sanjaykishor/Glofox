package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sanjaykishor/Glofox/internal/repository"
)

// BookingService handles business logic for bookings
type BookingService struct {
	bookingRepo *repository.BookingRepository
	classRepo   *repository.ClassRepository
}

// NewBookingService creates a new instance of BookingService
func NewBookingService(bookingRepo *repository.BookingRepository, classRepo *repository.ClassRepository) *BookingService {
	return &BookingService{
		bookingRepo: bookingRepo,
		classRepo:   classRepo,
	}
}

// CreateBookingRequest represents the data needed to create a booking
type CreateBookingRequest struct {
	MemberName string `json:"name" binding:"required"`
	Date       string `json:"date" binding:"required"`
	ClassID    string `json:"class_id"`
}

// CreateBooking creates a new booking
func (s *BookingService) CreateBooking(req *CreateBookingRequest) (*repository.Booking, error) {
	bookingDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("invalid date format, use YYYY-MM-DD")
	}

	if req.ClassID != "" {
		_, err := s.classRepo.GetByID(req.ClassID)
		if err != nil {
			return nil, errors.New("class not found")
		}
	}

	booking := &repository.Booking{
		ID:         uuid.New().String(),
		MemberName: req.MemberName,
		ClassID:    req.ClassID,
		Date:       bookingDate,
		CreatedAt:  time.Now(),
	}

	if err := s.bookingRepo.Create(booking); err != nil {
		return nil, err
	}

	return booking, nil
}

// GetAllBookings returns all bookings
func (s *BookingService) GetAllBookings() []*repository.Booking {
	return s.bookingRepo.GetAll()
}

// GetBookingByID retrieves a booking by its ID
func (s *BookingService) GetBookingByID(id string) (*repository.Booking, error) {
	return s.bookingRepo.GetByID(id)
}

// GetBookingsByDate retrieves all bookings for a specific date
func (s *BookingService) GetBookingsByDate(dateStr string) ([]*repository.Booking, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, errors.New("invalid date format, use YYYY-MM-DD")
	}

	return s.bookingRepo.GetBookingsByDate(date), nil
}
