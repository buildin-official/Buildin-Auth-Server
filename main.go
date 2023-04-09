package main

import (
	"log"
	"runtime"

	"github.com/gofiber/fiber/v2"

	"pentag.kr/BuildinAuth/configs"
	"pentag.kr/BuildinAuth/database"
	"pentag.kr/BuildinAuth/routers"
)

func main() {
	runtime.GOMAXPROCS(configs.Config.WAS.PROCESS_NUM)

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	database.ConnectDB()
	database.ConnectRedis()

	routers.Initialize(app)
	log.Fatal(app.Listen(":" + configs.Config.WAS.Port))
}
