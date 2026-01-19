package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) {
		c.JSON("Api is running")
	})

	app.Listen(":3030")
}