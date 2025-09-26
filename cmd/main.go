package main

import (
	"log"
	"siduk/config"
	"siduk/routes"
	"siduk/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.ConnectDB()
	if err := models.AutoMigrateKeluarga(db); err != nil {
		log.Fatal(err)
	}
	if err := models.AutoMigratePenduduk(db); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	routes.SetupRoutes(app, db)
	app.Static("/public", "./public")

	log.Fatal(app.Listen(":8080"))
}