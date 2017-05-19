package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/aravind741/Go-Gin-Crud/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)
 var ORM orm.Ormer
func LoginHandler(c *gin.Context) {
	user := models.Users{User_id:1}
	_ = ORM.Read(&user)
	fmt.Println("Login route has been hit", user)
}

