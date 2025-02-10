package main

import (
	"github.com/azkifairuz/catatin-api/config"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config.ConnectDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("api running succesfully!")
	})

	app.Listen(":3000")
}