package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"

	"fiber-mvc-webapp/controllers"
)

func main() {
	// Mustache template engine
	engine := mustache.New("./views", ".mustache")

	// Web MVC
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./static")

	// Index page handler
	app.Get("/", controllers.IndexPage)

	// Server listen port 8080
	err := app.Listen(":8080")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
