package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelpalhano/categories-ms/cmd/internal/repositories"
	use_cases "github.com/raphaelpalhano/categories-ms/cmd/internal/use-cases"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func badRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(
		http.StatusBadRequest,
		gin.H{
			"success": false,
			"error":   fmt.Sprintf("Invalid body: %s", err.Error()),
		})
}

func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		badRequest(ctx, err)
		return
	}

	useCase := use_cases.NewCreateCategoryUseCase(repository)

	err := useCase.Execute(body.Name)

	if err != nil {
		badRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
