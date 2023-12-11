package main

import (
	"github.com/magalhaes-andre/database"
	"github.com/magalhaes-andre/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
