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
		"code":     200,
		"message":  "OK",
		"email":    user.Email,
		"username": user.Username,
		"user-id":  jwtClaims.UserID,
	})

}

func ChangeUsername(c *fiber.Ctx) error { // 닉네임 변경 컨트롤러
	type ChangeUsernameRequest struct {
		Username string `json:"username" validate:"required,min=1,max=20,excludesall=;"`
	}
	json := new(ChangeUsernameRequest)
	if err := c.BodyParser(json); err != nil || json.Username == "" { // json을 구조체로 파싱하고, 닉네임이 비어있는지 확인
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
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
	user.Username = json.Username
	db.Save(&user)

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "OK",
	})
}
