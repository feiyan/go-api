package route

import (
	"github.com/gin-gonic/gin"
	region "gogin/app/controller"
	"gogin/app/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerToFile())
	r.GET("/region/detail/:id", region.Detail)
	r.GET("/region/sub", region.Sub)
	return r
}
