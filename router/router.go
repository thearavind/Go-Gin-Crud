package router

import (
	"github.com/gin-gonic/gin"
	"../models"
	"../handlers"
)

/* GetMainEngine - Creates the gin instance and returns it so that it can be used for dev and testing */
func GetMainEngine() *gin.Engine {
	models.ConnectToDb()
	httpRouter := gin.New()
	httpRouter.Use(gin.Logger())
	httpRouter.Use(gin.Recovery())
	handlers.ORM = models.GetOrmObject()
	httpRouter.POST("/api/users/login", handlers.LoginHandler)
	httpRouter.POST("/api/users", handlers.RegistrationHandler)
	RequireToken := httpRouter.Group("/")
	RequireToken.Use(handlers.TokenValidator())
	{
		RequireToken.GET("/api/user", handlers.FetchUser)
	}
	return httpRouter
}
