package main

import (
	"log"

	"github.com/naveenkakumanu/todoapp/db"
	"github.com/naveenkakumanu/todoapp/routes"
)

func main() {
	log.Println("TODO APP USING Gin Framework and Postgress")
	router := routes.Routes()
	db.DB()
	router.Run(":8090")
}
