package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"fiber-mvc-webapp/models"
)

func IndexPage(c *fiber.Ctx) error {
	name := utils.CopyString(c.Query("name", "Guest"))
	fmt.Printf("name: %v\n", name)
	// return c.SendString(fmt.Sprintf("Hello, %s!", name))
	return c.Render("index", models.User{
		Name: name,
	})
}
