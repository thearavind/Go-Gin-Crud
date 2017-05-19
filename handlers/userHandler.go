package handlers

import (
	"net/http"

	"github.com/aravind741/Go-Gin-Crud/models"
	"github.com/gin-gonic/gin"
)

// FetchUser - Gets the user ID from auth token and returns that user form the DB
func FetchUser(c *gin.Context) {
	_, hasToken := c.Get("token")
	if hasToken {
		userId, _ := c.Get("user_id")
		currentUser := models.Users{UserId: userId.(int)}
		_ = ORM.Read(&currentUser)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "user_id": currentUser.UserId,
			"user_name": currentUser.UserName, "email": currentUser.Email})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized,
			"message": "No auth token sent in the request"})
	}
}
