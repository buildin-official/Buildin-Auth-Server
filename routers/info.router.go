package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/BuildinAuth/controllers"
	"pentag.kr/BuildinAuth/middlewares"
)

func info(router *fiber.App) {
	info := router.Group("/info")
	info.Use(middlewares.JWTMiddlware())
	info.Get("/", func(c *fiber.Ctx) error {
		return controllers.CheckUser(c)
	})
	info.Get("/info", func(c *fiber.Ctx) error {
		return controllers.GetInfo(c)
	})
	info.Post("/changeUsername", func(c *fiber.Ctx) error {
		return controllers.ChangeUsername(c)
	})
}
