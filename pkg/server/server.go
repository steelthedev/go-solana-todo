package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func StartServer(port string) {

	viewsEngine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: viewsEngine,
	})

	// Intialize todo handler
	todoHandler := NewTodoHandler()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("I am the todo")
	})

	// Register all todo routes
	todoHandler.RegisterRoutes(app)

	if port != "" {
		app.Listen(port)
	} else {
		app.Listen("8000")
	}

}
