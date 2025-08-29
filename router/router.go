package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter(port string) {
	router := gin.Default()

	InitializeRoutes(router)

	router.Run(port)
}
