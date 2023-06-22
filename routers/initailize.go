package routers

import (
	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {

	index(router)
	auth(router)
	info(router)
	withdrawal(router)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
