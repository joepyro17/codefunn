package handlers

import (
	"github.com/gofiber/fiber/v2"
	database "pyro.com/codefunn/databases"
	"pyro.com/codefunn/models"
)
func CreateCategory(ctx *fiber.Ctx) error {
	db := database.DB
	category := new(models.Category)

	err := ctx.BodyParser(category)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "CreateCategory: Review Your Inputs", "data": err})
	}

	err = db.Create(&category).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Could Not Create a New Category", "data": err})
	}

	return ctx.Status(201).JSON(fiber.Map{"status": "success", "message": "Created Category", "data": category})
}

func GetCategories(ctx *fiber.Ctx) error {
	db := database.DB
	var categories []models.Category

	db.Find(&categories)

	if len(categories) == 0 {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No Categories Present", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Categories Found", "data": categories})
}

func GetCategoryById(ctx *fiber.Ctx) error {
	db := database.DB
	category := new(models.Category)

	db.Find(&category, "id=?", ctx.Params("id"))

	if category.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{"status":"error", "message":"No Category Present", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status":"success", "message":"Category Found", "data": category})


}

func UpdateCategory(ctx *fiber.Ctx) error{
	db := database.DB
	category := new(models.Category)

	db.Find(&category, "id=?", ctx.Params("id"))

	if category.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No Category Present", "data": nil})
	}

	newCategory := new(models.Category)

	err := ctx.BodyParser(newCategory)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review Your Inputs...", "data": err})
	}

	category.Name = newCategory.Name

	err = db.Save(&category).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Could Not Update Category", "data": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Updated Category", "data": category})
}

func DeleteCategory(ctx *fiber.Ctx) error {
	db := database.DB
	category := new(models.Category)

	db.Find(&category, "id=?", ctx.Params("id"))
	if category.ID == 0 {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No Category Present", "data": nil})
	}

	err := db.Delete(&category, "id", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to Delete Category", "data": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Deleted Category"})

}