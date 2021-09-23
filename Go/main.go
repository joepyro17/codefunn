package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"io/ioutil"
	"log"
	"os"
	Models "pyro.com/codefunn/models"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	jsonFile, err := os.Open("./mockups/tutorials.json")
	if err != nil {
		log.Fatal("Open Json failed")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tutorials Models.Tutorials

	json.Unmarshal(byteValue, &tutorials)

	//for i:= 0; i< len(tutorials.Tutorials); i++ {
	//	fmt.Println(tutorials.Tutorials[i].Title)
	//}

	app.Get("/tutorials", func(ctx *fiber.Ctx) error {
		return ctx.JSON(tutorials)
	})

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Error Listen to port 3000")
	}
}