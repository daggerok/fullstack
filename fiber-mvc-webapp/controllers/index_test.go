package controllers_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
	"github.com/stretchr/testify/require"

	"fiber-mvc-webapp/controllers"
)

func TestIndexPage(t *testing.T) {
	// Mustache template engine
	engine := mustache.New("../views", ".mustache")

	// Web MVC
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "../static")

	// Register the IndexPage handler with the test app
	app.Get("/", controllers.IndexPage)

	// Create a request with a query parameter
	req, err := http.NewRequest("GET", "/?name=Test", nil)
	require.NoError(t, err)

	// Call the handler with the request
	resp, err := app.Test(req, 1)
	require.NoError(t, err)

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	// Assert expected response status and content
	expectedBody := "Hello Test!"
	require.Contains(t, string(body), expectedBody)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
