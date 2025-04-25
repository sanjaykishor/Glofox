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
	"github.com/stretchr/testify/assert"
)

func setupClassTestRouter() (*gin.Engine, *repository.ClassRepository) {
	gin.SetMode(gin.TestMode)

	classRepo := repository.NewClassRepository()

	classService := service.NewClassService(classRepo)
	classHandler := NewClassHandler(classService)

	router := gin.New()
	classHandler.RegisterRoutes(router.Group("/api/v1"))

	return router, classRepo
}

func TestCreateClass(t *testing.T) {
	router, _ := setupClassTestRouter()

	startDate := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	endDate := time.Now().Add(48 * time.Hour).Format("2006-01-02")

	validRequest := map[string]interface{}{
		"name":       "Yoga Class",
		"start_date": startDate,
		"end_date":   endDate,
		"capacity":   20,
	}

	jsonData, _ := json.Marshal(validRequest)
	req, _ := http.NewRequest("POST", "/api/v1/classes", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Should return status code 201")

	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should parse response JSON without error")
	assert.True(t, response.Success, "Response success should be true")

	// Missing start_date, end_date, and capacity
	invalidRequest := map[string]any{
		"name": "Yoga Class",
	}

	jsonData, _ = json.Marshal(invalidRequest)
	req, _ = http.NewRequest("POST", "/api/v1/classes", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should return status code 400")

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should parse response JSON without error")
	assert.False(t, response.Success, "Response success should be false")

	// Invalid date format
	invalidDateRequest := map[string]any{
		"name":       "Yoga Class",
		"start_date": "invalid-date",
		"end_date":   endDate,
		"capacity":   20,
	}

	jsonData, _ = json.Marshal(invalidDateRequest)
	req, _ = http.NewRequest("POST", "/api/v1/classes", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Should return status code 400")
	assert.False(t, response.Success, "Response success should be false")
}

func TestGetAllClasses(t *testing.T) {
	router, classRepo := setupClassTestRouter()

	class := &repository.Class{
		ID:        "11111111-1111-1111-1111-111111111111",
		Name:      "Yoga",
		StartDate: time.Now().Add(24 * time.Hour),
		EndDate:   time.Now().Add(48 * time.Hour),
		Capacity:  20,
	}
	classRepo.Create(class)

	req, _ := http.NewRequest("GET", "/api/v1/classes", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Should return status code 200")

	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should parse response JSON without error")
	assert.True(t, response.Success, "Response success should be true")

	// Check that we got classes in the response
	classesData, ok := response.Data.([]any)
	assert.True(t, ok, "Data should be an array of classes")
	assert.NotEmpty(t, classesData, "Should return at least one class")
}

func TestGetClassByID(t *testing.T) {
	router, classRepo := setupClassTestRouter()

	class := &repository.Class{
		ID:        "11111111-1111-1111-1111-111111111111",
		Name:      "Yoga",
		StartDate: time.Now().Add(24 * time.Hour),
		EndDate:   time.Now().Add(48 * time.Hour),
		Capacity:  20,
	}
	classRepo.Create(class)

	req, _ := http.NewRequest("GET", "/api/v1/classes/11111111-1111-1111-1111-111111111111", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Should return status code 200")

	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should parse response JSON without error")
	assert.True(t, response.Success, "Response success should be true")

	// Non-existent class
	req, _ = http.NewRequest("GET", "/api/v1/classes/non-existent-id", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Should return status code 404")

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Should parse response JSON without error")
	assert.False(t, response.Success, "Response success should be false")
}
