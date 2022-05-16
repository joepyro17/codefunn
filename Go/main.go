package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	database "pyro.com/codefunn/databases"
	"pyro.com/codefunn/routers"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Fiber
	app := fiber.New()

	// Logging
	file, logErr := os.OpenFile("./api_log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if logErr != nil {
		log.Fatalf("Error Opening File: %v", logErr)
	}
	defer file.Close()

	// Config file
	loggerConfig := logger.Config{
		Next:   nil,
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${body} ${resBody}\n",
		Output: file,
	}

	app.Use(logger.New(loggerConfig))

	app.Use(cors.New())

	// DB
	database.DbConnection()

	// Routes
	routers.SetupRoutes(app)

	port := os.Getenv("FIBER_PORT")

	err = app.Listen(port)
	if err != nil {
		log.Fatal("Error Listen to port 3000")
	}
}
