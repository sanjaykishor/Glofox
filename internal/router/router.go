package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sanjaykishor/Glofox/internal/handler"
	"github.com/sanjaykishor/Glofox/internal/middleware"
)

func Setup(
	classHandler *handler.ClassHandler,
	bookingHandler *handler.BookingHandler,
) *gin.Engine {

	router := gin.Default()
	middleware.Setup(router)
	setupAPIRoutes(router, classHandler, bookingHandler)

	return router
}

// setupAPIRoutes configures all the API routes for the application
func setupAPIRoutes(
	router *gin.Engine,
	classHandler *handler.ClassHandler,
	bookingHandler *handler.BookingHandler,
) {
	api := router.Group("/api/v1")
	{
		// Register class routes
		classHandler.RegisterRoutes(api)

		// Register booking routes
		bookingHandler.RegisterRoutes(api)
	}

}
