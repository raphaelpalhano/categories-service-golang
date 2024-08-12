package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raphaelpalhano/categories-ms/cmd/internal/repositories"
	use_cases "github.com/raphaelpalhano/categories-ms/cmd/internal/use-cases"
	"github.com/raphaelpalhano/categories-ms/cmd/pkg/utils"
)

func ListCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {

	useCase := use_cases.NewListCategoryUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		utils.BadRequest(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "categories": categories})
}
