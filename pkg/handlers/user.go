package handlers

import (
	"log"
	"net/http"

	"github.com/TonySultan/expense-tracker/pkg/models"
	"github.com/TonySultan/expense-tracker/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type RegisterResponse struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func RegisterHandler(ctx *gin.Context) {
	resp := RegisterResponse{}
	err := ctx.BindJSON(&resp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		log.Printf("RegisterHandler: BindJSON error: %s", err.Error())
		return
	}
	if resp.Login == "" || resp.Password == "" {
		ctx.JSON(http.StatusBadRequest, "login or password is empty")
		log.Printf("RegisterHandler: BindJSON error: login or password is empty")
		return
	}

	id, err := repositories.CreateUser(resp.Login, resp.Password)
	ctx.JSON(http.StatusOK, RegisterRequest{ID: id, Message: "ok"})
}

func LoginHandler(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storedUser, exists := repositories.Users[user.Login]
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}

	if user.Password != storedUser.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный пароль"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Вход выполнен успешно!"})
}
