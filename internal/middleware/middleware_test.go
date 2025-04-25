package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestLogger(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var buf bytes.Buffer
	gin.DefaultWriter = &buf

	router := gin.New()
	router.Use(RequestLogger())

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(recorder, req)

	logOutput := buf.String()
	assert.Contains(t, logOutput, "[GIN]")
	assert.Contains(t, logOutput, "GET")
	assert.Contains(t, logOutput, "/test")
	assert.Contains(t, logOutput, "200")

	elapsedTimePattern := `\d+(\.\d+)?(µs|ms|s|m|h)`
	assert.Regexp(t, elapsedTimePattern, logOutput)
}

func TestCORS(t *testing.T) {
	t.Run("Regular Request", func(t *testing.T) {
		router := gin.New()
		router.Use(CORS())

		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "test")
		})

		recorder := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, "*", recorder.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "true", recorder.Header().Get("Access-Control-Allow-Credentials"))
		assert.NotEmpty(t, recorder.Header().Get("Access-Control-Allow-Headers"))
		assert.NotEmpty(t, recorder.Header().Get("Access-Control-Allow-Methods"))

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "test", recorder.Body.String())
	})

	t.Run("OPTIONS Request", func(t *testing.T) {
		router := gin.New()
		router.Use(CORS())

		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "test")
		})

		recorder := httptest.NewRecorder()
		req, _ := http.NewRequest("OPTIONS", "/test", nil)
		router.ServeHTTP(recorder, req)

		assert.Equal(t, "*", recorder.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "true", recorder.Header().Get("Access-Control-Allow-Credentials"))
		assert.NotEmpty(t, recorder.Header().Get("Access-Control-Allow-Headers"))
		assert.NotEmpty(t, recorder.Header().Get("Access-Control-Allow-Methods"))

		assert.Equal(t, http.StatusNoContent, recorder.Code) // 204 No Content
		assert.Empty(t, recorder.Body.String())
	})
}

