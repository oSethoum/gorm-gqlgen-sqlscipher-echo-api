package main

import (
	"gormql/db"
	"gormql/graph/models"
	"gormql/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&models.User{}, &models.Todo{})
	e := echo.New()
	routes.Init(e)

	log.Fatal(e.Start(":5000"))
}
