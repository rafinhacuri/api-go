package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("/opening", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Openiiiiing"}) })
	v1.POST("/opening", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Openiiiiing Post"}) })
	v1.DELETE("/opening", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Openiiiiing Delete"}) })
	v1.PUT("/opening", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Openiiiiing Put"}) })
}
