package main


import (
	"github.com/aravind741/Go-Gin-Crud/models"
	"github.com/aravind741/Go-Gin-Crud/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // import your used driver
)
func main() {
	models.ConnectToDb()
	//o := orm.NewOrm()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.POST("api/users/login", handlers.LoginHandler)
	handlers.ORM = models.GetOrmObject()
/*	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", CreateTodo)
		v1.GET("/", FetchAllTodo)
		v1.GET("/:id", FetchSingleTodo)
		v1.PUT("/:id", UpdateTodo)
		v1.DELETE("/:id", DeleteTodo)
	}*/
	router.Run("localhost:1445")
/*	user := models.Users{User_id:1}
	_ = o.Read(&user)
	fmt.Println("User ID 1 is ->", user)*/
}
