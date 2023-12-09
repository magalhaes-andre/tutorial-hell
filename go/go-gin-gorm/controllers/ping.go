package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"Message":      "Hi there!",
		"Current time": time.Now(),
	})
}
