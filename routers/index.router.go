package routers

import (
	"github.com/gofiber/fiber/v2"
)

func index(router *fiber.App) {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

}
