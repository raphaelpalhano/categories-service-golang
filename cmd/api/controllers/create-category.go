package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelpalhano/categories-ms/cmd/internal/repositories"
	use_cases "github.com/raphaelpalhano/categories-ms/cmd/internal/use-cases"
	"github.com/raphaelpalhano/categories-ms/cmd/pkg/utils"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.BadRequest(ctx, err)
		return
	}

	useCase := use_cases.NewCreateCategoryUseCase(repository)

	err := useCase.Execute(body.Name)

	if err != nil {
		utils.BadRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
