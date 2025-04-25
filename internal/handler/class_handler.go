package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanjaykishor/Glofox/internal/service"
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

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var request service.CreateClassRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request data: " + err.Error(),
		})
		return
	}

	class, err := h.classService.CreateClass(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "Class created successfully",
		Data:    class,
	})
}

// GetAllClasses returns all classes
func (h *ClassHandler) GetAllClasses(c *gin.Context) {
	classes := h.classService.GetAllClasses()
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    classes,
	})
}

// GetClassByID retrieves a class by its ID
func (h *ClassHandler) GetClassByID(c *gin.Context) {
	id := c.Param("id")
	class, err := h.classService.GetClassByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    class,
	})
}
