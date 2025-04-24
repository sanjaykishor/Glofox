package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sanjaykishor/Glofox/internal/handler"
	"github.com/sanjaykishor/Glofox/internal/repository"
	"github.com/sanjaykishor/Glofox/internal/router"
	"github.com/sanjaykishor/Glofox/internal/service"
)

func main() {
	// Initialize repositories
	classRepo := repository.NewClassRepository()
	bookingRepo := repository.NewBookingRepository()

	// Initialize services
	classService := service.NewClassService(classRepo)
	bookingService := service.NewBookingService(bookingRepo, classRepo)

	// Initialize handlers
	classHandler := handler.NewClassHandler(classService)
	bookingHandler := handler.NewBookingHandler(bookingService)

	// Initialize router
	r := router.Setup(classHandler, bookingHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Println("Server starting on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
