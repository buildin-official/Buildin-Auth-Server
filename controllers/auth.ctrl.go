package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	guuid "github.com/google/uuid"
	"pentag.kr/BuildinAuth/database"
	"pentag.kr/BuildinAuth/models"
	"pentag.kr/BuildinAuth/utils"
)

func Login(c *fiber.Ctx) error { // 로그인 컨트롤러
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
	newUUID, _ := guuid.NewRandom()
	newRefreshToken := models.RefreshToken{
		ID:     newUUID,
		UserID: found.ID,
	}
	db.Create(&newRefreshToken)
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
		"refresh-token": newUUID.String(),
	})
}

func Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Username string `json:"username" validate:"required,min=1,max=20,excludesall=;"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=30,excludesall=;"`
	}

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

	db := database.DB

	foundUnvalidatedUser := models.UnvalidatedUser{}
	unvalidatedUserquery := models.UnvalidatedUser{Email: json.Email}
	if err := db.First(&foundUnvalidatedUser, &unvalidatedUserquery).Error; err != gorm.ErrRecordNotFound {
		return c.Status(409).JSON(fiber.Map{
			"code":    410,
			"message": "Already registered, but not verified",
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
	newUnvalidatedUser := models.UnvalidatedUser{ID: userTempID, Email: json.Email, Password: hashedPassword, Username: json.Username}
	db.Create(&newUnvalidatedUser)

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func Refresh(c *fiber.Ctx) error {
	refreshTokenString := c.Query("refresh-token", "")
	refreshTokenUUID, err := guuid.Parse(refreshTokenString)
	db := database.DB
	err, refreshToken := models.FindRefreshToken(db, refreshTokenUUID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "Refresh Token not found",
		})
	}
	accessToken, err := utils.CreateToken(refreshToken.UserID.String())
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

func Logout(c *fiber.Ctx) error {
	refreshTokenString := c.Query("refresh-token", "")
	db := database.DB
	refreshTokenUUID, err := guuid.Parse(refreshTokenString)
	err, refreshToken := models.FindRefreshToken(db, refreshTokenUUID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "Refresh Token not found",
		})
	}
	db.Delete(&refreshToken)

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
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
	newUser := models.User{ID: newUserID, Email: foundUnvalidatedUser.Email, Password: foundUnvalidatedUser.Password, Username: foundUnvalidatedUser.Username}
	db.Create(&newUser)
	db.Delete(&foundUnvalidatedUser)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}

func RequestChangePassword(c *fiber.Ctx) error {
	type RequestChangePasswordRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=30,excludesall=;"`
	}

	json := new(RequestChangePasswordRequest)
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

	db := database.DB

	foundUser := models.User{}
	Userquery := models.User{Email: json.Email}
	if err := db.First(&foundUser, &Userquery).Error; err == gorm.ErrRecordNotFound {
		return c.Status(200).JSON(fiber.Map{
			"code":    200,
			"message": "success",
		})
	}

	changePasswordCode := guuid.New()
	hashedPassword := utils.HashPassword(json.Password)

	newChangePasswordObject := models.ChangePasswordCode{ID: changePasswordCode, UserID: foundUser.ID, Password: hashedPassword}
	db.Create(&newChangePasswordObject)

	changePasswordEmail := utils.Mail{}

	err := changePasswordEmail.SendChangePasswordEmail(json.Email, changePasswordCode.String())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": "Internal Server Error",
		})
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})

}

func VerifyChangePassword(c *fiber.Ctx) error {
	token := c.Query("token", "")
	changePasswordCode, err := guuid.Parse(token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Qurey or Token",
		})
	}
	db := database.DB
	foundChangePasswordCode := models.ChangePasswordCode{}
	changePasswordCodeQuery := models.ChangePasswordCode{ID: changePasswordCode}
	if err := db.First(&foundChangePasswordCode, &changePasswordCodeQuery).Error; err == gorm.ErrRecordNotFound {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Qurey or Token",
		})
	}
	foundUser := models.User{}
	userQuery := models.User{ID: foundChangePasswordCode.UserID}
	if err := db.First(&foundUser, &userQuery).Error; err == gorm.ErrRecordNotFound {
		return c.Status(400).JSON(fiber.Map{
			"code":    500,
			"message": "Internal Server Error",
		})
	}
	foundUser.Password = foundChangePasswordCode.Password
	db.Save(&foundUser)
	db.Delete(&foundChangePasswordCode)
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "sucess",
	})
}
