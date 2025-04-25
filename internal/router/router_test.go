package router

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sanjaykishor/Glofox/internal/handler"
	"github.com/sanjaykishor/Glofox/internal/repository"
	"github.com/sanjaykishor/Glofox/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestRouterSetup(t *testing.T) {
	classRepo := repository.NewClassRepository()
	bookingRepo := repository.NewBookingRepository()

	classService := service.NewClassService(classRepo)
	bookingService := service.NewBookingService(bookingRepo, classRepo)

	classHandler := handler.NewClassHandler(classService)
	bookingHandler := handler.NewBookingHandler(bookingService)

	gin.SetMode(gin.TestMode)

	router := Setup(classHandler, bookingHandler)

	assert.NotNil(t, router, "Router should not be nil")

	routes := router.Routes()
	assert.NotEmpty(t, routes, "Router should have routes registered")
}

func TestSetupAPIRoutes(t *testing.T) {
	classRepo := repository.NewClassRepository()
	bookingRepo := repository.NewBookingRepository()

	classService := service.NewClassService(classRepo)
	bookingService := service.NewBookingService(bookingRepo, classRepo)

	classHandler := handler.NewClassHandler(classService)
	bookingHandler := handler.NewBookingHandler(bookingService)

	gin.SetMode(gin.TestMode)

	router := gin.New()

	setupAPIRoutes(router, classHandler, bookingHandler)

	routes := router.Routes()
	assert.NotEmpty(t, routes, "Router should have routes registered")

	apiRouteFound := false
	for _, route := range routes {
		if route.Path == "/api/v1/classes" || route.Path == "/api/v1/bookings" {
			apiRouteFound = true
			break
		}
	}

	assert.True(t, apiRouteFound, "Router should have API routes registered")
}
