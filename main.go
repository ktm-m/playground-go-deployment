package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/ktm-m/playground-go-deployment/internal/config"
	"github.com/ktm-m/playground-go-deployment/internal/handler"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	app := fiber.New()
	appConfig := config.GetConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", appConfig.Redis.Host, appConfig.Redis.Port),
		Password: appConfig.Redis.Password,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic("[Main] cannot connect to redis: " + err.Error())
	}
	log.Println("[Main] connect to redis successfully")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok", appConfig.DB.Host, appConfig.DB.Port, appConfig.DB.User, appConfig.DB.Password, appConfig.DB.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("[Main] cannot connect to database: " + err.Error())
	}
	DB = db
	log.Println("[Main] connect to database successfully")

	handlers := handler.NewHandler(app)
	handlers.RegisterRoutes()

	if err = app.Listen(fmt.Sprintf(":%s", appConfig.HTTP.Port)); err != nil {
		return
	}
}
