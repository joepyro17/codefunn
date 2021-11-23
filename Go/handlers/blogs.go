package handlers

import (
	"github.com/gofiber/fiber/v2"
	database "pyro.com/codefunn/databases"
	. "pyro.com/codefunn/models"
)


func CreateBlog(ctx *fiber.Ctx) error {
	db := database.DB
	blog := new(Blog)

	err := ctx.BodyParser(blog)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review Your Inputs", "data": err})
	}

	err = db.Create(&blog).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Could Not Create a New Blog", "data": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Created Blog", "data": blog})

}

func GetBlogs(ctx *fiber.Ctx) error {
	db := database.DB
	var blogs []Blog

	db.Find(&blogs)

	if len(blogs) == 0 {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No Blogs Present", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Blogs Found", "data": blogs})
}

func GetBlogById(ctx *fiber.Ctx) error {
	db := database.DB
	var blog Blog

	db.Find(&blog, "ID=?", ctx.Params("id"))

	if blog.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No Blog Present", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Blog Found", "data": blog})

}

func UpdateBlog(ctx *fiber.Ctx) error {
	db := database.DB
	blog := new(Blog)

	db.Find(&blog, "id=?", ctx.Params("id"))

	if blog.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No Blog Present", "data": nil})
	}

	newBlog := new(Blog)

	err := ctx.BodyParser(newBlog)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review Your Inputs...", "data": err})
	}

	blog.Title = newBlog.Title
	blog.Content = newBlog.Content
	blog.FeatureImageURL = newBlog.FeatureImageURL

	err = db.Save(&blog).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Could Not Update blog", "data": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Updated Blog", "data": blog})
}

func DeleteBlog(ctx *fiber.Ctx) error {
	db := database.DB
	blog := new(Blog)

	db.Find(&blog, "id=?", ctx.Params("id"))
	if blog.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No Blog Present", "data": nil})
	}

	err := db.Delete(&blog, "id=?").Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to Delete Blog", "data": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Deleted Blog"})
}