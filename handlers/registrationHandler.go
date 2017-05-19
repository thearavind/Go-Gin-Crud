package handlers

import (
	"fmt"
	"net/http"

	"github.com/aravind741/Go-Gin-Crud/models"
	"github.com/gin-gonic/gin"
)

// RegistrationRequest - Registration request parameter format
type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

// RegistrationHandler - Encrypts the new users password and stores user details in the DB
func RegistrationHandler(c *gin.Context) {
	var regRequest RegistrationRequest
	c.BindJSON(&regRequest)
	if validateNewUser(regRequest.Email, c) {
		encryptedPassword, err := encrypt([]byte(regRequest.Password))
		if err != nil {
			fmt.Println("Failed to encrypt the password:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError,
				"message": "Failed encrypt the password"})
		} else {
			newUser := models.Users{
				Email:    regRequest.Email,
				UserName: regRequest.UserName,
				Password: encryptedPassword,
			}
			_, err = ORM.Insert(&newUser)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "email": newUser.Email,
					"user_name": newUser.UserName, "user_id": newUser.UserId})
			}
		}
	}
}

// validateNewUser - Checks weather the users email_id is already present in the database
func validateNewUser(userEmail string, c *gin.Context) (s bool) {
	var queryUsers []models.Users
	ORM.QueryTable("users").Filter("email__exact", userEmail).All(&queryUsers)
	if len(queryUsers) > 0 {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict,
			"message": "User has been registered already"})
		return false
	}
	return true
}
