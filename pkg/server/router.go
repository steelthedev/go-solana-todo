package server

import "github.com/gofiber/fiber/v2"

func (th *TodoHandler) RegisterRoutes(router *fiber.App) {
	router.Get("/", th.Home)
}
