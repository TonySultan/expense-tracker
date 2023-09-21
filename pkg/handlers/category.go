package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}

func CreateCategory(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
