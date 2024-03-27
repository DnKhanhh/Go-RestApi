package routes

import (
	"example.com/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
