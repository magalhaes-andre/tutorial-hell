package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/magalhaes-andre/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.GET("/users", controllers.GetAllUsers)
	r.Run(":8085")
}
