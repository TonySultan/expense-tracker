package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccounts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}

func CreateAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")

}
