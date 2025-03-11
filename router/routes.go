package router

import (
	"github.com/FelipePassos/api-go/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("api/v1")
	{
		v1.GET("/status", handler.ShowAbout)
	}
}
