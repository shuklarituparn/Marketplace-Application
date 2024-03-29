package main

import (
	"fmt"
	"github.com/shuklarituparn/marketplace/internal/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error opening the env files %v", err)
	}

	//gin.SetMode(gin.ReleaseMode)
	//gin.DisableConsoleColor()
	//
	//logs := logger.InitLog()
	//defer func(logs *os.File) {
	//	err := logs.Close()
	//	if err != nil {
	//
	//	}
	//}(logs)
	PORT := os.Getenv("SERVER_PORT")
	//gin.DefaultWriter = io.MultiWriter(logs)

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	fmt.Println(os.Getwd())
	router.Static("/api/static", "./static")
	router.Static("/api/uploads", "./uploads")
	router.NoRoute(handlers.NoRouteHandler)
	router.GET("/api/v1/", handlers.WelcomeHandler)
	router.GET("/api/v1/login", handlers.LoginHandler)
	router.POST("/api/v1/users/login", handlers.ProcessLoginHandler)
	router.GET("/api/v1/register", handlers.RegistrationHandler)
	router.POST("/api/v1/users/register", handlers.ProcessRegistrationHandler)

	err := router.Run(":" + PORT)
	if err != nil {
		return
	}
}
