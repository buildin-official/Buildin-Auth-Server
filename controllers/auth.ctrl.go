package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	guuid "github.com/google/uuid"
	"pentag.kr/BuildinAuth/database"
	"pentag.kr/BuildinAuth/models"
	"pentag.kr/BuildinAuth/utils"
)

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	db := database.DB
	json := new(LoginRequest)
	if err := c.BodyParser(json); err != nil || json.Email == "" || json.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

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
	newUuid, _ := guuid.NewRandom()
	refreshToken := newUuid.String()
	database.RDB.SetEx(c.Context(), refreshToken, found.ID.String(), time.Hour*24*7)
	accessToken, err := utils.CreateToken(found.ID.String())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": "Internal Server Error",
		})
	}
	return c.JSON(fiber.Map{
		"code":          200,
		"message":       "sucess",
		"access-token":  accessToken,
		"refresh-token": refreshToken,
	})
}

func Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=30,excludesall=;"`
	}
	db := database.DB
	json := new(RegisterRequest)
	if err := c.BodyParser(json); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	if err := utils.ValidateStruct(json); err != nil || !utils.VerifyPassword(json.Password) {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	foundUnvalidatedUser := models.UnvalidatedUser{}
	unvalidatedUserquery := models.UnvalidatedUser{Email: json.Email}
	if err := db.First(&foundUnvalidatedUser, &unvalidatedUserquery).Error; err != gorm.ErrRecordNotFound {
		return c.Status(409).JSON(fiber.Map{
			"code":    409,
			"message": "User already exists",
		})
	}

	foundUser := models.User{}
	userQuery := models.User{Email: json.Email}
	if err := db.First(&foundUser, &userQuery).Error; err != gorm.ErrRecordNotFound {
		return c.Status(409).JSON(fiber.Map{
			"code":    409,
			"message": "User already exists",
		})
	}

	userTempID := guuid.New()

	validationEmail := utils.Mail{}
	err := validationEmail.SendVerificationEmail(json.Email, userTempID.String())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": "Internal Server Error",
		})
	}

	hashedPassword := utils.HashPassword(json.Password)
	newUnvalidatedUser := models.UnvalidatedUser{ID: userTempID, Email: json.Email, Password: hashedPassword}
	db.Create(&newUnvalidatedUser)

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func Refresh(c *fiber.Ctx) error {
	refreshToken := c.Query("refresh-token", "")
	if refreshToken == "" {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	userId, err := database.RDB.Get(c.Context(), refreshToken).Result()
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"code":    401,
			"message": "Unauthorized",
		})
	}
	accessToken, err := utils.CreateToken(userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": "Internal Server Error",
		})
	}
	return c.JSON(fiber.Map{
		"code":         200,
		"message":      "sucess",
		"access-token": accessToken,
	})
}

func Verify(c *fiber.Ctx) error {
	token := c.Query("token", "")
	userUUID, err := guuid.Parse(token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Qurey or Token",
		})
	}
	db := database.DB
	foundUnvalidatedUser := models.UnvalidatedUser{}
	unvalidatedUserquery := models.UnvalidatedUser{ID: userUUID}
	if err := db.First(&foundUnvalidatedUser, &unvalidatedUserquery).Error; err == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "User not found",
		})
	}
	newUserID := guuid.New()
	newUser := models.User{ID: newUserID, Email: foundUnvalidatedUser.Email, Password: foundUnvalidatedUser.Password}
	db.Create(&newUser)
	db.Delete(&foundUnvalidatedUser)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}
