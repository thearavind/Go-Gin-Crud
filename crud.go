package main

import (
	"github.com/aravind741/Go-Gin-Crud/router"
	_ "github.com/lib/pq"
)

func main() {
	/*Fetch the instance of the gin framework and run the server */
	router.GetMainEngine().Run(":1445")
}
