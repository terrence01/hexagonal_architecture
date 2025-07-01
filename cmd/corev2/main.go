package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"hexagonal-architecture/internal/corev2/adapters/controller/user"
	"hexagonal-architecture/internal/corev2/adapters/email"
	"hexagonal-architecture/internal/corev2/adapters/http/router"
	"hexagonal-architecture/internal/corev2/adapters/repository/memory"
	"hexagonal-architecture/internal/corev2/application/user/createuserservice"
	"hexagonal-architecture/internal/corev2/application/user/deleteuserservice"
	"hexagonal-architecture/internal/corev2/application/user/getuserservice"
)

func main() {
	// Create repository (output adapter)
	userRepository := memory.NewUserRepository()

	// Create email sender (output adapter)
	emailSender := email.NewEmailSender()

	// Create services (use case implementations)
	createUserService := createuserservice.NewService(userRepository, emailSender)
	deleteUserService := deleteuserservice.NewService(userRepository, emailSender)
	getUserService := getuserservice.NewService(userRepository)

	// Setup Gin router
	ginEngine := gin.Default()

	// Create HTTP controller (input adapter)
	userController := user.NewController(
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
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
