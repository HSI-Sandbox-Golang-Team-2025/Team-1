package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/orgs/HSI-Sandbox-Golang-Team-2025/pkg/config"
)

func main() {
	cfg := config.LoadEnv()
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()
	log.Printf("Starting server in %s mode on port %s", cfg.AppEnv, cfg.AppPort)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(cfg.AppEnv)
	})

	app.Listen(":" + cfg.AppPort)
}
