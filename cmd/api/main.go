package main

import (
	"log"
	"authentication-service/internal/config"
	"authentication-service/internal/delivery/http/route"
	"authentication-service/internal/repository/postgresql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	log.Println(cfg)

	db, err := postgresql.NewConnector(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	cfg.DB = db

	err = postgresql.Migrate(db)
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New()

	app.Use(logger.New())

	// Setup routes
	route.SetupRoutes(app, cfg)

	// Start server
	log.Fatal(app.Listen(cfg.ServerAddress))
}
