package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raphaelpalhano/categories-ms/cmd/api/controllers"
	"github.com/raphaelpalhano/categories-ms/cmd/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.GET("/", func(ctx *gin.Context) {
		controllers.ListCategory(ctx, inMemoryCategoryRepository)
	})
}
