package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/magalhaes-andre/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.FetchUserById)
	r.GET("/users/age/:age", controllers.FetchByAge)
	r.POST("/users", controllers.CreateUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUserById)
	r.Run(":8085")
}
