package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactions(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}

func CreateTransaction(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
