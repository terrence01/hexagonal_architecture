package main

import (
	"context"
	httpAdapter "hexagonal-architecture/internal/core/adapters/in/controller/user"
	"hexagonal-architecture/internal/core/adapters/in/http/router"
	"hexagonal-architecture/internal/core/adapters/out/email"
	"hexagonal-architecture/internal/core/adapters/out/repository/memory"
	"hexagonal-architecture/internal/core/application/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create repository (output adapter)
	userRepository := memory.NewUserRepository()

	// Create email sender (output adapter)
	emailSender := email.NewEmailSender()

	// Create services (use case implementations)
	createUserService := user.NewCreateUserService(userRepository, emailSender)
	deleteUserService := user.NewDeleteUserService(userRepository, emailSender)
	getUserService := user.NewGetUserService(userRepository, emailSender)

	// Setup Gin router
	ginEngine := gin.Default()

	// Create HTTP controller (input adapter)
	userController := httpAdapter.NewUserController(
		createUserService,
		deleteUserService,
		getUserService,
	)

	// Create and register routes
	ginRouter := router.NewGinRouter(userController)
	ginRouter.RegisterRoutes(ginEngine)

	// Configure server
	server := &http.Server{
		Addr:    ":8080",
		Handler: ginEngine,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server started on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server gracefully stopped")
}
