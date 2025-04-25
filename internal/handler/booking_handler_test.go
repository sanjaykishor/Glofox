package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanjaykishor/Glofox/internal/repository"
	"github.com/sanjaykishor/Glofox/internal/service"
	"github.com/sanjaykishor/Glofox/internal/validation"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, *repository.BookingRepository, *repository.ClassRepository) {
	gin.SetMode(gin.TestMode)

	bookingRepo := repository.NewBookingRepository()
	classRepo := repository.NewClassRepository()

	class := &repository.Class{
		ID:        "test-class-1",
		Name:      "Yoga",
		StartDate: time.Now().Add(24 * time.Hour), // Future date
		EndDate:   time.Now().Add(48 * time.Hour),
		Capacity:  20,
	}
	classRepo.Create(class)

	bookingService := service.NewBookingService(bookingRepo, classRepo)
	bookingHandler := NewBookingHandler(bookingService)

	router := gin.New()
	bookingHandler.RegisterRoutes(router.Group("/api/v1"))

	return router, bookingRepo, classRepo
}

func TestCreateBooking(t *testing.T) {
	router, _, _ := setupTestRouter()

	validRequest := map[string]interface{}{
		"name":     "John Doe",
		"date":     time.Now().Format("2006-01-02"),
		"class_id": "test-class-1",
	}

	jsonData, _ := json.Marshal(validRequest)
	req, _ := http.NewRequest("POST", "/api/v1/bookings", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Should return status code 201")

	var response validation.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should parse response JSON without error")
	assert.True(t, response.Success, "Response success should be true")

	invalidRequest := map[string]interface{}{
		"name": "John Doe",
	}

	jsonData, _ = json.Marshal(invalidRequest)
	req, _ = http.NewRequest("POST", "/api/v1/bookings", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should return status code 400")

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should parse response JSON without error")
	assert.False(t, response.Success, "Response success should be false")
	assert.Contains(t, response.Error, "date is required", "Error message should indicate missing date field")
}

func TestGetAllBookings(t *testing.T) {
	router, bookingRepo, _ := setupTestRouter()

	booking := &repository.Booking{
		ID:         "test-booking-1",
		MemberName: "John Doe",
		ClassID:    "test-class-1",
		Date:       time.Now(),
		CreatedAt:  time.Now(),
	}
	bookingRepo.Create(booking)

	req, _ := http.NewRequest("GET", "/api/v1/bookings", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Should return status code 200")

	var response validation.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should parse response JSON without error")
	assert.True(t, response.Success, "Response success should be true")

	bookingsData, ok := response.Data.([]interface{})
	assert.True(t, ok, "Data should be an array of bookings")
	assert.NotEmpty(t, bookingsData, "Should return at least one booking")
}

func TestGetBookingByID(t *testing.T) {
	router, bookingRepo, _ := setupTestRouter()

	booking := &repository.Booking{
		ID:         "test-booking-1",
		MemberName: "John Doe",
		ClassID:    "test-class-1",
		Date:       time.Now(),
		CreatedAt:  time.Now(),
	}
	bookingRepo.Create(booking)

	req, _ := http.NewRequest("GET", "/api/v1/bookings/test-booking-1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Should return status code 200")

	var response validation.Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should parse response JSON without error")
	assert.True(t, response.Success, "Response success should be true")

	req, _ = http.NewRequest("GET", "/api/v1/bookings/non-existent", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Should return status code 404")

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should parse response JSON without error")
	assert.False(t, response.Success, "Response success should be false")
}
