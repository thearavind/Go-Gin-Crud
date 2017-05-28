package provider

import (
	"github.com/aravind741/Go-Gin-Crud/router"
	"github.com/gin-gonic/gin"
)

// ProvideRouter - Provides the gin instance for the test files
func ProvideRouter() *gin.Engine {
	return router.GetMainEngine()
}
