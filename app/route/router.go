package route

import (
	"github.com/gin-gonic/gin"
	region "gogin/app/controller"
	"gogin/app/middleware"
)

func SetupRouter() *gin.Engine {
	// init gin engine
	r := gin.Default()

	// middleware logs
	r.Use(middleware.LoggerToFile())

	// route group region
	regionGroup := r.Group("/region")
	regionGroup.POST("/add", region.Add)
	regionGroup.GET("/detail", region.Detail)
	regionGroup.GET("/sub", region.Sub)
	regionGroup.GET("/tree", region.Tree)
	return r
}
