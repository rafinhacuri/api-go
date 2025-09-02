package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-go.git/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, request.ID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Opening not found"})
		return
	}

	opening.Role = request.Role
	opening.Company = request.Company
	opening.Location = request.Location
	opening.Remote = *request.Remote
	opening.Link = request.Link
	opening.Salary = request.Salary

	if err := db.Save(&opening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update opening"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Opening updated successfully", "opening": opening})
}
