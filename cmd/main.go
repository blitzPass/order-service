package main

import (
	"fmt"
	"os"
	"ticket-service/infrastructure"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := os.Getenv("DB_SSL_MODE")
	
	// Database connection
	var dataSource string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, dbSslMode)
	db, _ := infrastructure.NewPostgresConnection(dataSource)
	err := db.Ping()
	if err != nil {
		panic("Database error: " + err.Error())
	}


	app.Get("/", func(c fiber.Ctx) {
		c.JSON("Api is running")
	})

	app.Listen(":3030")
}