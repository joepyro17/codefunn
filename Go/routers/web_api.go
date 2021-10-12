package routers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"net"
	topic "pyro.com/codefunn/services/topics"
)

func SetupWebAPIRoutes(app *fiber.App, db *gorm.DB){

	//api
	api := app.Group("/web_api")

	//Topics
	api.Get("/topics", func(ctx *fiber.Ctx) error {
		allTopics := topic.GetAllTopics(db)
		return ctx.JSON(allTopics)
	})
	api.Get("/topics/:id", func(ctx *fiber.Ctx) error {
		topicById := topic.GetTopicById(db, ctx.Params("id"))
		return ctx.JSON(topicById)
	})
	api.Post("/topics", func(ctx *fiber.Ctx) error {
		newTopic, err := topic.AddTopic(db, ctx)

		if err != nil {
			return ctx.JSON(err)
		}

		return ctx.JSON(newTopic)
	})

	api.Get("/get_ip", func(ctx *fiber.Ctx) error {
		addresses, err := net.InterfaceAddrs()
		if err != nil {
			log.Fatal(err)
		}

		var currentIP string

		for _, addr := range addresses {

			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					currentIP = ipnet.IP.String()
				}
			}
		}

		return ctx.SendString("IP address = " + currentIP)
	})

	//api.Get("/mockups/tutorials", func(ctx *fiber.Ctx) error {
	//	return ctx.JSON(GetAllTutorials())
	//})
}