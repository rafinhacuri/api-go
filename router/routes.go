package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-go.git/docs"
	"github.com/rafinhacuri/api-go.git/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	handler.InitHandler()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	v1.GET("/openings", handler.ListOpeningsHandler)
	v1.GET("/opening", handler.GetOpeningHandler)
	v1.POST("/opening", handler.CreateOpeningHandler)
	v1.DELETE("/opening", handler.DeleteOpeningHandler)
	v1.PUT("/opening", handler.UpdateOpeningHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
