package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	. "pyro.com/codefunn/configs"
	. "pyro.com/codefunn/routers"
)

func main() {
	// Fiber
	app := fiber.New()
	app.Use(cors.New())

	// Routes
	SetupRoutes(app)

	err := app.Listen(PORT)
	if err != nil {
		log.Fatal("Error Listen to port 3000")
	}
}