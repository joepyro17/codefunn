package routers

import (
	"github.com/gofiber/fiber/v2"
	handler "pyro.com/codefunn/handlers"
)

func SetupWebAPIRoutes(app *fiber.App){

	//api
	api := app.Group("/web_api")

	//Blogs
	api.Post("/blogs", handler.CreateBlog)
	api.Get("/blogs", handler.GetBlogs)
	api.Get("/blogs/:id", handler.GetBlogById)
	api.Put("/blogs/:id", handler.UpdateBlog)
	api.Delete("/blogs/:id", handler.DeleteBlog)


	// IP
	api.Get("/get_ip", handler.GetIP)
}