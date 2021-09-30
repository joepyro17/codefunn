package routers

import (
	"github.com/gofiber/fiber/v2"
	. "pyro.com/codefunn/mockups"
)

func SetupAPIRoutes(app *fiber.App){

	//api
	api := app.Group("/api")

	api.Get("/tutorials", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Get all Tutorials")
	})
	
	// mockups
	api.Get("/mockups/tutorials", func(ctx *fiber.Ctx) error {
		return ctx.JSON(GetAllTutorials())
	})
}