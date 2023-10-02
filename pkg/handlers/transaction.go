package handlers

import (
	"net/http"
	"strconv"

	"github.com/TonySultan/expense-tracker/pkg/models"
	"github.com/gin-gonic/gin"
)

var transactions []models.Transaction

func GetTransactions(ctx *gin.Context) {
	userID := ctx.Param("id")
	intID, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userTransactions := []models.Transaction{}

	for _, transaction := range transactions {
		var accountExists, categoryExists bool
		for _, acc := range accounts {
			if acc.UserId == transaction.AccountID && acc.UserId == intID {
				accountExists = true
				break
			}
		}
		for _, cat := range categories {
			if cat.UserId == transaction.CategoryID && cat.UserId == intID {
				categoryExists = true
				break
			}
		}

		if accountExists && categoryExists {
			userTransactions = append(userTransactions, transaction)
		}
	}
	ctx.JSON(http.StatusOK, userTransactions)

}

func CreateTransaction(ctx *gin.Context) {
	var inputData struct {
		CategoryID int     `json:"categoryId"`
		AccountID  int     `json:"accountId"`
		Amount     float64 `json:"amount"`
		Comment    string  `json:"comment"`
		UserId     int     `json:"userId"`
	}
	var accountExists, categoryExists bool
	var account models.Account

	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := &inputData.UserId

	for _, acc := range accounts {
		if acc.UserId == inputData.AccountID && acc.UserId == *userID {
			accountExists = true
			account = acc
			break
		}
	}

	for _, cat := range categories {
		if cat.UserId == inputData.CategoryID && cat.UserId == *userID {
			categoryExists = true
			break
		}
	}

	if !accountExists || !categoryExists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account or category"})
		return
	}
	newTransaction := models.Transaction{
		Amount:     inputData.Amount,
		Comment:    inputData.Comment,
		AccountID:  inputData.AccountID,
		CategoryID: inputData.CategoryID,
	}
	transactions = append(transactions, newTransaction)

	account.Cash += newTransaction.Amount
	ctx.JSON(http.StatusCreated, newTransaction)

}
