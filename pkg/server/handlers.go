package server

import "github.com/gofiber/fiber/v2"

type TodoHandler struct{}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

func (t *TodoHandler) Home(ctx *fiber.Ctx) error {
	context := fiber.Map{}
	return ctx.Render("index", context)
}
