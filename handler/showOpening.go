package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-go.git/schemas"
)

type ShowOpeningResponse struct {
	Opening schemas.OpeningResponse `json:"opening"`
}

// @BasePath /api/v1

// @Summary Show a job opening
// @Description Show a job opening
// @Tags opening
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} ShowOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Opening not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"opening": opening})
}
