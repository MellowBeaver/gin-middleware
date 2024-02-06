package main

import (
	"fmt"

	"gin-middleware/api/handlers"
	"gin-middleware/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	//API using JWT (JSON Web Tokens) for authentication and a logging middleware.
	router := gin.Default()

	// Middleware
	router.Use(middleware.LoggerMiddleware)

	// Route without authentication
	router.POST("/sam-login", handlers.LoginHandler)

	// Route with authentication middleware
	authGroup := router.Group("/sam-api")
	authGroup.Use(middleware.AuthMiddleware)
	{
		authGroup.GET("/secured", handlers.SecuredHandler)
	}

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
