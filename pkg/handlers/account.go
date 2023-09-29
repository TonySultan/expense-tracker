package handlers

import (
	"net/http"

	"github.com/TonySultan/expense-tracker/pkg/models"
	"github.com/gin-gonic/gin"
)

var accounts []models.Account

func GetAccounts(ctx *gin.Context) {
	userID, _ := ctx.Get("id")
	userID = userID.(int)
	userAccounts := []models.Account{}
	for _, account := range accounts {
		if account.UserId == userID {
			userAccounts = append(userAccounts, account)
		}
	}
	ctx.JSON(http.StatusOK, userAccounts)
}

func CreateAccount(ctx *gin.Context) {
	var inputData struct {
		Name string  `json:"name"`
		Cash float64 `json:"cash"`
	}

	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := ctx.Get("id")
	userID = userID.(int)
	newAccount := models.Account{
		Name:   inputData.Name,
		Cash:   inputData.Cash,
		UserId: userID,
	}
	accounts = append(accounts, newAccount)
	ctx.JSON(http.StatusCreated, "newAccount")

}
