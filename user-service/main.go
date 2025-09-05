package main

import (
	"fmt"
	"log"
	"os"

	"github.com/devashish0812/user-service/config"
	"github.com/devashish0812/user-service/handlers"
	"github.com/devashish0812/user-service/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1) Initialize Mongo via config (reads MONGO_URI from env)
	_ = godotenv.Load()
	mongoCfg := config.InitMongo()
	// 2) Wire layers
	userService := services.NewUserService(mongoCfg.UserCol) // signup service
	userHandler := handlers.NewUserHandler(userService)      // signup handler

	authService := services.NewAuthService(mongoCfg.JWTSecret, mongoCfg)
	loginService := services.NewLoginService(mongoCfg)
	loginHandler := handlers.NewLoginHandler(loginService, *authService)

	fmt.Println("After NewAuthService")
	println("After NewAuthService1")

	authMiddleware := handlers.NewAuthMiddleware(mongoCfg.JWTSecret, *authService)
	authHandler := handlers.NewAuthHandler(*authService)

	// 3) Routes
	r := gin.Default()

	users := r.Group("/users")
	{
		users.POST("/signup", userHandler.Signup)
		users.POST("/login", loginHandler.Login)
	}
	r.POST("/auth/refresh", authMiddleware.RequireAuth(), authHandler.GetRefreshToken)
	// 4) Start
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
