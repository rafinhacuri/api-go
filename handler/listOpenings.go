package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-go.git/schemas"
)

type ListOpeningResponse struct {
	Openings []schemas.OpeningResponse `json:"openings"`
}

// @BasePath /api/v1

// @Summary List job openings
// @Description List job openings
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	var openings []schemas.Opening

	if err := db.Find(&openings).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve openings"})
		return
	}

	ctx.JSON(200, gin.H{"openings": openings})
}
