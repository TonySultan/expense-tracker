package handlers

import (
	"net/http"
	"strconv"

	"github.com/TonySultan/expense-tracker/pkg/models"
	"github.com/gin-gonic/gin"
)

var accounts []models.Account

var UserAccounts = make(map[string]models.Account)

func GetAccounts(ctx *gin.Context) {
	userID := ctx.Param("id")
	// intID, err := strconv.Atoi(userID)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// TODO: логиканы сайке синше озгерту
	// userAccounts := []models.Account{}
	userAccounts := UserAccounts[userID]
	ctx.JSON(http.StatusOK, userAccounts)
}

func CreateAccount(ctx *gin.Context) {
	var inputData struct {
		Name   string  `json:"name"`
		Cash   float64 `json:"cash"`
		UserId int
	}

	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: call the function
	// userHasAccount()
	// TODO: жана страктка сай озгерту

	newAccount := models.Account{
		Name:   inputData.Name,
		Cash:   inputData.Cash,
		UserId: inputData.UserId,
	}
	strID := strconv.Itoa(inputData.UserId)
	UserAccounts[strID] = newAccount

	// TODO: send newAccount as response
	ctx.JSON(http.StatusCreated, newAccount)
}

// func userHasAccount(userId int, accountName string) bool {
// 	// accounts.find()

// 	return true
// }
