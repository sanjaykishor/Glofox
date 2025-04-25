package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanjaykishor/Glofox/internal/service"
	"github.com/sanjaykishor/Glofox/internal/validation"
)

type BookingHandler struct {
	bookingService *service.BookingService
}

func NewBookingHandler(bookingService *service.BookingService) *BookingHandler {
	return &BookingHandler{
		bookingService: bookingService,
	}
}

func (h *BookingHandler) RegisterRoutes(router gin.IRouter) {
	bookingsGroup := router.Group("/bookings")
	{
		bookingsGroup.POST("", h.CreateBooking)
		bookingsGroup.GET("", h.GetAllBookings)
		bookingsGroup.GET("/:id", h.GetBookingByID)
		bookingsGroup.GET("/date/:date", h.GetBookingsByDate)
	}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var request service.CreateBookingRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		validation.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	booking, err := h.bookingService.CreateBooking(&request)
	if err != nil {
		validation.ServiceErrorResponse(c, err)
		return
	}

	validation.SuccessResponse(c, http.StatusCreated, "Booking created successfully", booking)
}

// GetAllBookings returns all bookings
func (h *BookingHandler) GetAllBookings(c *gin.Context) {
	bookings := h.bookingService.GetAllBookings()
	validation.SuccessResponse(c, http.StatusOK, "", bookings)
}

// GetBookingByID retrieves a booking by its ID
func (h *BookingHandler) GetBookingByID(c *gin.Context) {
	id := c.Param("id")
	booking, err := h.bookingService.GetBookingByID(id)
	if err != nil {
		validation.ServiceErrorResponse(c, err)
		return
	}

	validation.SuccessResponse(c, http.StatusOK, "", booking)
}

// GetBookingsByDate retrieves all bookings for a specific date
func (h *BookingHandler) GetBookingsByDate(c *gin.Context) {
	date := c.Param("date")
	bookings, err := h.bookingService.GetBookingsByDate(date)

	if err != nil {
		validation.ServiceErrorResponse(c, err)
		return
	}

	validation.SuccessResponse(c, http.StatusOK, "", bookings)
}
