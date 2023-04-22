package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/BuildinAuth/utils"
)

func JWTMiddlware() fiber.Handler {

	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization", "")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(401).JSON(fiber.Map{
				"error": "Bad Authorization header",
			})
		}

		accessToken := strings.TrimPrefix(authHeader, "Bearer ")
		// we parse our jwt token and check for validity against our secret
		claims, err := utils.ParseToken(accessToken)

		if err == nil {
			fmt.Println(claims.UserID)
			c.Locals("jwtClaims", *claims)
			return c.Next()
		}

		return c.Status(498).JSON(fiber.Map{
			"error": "Bad Access Token",
		})
	}
}
