package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-go.git/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
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

	if err := db.Delete(&opening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete opening"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Opening deleted successfully", "opening": opening})
}
