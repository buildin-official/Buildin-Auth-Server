package routers

import (
	"github.com/gofiber/fiber/v2"
	"pentag.kr/BuildinAuth/controllers"
)

func withdrawal(router *fiber.App) {
	withdrawal := router.Group("/withdrawal")
	withdrawal.Get("/", func(c *fiber.Ctx) error {
		return controllers.WithdrawalAccount(c)
	})
}
