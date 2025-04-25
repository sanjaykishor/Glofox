package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanjaykishor/Glofox/internal/service"
	"github.com/sanjaykishor/Glofox/internal/validation"
)

type ClassHandler struct {
	classService *service.ClassService
}

func NewClassHandler(classService *service.ClassService) *ClassHandler {
	return &ClassHandler{
		classService: classService,
	}
}

func (h *ClassHandler) RegisterRoutes(router gin.IRouter) {
	classesGroup := router.Group("/classes")
	{
		classesGroup.POST("", h.CreateClass)
		classesGroup.GET("", h.GetAllClasses)
		classesGroup.GET("/:id", h.GetClassByID)
	}
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var request service.CreateClassRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		validation.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	class, err := h.classService.CreateClass(&request)
	if err != nil {
		validation.ServiceErrorResponse(c, err)
		return
	}

	validation.SuccessResponse(c, http.StatusCreated, "Class created successfully", class)
}

// GetAllClasses returns all classes
func (h *ClassHandler) GetAllClasses(c *gin.Context) {
	classes := h.classService.GetAllClasses()
	validation.SuccessResponse(c, http.StatusOK, "", classes)
}

// GetClassByID retrieves a class by its ID
func (h *ClassHandler) GetClassByID(c *gin.Context) {
	id := c.Param("id")
	class, err := h.classService.GetClassByID(id)
	if err != nil {
		validation.ServiceErrorResponse(c, err)
		return
	}

	validation.SuccessResponse(c, http.StatusOK, "", class)
}
