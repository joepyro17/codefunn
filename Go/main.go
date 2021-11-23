package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
	database "pyro.com/codefunn/databases"
	. "pyro.com/codefunn/routers"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Fiber
	app := fiber.New()
	app.Use(cors.New())

	// DB
	database.DbConnection()

	// Routes
	SetupRoutes(app)

	port := os.Getenv("FIBER_PORT")

	err = app.Listen(port)
	if err != nil {
		log.Fatal("Error Listen to port 3000")
	}
}