package handlers

import (
	"net/http"

	"github.com/TonySultan/expense-tracker/pkg/models"
	"github.com/gin-gonic/gin"
)

var accounts []models.Account

/*
map = {
	userId: {
		accountName: accountCash
	}
}
*/


func GetAccounts(ctx *gin.Context) {
	userID, _ := ctx.Get("id") // .Param("id")
	userID = userID.(int) // TODO: int, err := strconv.Atoi("12345")
	// TODO: логиканы сайкесинше озгерту
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
		UserId int `json:"userId"` // TODO: барлык пост бади дурстау
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
	accounts = append(accounts, newAccount)
	
	// TODO: send newAccount as response
	ctx.JSON(http.StatusCreated, "newAccount")
}

func userHasAccount(userId int, accountName string) bool {
	// accounts.find()

	return true
}