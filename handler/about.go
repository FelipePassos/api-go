package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAbout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "github.com/FelipePassos/api-go V1",
	})
}
