package configs

import (
	"os"
	"strconv"
)

type WASConfig struct { // 인증서버 환경변수 구조체
	Host        string
	Port        string
	JWT_SECRET  string
	JWT_EXPIRE  int
	PROCESS_NUM int // gorutine에서 사용할 프로세스 개수
}

type DBConfig struct { //Postgres 환경변수 구조체
	Host     string
	Port     int
	User     string
	Pass     string
	Database string
}

type RedisConfig struct { // Redis 환경변수 구조체
	Host     string
	Port     int
	Password string
	DB       int
}

type SMTPConfig struct {
	Host       string
	Port       int
	User       string
	Pass       string
	Sender     string // 보내는 사람의 이메일 주소
	SenderName string // 보내는 사람의 이름
}

type MainConfig struct { // 중앙 구조체
	WAS   WASConfig
	DB    DBConfig
	Redis RedisConfig
	SMTP  SMTPConfig
}

func getEnv_s(key string) string { // 환경변수를 가져오고 없으면 에러를 발생시키는 함수
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic("Environment variable " + key + " not found")
}

func getAllEnv() MainConfig { // 모든 환경변수를 가져오는 함수
	var err error
	jwtExpire, err := strconv.Atoi(getEnv_s("JWT_EXPIRE"))
	if err != nil {
		panic("Environment variable JWT_EXPIRE is not a number")
	}

	dbPort, err := strconv.Atoi(getEnv_s("POSTGRES_PORT"))

	if err != nil {
		panic("Environment variable DB_PORT is not a number")
	}
	redisPort, err := strconv.Atoi(getEnv_s("REDIS_PORT"))
	if err != nil {
		panic("Environment variable REDIS_PORT is not a number")
	}
	smtpPort, err := strconv.Atoi(getEnv_s("SMTP_PORT"))
	if err != nil {
		panic("Environment variable SMTP_PORT is not a number")
	}
	processNum, err := strconv.Atoi(getEnv_s("PROCESS_NUM"))
	if err != nil {
		panic("Environment variable PROCESS_NUM is not a number")
	}

	return MainConfig{
		WAS: WASConfig{
			Host:        getEnv_s("WAS_HOST"),
			Port:        getEnv_s("WAS_PORT"),
			JWT_SECRET:  getEnv_s("JWT_SECRET"),
			JWT_EXPIRE:  jwtExpire,
			PROCESS_NUM: processNum,
		},
		DB: DBConfig{
			Host:     getEnv_s("POSTGRES_HOST"),
			Port:     dbPort,
			User:     getEnv_s("POSTGRES_USER"),
			Pass:     getEnv_s("POSTGRES_PASSWORD"),
			Database: getEnv_s("POSTGRES_DB"),
		},
		Redis: RedisConfig{
			Host:     getEnv_s("REDIS_HOST"),
			Port:     redisPort,
			Password: getEnv_s("REDIS_PASSWORD"),
			DB:       0,
		},
		SMTP: SMTPConfig{
			Host:       getEnv_s("SMTP_HOST"),
			Port:       smtpPort,
			User:       getEnv_s("SMTP_USER"),
			Pass:       getEnv_s("SMTP_PASSWORD"),
			Sender:     getEnv_s("SMTP_SENDER"),
			SenderName: getEnv_s("SMTP_SENDER_NAME"),
		},
	}
}

var Config = getAllEnv() // 환경변수 객체
