package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-go.git/schemas"
)

type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Opening schemas.OpeningResponse `json:"opening"`
}

// @BasePath /api/v1

// @Summary Update a job opening
// @Description Update a job opening
// @Tags opening
// @Accept json
// @Produce json
// @Param request body UpdateOpeningRequest true "Request Body"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
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
