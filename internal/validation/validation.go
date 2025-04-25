package validation

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Response is the standard API response structure
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// ValidateRequest handles the validation of the request from the client
func ValidateRequest(err error) (string, bool) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var fieldErrors []string
		for _, e := range ve {
			fieldName := e.Field()
			fieldName = mapFieldNameToJSON(fieldName)

			message := getValidationErrorMessage(e.Tag(), fieldName)
			fieldErrors = append(fieldErrors, message)
		}

		if len(fieldErrors) > 0 {
			return strings.Join(fieldErrors, ", "), true
		}
	}

	return "", false
}

// mapFieldNameToJSON converts struct field names to JSON field names
func mapFieldNameToJSON(fieldName string) string {
	switch fieldName {
	case "MemberName":
		return "name"
	case "ClassID":
		return "class_id"
	case "StartDate":
		return "start_date"
	case "EndDate":
		return "end_date"
	default:
		return strings.ToLower(fieldName)
	}
}

// getValidationErrorMessage returns an appropriate error message based on the validation tag
func getValidationErrorMessage(tag string, fieldName string) string {
	switch tag {
	case "required":
		return fieldName + " is required"
	case "min":
		return fieldName + " is below minimum value"
	case "max":
		return fieldName + " is above maximum value"
	case "email":
		return fieldName + " is not a valid email address"
	default:
		return fieldName + " is invalid: " + tag
	}
}

// ErrorResponse sends a standardized error response
func ErrorResponse(c *gin.Context, statusCode int, err error) {
	errorMessage := err.Error()

	if customMsg, isValidationErr := ValidateRequest(err); isValidationErr {
		errorMessage = customMsg
	}

	c.JSON(statusCode, Response{
		Success: false,
		Error:   errorMessage,
	})
}

// ServiceErrorResponse handles service related errors with appropriate status codes
func ServiceErrorResponse(c *gin.Context, err error) {
	statusCode := determineStatusCode(err)

	c.JSON(statusCode, Response{
		Success: false,
		Error:   err.Error(),
	})
}

// determineStatusCode analyzes the error message and returns an appropriate HTTP status code
func determineStatusCode(err error) int {
	errMsg := strings.ToLower(err.Error())

	if strings.Contains(errMsg, "not found") {
		return http.StatusNotFound
	} else if strings.Contains(errMsg, "unauthorized") {
		return http.StatusUnauthorized
	} else if strings.Contains(errMsg, "forbidden") {
		return http.StatusForbidden
	} else if strings.Contains(errMsg, "conflict") || strings.Contains(errMsg, "already exists") {
		return http.StatusConflict
	}

	return http.StatusBadRequest
}

// SuccessResponse sends a standardized success response
func SuccessResponse(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}
