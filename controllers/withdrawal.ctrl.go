package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"pentag.kr/BuildinAuth/database"
	"pentag.kr/BuildinAuth/models"
	"pentag.kr/BuildinAuth/utils"
)

func WithdrawalAccount(c *fiber.Ctx) error { // 회원 탈퇴 컨트롤러
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json := new(LoginRequest)
	if err := c.BodyParser(json); err != nil || json.Email == "" || json.Password == "" { // json을 구조체로 파싱하고, 이메일과 비밀번호가 비어있는지 확인
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	db := database.DB
	found := models.User{}
	query := models.User{Email: json.Email}
	err := db.First(&found, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "User not found",
		})
	}

	if !utils.ComparePasswords(found.Password, json.Password) {
		return c.Status(401).JSON(fiber.Map{
			"code":    401,
			"message": "Invalid Password",
		})
	}
	db.Delete(&found)

	allUserRefreshTokens := []models.RefreshToken{}
	db.Where("user_id = ?", found.ID).Find(&allUserRefreshTokens)
	for _, refreshToken := range allUserRefreshTokens {
		db.Delete(&refreshToken)
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "OK",
	})
}
