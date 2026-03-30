package routes

import (
	"crud_go/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	//grouping routes that require authentication

	authRoutes := server.Group("/")
	authRoutes.Use(middlewares.Authenticate)

	authRoutes.POST("/events", createEvent)
	authRoutes.PUT("/events/:id", updateEvent)
	authRoutes.DELETE("/events/:id", deletEvent)

	// server.POST("/events", middlewares.Authenticate, createEvent)
	// server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	// server.DELETE("/events/:id", middlewares.Authenticate, deletEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
