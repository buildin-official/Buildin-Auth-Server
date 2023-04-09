package controllers

import (
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
	"pentag.kr/BuildinAuth/database"
	"pentag.kr/BuildinAuth/models"
	"pentag.kr/BuildinAuth/utils"
)

func CheckUser(c *fiber.Ctx) error { // 유저가 있는지 확인하는 컨트롤러
	jwtClaims := c.Locals("jwtClaims").(utils.AuthTokenClaims)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "OK",
		"user-id": jwtClaims.UserID,
	})
}

func GetInfo(c *fiber.Ctx) error { // 유저 정보를 가져오는 컨트롤러
	jwtClaims := c.Locals("jwtClaims").(utils.AuthTokenClaims)
	db := database.DB
	user := models.User{}
	userUUID, err := guuid.Parse(jwtClaims.UserID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": "Internal Server Error",
		})
	}
	query := models.User{ID: userUUID}
	err = db.First(&user, &query).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": "Internal Server Error",
		})
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "OK",
		"email":   user.Email,
		"user-id": jwtClaims.UserID,
	})

}
