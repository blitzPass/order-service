package main

import (
	"fmt"
	"log"
	"order-service/infrastructure"
	"order-service/pkg"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New(fiber.Config{
		ErrorHandler: pkg.ErrorHandler(),
	})

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	dataSource := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, dbSslMode,
	)

	db, _ := infrastructure.NewPostgresConnection(dataSource)
	if err := db.Ping(); err != nil {
		panic("Database error: " + err.Error())
	}

	app.Get("/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		if id == "error" {
			return fiber.NewError(fiber.StatusInternalServerError, "this is an error")
		}

		return c.JSON("Api is running")
	})

	log.Fatal(app.Listen(":3030"))
}
