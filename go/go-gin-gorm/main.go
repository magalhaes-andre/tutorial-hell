package main

import (
	"github.com/magalhaes-andre/models"
	"github.com/magalhaes-andre/routes"
)

func main() {
	models.Users = []models.User{
		{Name: "André", Age: 23},
		{Name: "Luciana", Age: 24},
	}
	routes.HandleRequests()
}
