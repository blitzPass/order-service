package main

import (
	// "context"
	"context"
	"fmt"
	"log"
	"order-service/infrastructure"
	"order-service/infrastructure/kafka"
	"os/signal"
	"syscall"

	// "order-service/infrastructure/kafka"
	// "order-service/infrastructure/kafka/handlers"
	"order-service/pkg"
	"os"

	// "os/signal"
	// "syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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


	// Kafka Consumer
	consumer := kafka.NewConsumer(
		[]string{"localhost:9092"},
		"order.events",
		"order-service",
	)

	go consumer.Start(ctx)

	
	app.Get("/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		if id == "error" {
			return fiber.NewError(fiber.StatusInternalServerError, "this is an error")
		}

		return c.JSON("Api is running")
	})

	// Graceful shutdown
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		log.Println("shutdown signal received")
		cancel()
		consumer.Close()
		_ = app.Shutdown()
	}()



	log.Fatal(app.Listen(":3030"))
}