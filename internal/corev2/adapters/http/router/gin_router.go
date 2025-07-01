package router

import (
	httpAdapter "hexagonal-architecture/internal/corev2/adapters/controller/user"

	"github.com/gin-gonic/gin"
)

// GinRouter handles routing for the Gin framework
type GinRouter struct {
	userController *httpAdapter.Controller
}

// NewGinRouter creates a new GinRouter
func NewGinRouter(userController *httpAdapter.Controller) *GinRouter {
	return &GinRouter{
		userController: userController,
	}
}

// RegisterRoutes registers all routes to the Gin engine
func (r *GinRouter) RegisterRoutes(engine *gin.Engine) {
	// Register index route
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hexagonal Architecture Example - User Management API",
		})
	})

	// User routes
	userGroup := engine.Group("/users")
	{
		userGroup.POST("", r.userController.HandleCreateUser)
		userGroup.GET("/:id", r.userController.HandleGetUser)
		userGroup.DELETE("/:id", r.userController.HandleDeleteUser)
	}
}
