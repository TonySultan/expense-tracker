package main

import (
	"github.com/TonySultan/expense-tracker/pkg/handlers"
	"github.com/TonySultan/expense-tracker/pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/registration", handlers.RegisterHandler)
	r.POST("/login", handlers.LoginHandler)

	r.GET("/accounts/:id", handlers.GetAccounts)
	r.POST("/accounts", handlers.CreateAccount)

	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)

	r.GET("/transactions", handlers.GetTransactions)
	r.POST("/transactions", handlers.CreateTransaction)

	r.Run(":8080")
}

type ApiAnswer struct {
	Message string      `json:"msg"`
	User    models.User `json:"user"`
}
