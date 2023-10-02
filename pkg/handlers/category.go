package handlers

import (
	"net/http"
	"strconv"

	"github.com/TonySultan/expense-tracker/pkg/models"
	"github.com/gin-gonic/gin"
)

var categories []models.Category

func GetCategories(ctx *gin.Context) {
	userID := ctx.Param("id")
	intID, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userCategories := []models.Category{}

	for _, category := range categories {
		if category.UserId == intID {
			userCategories = append(userCategories, category)
		}
	}
	ctx.JSON(http.StatusOK, userCategories)
}

func CreateCategory(ctx *gin.Context) {

	var inputData struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		UserId int
	}

	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := &inputData.UserId

	newCategory := models.Category{
		Name:   inputData.Name,
		Type:   inputData.Type,
		UserId: *userID,
	}

	categories = append(categories, newCategory)

	ctx.JSON(http.StatusCreated, newCategory)
}
