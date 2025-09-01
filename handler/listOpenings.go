package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-go.git/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	var openings []schemas.Opening

	if err := db.Find(&openings).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve openings"})
		return
	}

	ctx.JSON(200, gin.H{"openings": openings})
}
