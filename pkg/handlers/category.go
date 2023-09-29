package handlers

import (
	"net/http"

	"github.com/TonySultan/expense-tracker/pkg/models"
	"github.com/gin-gonic/gin"
)

var categories []models.Category

func GetCategories(ctx *gin.Context) {
	userID, _ := ctx.Get("id")
	userID = userID.(int)
	userCategories := []models.Category{}

	for _, category := range categories {
		if category.UserId == userID {
			userCategories = append(userCategories, category)
		}
	}
	ctx.JSON(http.StatusOK, userCategories)
}

func CreateCategory(ctx *gin.Context) {

	var inputData struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := ctx.Get("id")
	userID = userID.(int)

	newCategory := models.Category{
		Name:   inputData.Name,
		Type:   inputData.Type,
		UserId: userID,
	}

	categories = append(categories, newCategory)

	ctx.JSON(http.StatusCreated, newCategory)
}
