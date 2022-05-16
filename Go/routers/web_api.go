package routers

import (
	"github.com/gofiber/fiber/v2"
	handler "pyro.com/codefunn/handlers"
)

func SetupWebAPIRoutes(app *fiber.App) {

	//api
	api := app.Group("/web_api")

	//Blogs
	api.Post("/blogs", handler.CreateBlog)
	api.Get("/blogs", handler.GetBlogs)
	api.Get("/blogs/:id", handler.GetBlogById)
	api.Put("/blogs/:id", handler.UpdateBlog)
	api.Delete("/blogs/:id", handler.DeleteBlog)

	api.Get("/blogs/category/:category_id", handler.GetBlogByCategoryId)

	//Categories
	api.Post("/categories", handler.CreateCategory)
	api.Get("/categories", handler.GetCategories)
	api.Get("/categories/:id", handler.GetCategoryById)
	api.Put("/categories/:id", handler.UpdateCategory)
	api.Delete("/categories/:id", handler.DeleteCategory)

	// IP
	api.Get("/get_ip", handler.GetIP)
}
