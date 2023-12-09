package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/magalhaes-andre/models"
)

func GetAllUsers(context *gin.Context) {
	context.JSON(200, models.Users)
}
