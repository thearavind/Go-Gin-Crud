package handlers

import (
	"fmt"
	"net/http"

	"github.com/aravind741/Go-Gin-Crud/models"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

// ORM - Global ORM object for the handler package
var ORM orm.Ormer

// LoginRequest - Login JSON request format
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler - Handler function for the login route
func LoginHandler(c *gin.Context) {
	var user []models.Users
	var loginRequest LoginRequest
	c.BindJSON(&loginRequest)
	ORM.QueryTable("users").Filter("email__exact", loginRequest.Email).All(&user)
	if len(user) != 0 {
		if decryptedPass, _ := decrypt(user[0].Password); decryptedPass == loginRequest.Password {
			if token, err := generateJWT(user[0].UserId); err != nil {
				fmt.Println("Failed to generate the JWT Token:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError,
					"message": "Failed to generate JWT Token"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "email": user[0].Email,
					"user_name": user[0].UserName, "token": token})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized,
				"message": "password is not valid"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized,
			"message": "User does not exist"})
	}

}

/*TODO build a JSON response generator*/
