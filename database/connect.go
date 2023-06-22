package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"pentag.kr/BuildinAuth/configs"
	"pentag.kr/BuildinAuth/models"
)

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	dbConfig := configs.Config.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Seoul", dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.Database, dbConfig.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(
		&models.User{},
		&models.UnvalidatedUser{},
		&models.ChangePasswordCode{},
		&models.RefreshToken{},
	)
	if err != nil {
		panic(err)
	}
}
