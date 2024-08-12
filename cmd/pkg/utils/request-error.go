package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(
		http.StatusBadRequest,
		gin.H{
			"success": false,
			"error":   fmt.Sprintf("Invalid body: %s", err.Error()),
		})
}
