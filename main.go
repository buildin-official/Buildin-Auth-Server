package main

import (
	"log"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"pentag.kr/BuildinAuth/configs"
	"pentag.kr/BuildinAuth/database"
	"pentag.kr/BuildinAuth/routers"
)

func main() {
	runtime.GOMAXPROCS(configs.Config.WAS.PROCESS_NUM)

	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	//CORS Setting
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://buildin.kr",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	database.ConnectDB()

	routers.Initialize(app)
	log.Fatal(app.Listen(":" + configs.Config.WAS.Port))
}
