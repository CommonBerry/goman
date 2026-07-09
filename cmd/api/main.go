package main

import (
	"log"

	"github.com/CommonBerry/goman/cmd/routes"
	"github.com/CommonBerry/goman/internal/infra"
	"github.com/gofiber/fiber/v3"
)

func main() {
	port := ":3000"
	app := fiber.New()

	routes.SetupRoutes(app, new(infra.DataBase))

	log.Fatal(app.Listen(port))
}
