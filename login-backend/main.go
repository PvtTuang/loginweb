package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"

	"login-backend/handler"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("ไม่สามารถโหลดไฟล์ .env:", err)
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		log.Fatal("JWT_SECRET ไม่ได้ตั้งค่าในไฟล์ .env")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	handler.SetRedisClient(rdb, ctx)
	handler.SetJWTSecret(jwtSecret)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Login)

	api := r.Group("/api")
	api.Use(handler.AuthMiddleware())
	{
		api.POST("/logout", handler.Logout)
	}

	r.Run(":8080")

}
