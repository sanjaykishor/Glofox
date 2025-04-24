package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanjaykishor/Glofox/internal/service"
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
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request data: " + err.Error(),
		})
		return
	}

	booking, err := h.bookingService.CreateBooking(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "Booking created successfully",
		Data:    booking,
	})
}

// GetAllBookings returns all bookings
func (h *BookingHandler) GetAllBookings(c *gin.Context) {
	bookings := h.bookingService.GetAllBookings()
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    bookings,
	})
}

// GetBookingByID retrieves a booking by its ID
func (h *BookingHandler) GetBookingByID(c *gin.Context) {
	id := c.Param("id")
	booking, err := h.bookingService.GetBookingByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    booking,
	})
}

// GetBookingsByDate retrieves all bookings for a specific date
func (h *BookingHandler) GetBookingsByDate(c *gin.Context) {
	date := c.Param("date")
	bookings, err := h.bookingService.GetBookingsByDate(date)

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    bookings,
	})
}
