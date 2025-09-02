package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-go.git/schemas"
)

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Opening schemas.OpeningResponse `json:"opening"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

// @BasePath /api/v1

// @Summary Create a new job opening
// @Description Create a new job opening
// @Tags opening
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request Body"
// @Success 201 {object} CreateOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Failed to create opening: %v", err)
		ctx.JSON(500, gin.H{"error": "Failed to create opening"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Opening created successfully", "opening": opening})
}
