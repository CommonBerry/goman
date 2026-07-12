package main

import (
	"context"
	"log"
	"os"

	"github.com/CommonBerry/goman/cmd/routes"
	"github.com/CommonBerry/goman/internal/infra"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load enviroment file: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if port[0] != ':' {
		port = ":" + port
	}

	app := fiber.New()

	ctx := context.Background()

	dataBase, err := infra.NewPostgresDataBase(ctx)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	routes.SetupRoutes(app, dataBase)

	log.Fatal(app.Listen(port))
}
