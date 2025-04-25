package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)


// Setup configures global middleware for the application
func Setup(router *gin.Engine) {
	router.Use(RequestLogger())
	router.Use(CORS())
}

// RequestLogger returns a middleware that logs request details
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		path := c.Request.URL.Path
		method := c.Request.Method
		statusCode := c.Writer.Status()

		logMessage := fmt.Sprintf("[GIN] %s | %s | %s | %d | %s\n",
			time.Now().Format("2006/01/02 - 15:04:05"),
			method,
			path,
			statusCode,
			latency.String())

		_, err := gin.DefaultWriter.Write([]byte(logMessage))
		if err != nil {
			log.Println("Failed to write request log:", err)
		}
	}
}

// CORS returns a middleware for handling Cross-Origin Resource Sharing
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // TODO: Change to specific origin
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

