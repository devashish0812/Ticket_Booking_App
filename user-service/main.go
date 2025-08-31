package main

import (
	"fmt"
	"log"
	"time"

	"github.com/devashish0812/user-service/config"
	"github.com/devashish0812/user-service/handlers"
	"github.com/devashish0812/user-service/services"
	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // React dev server
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	users := r.Group("/users")
	{
		users.POST("/signup", userHandler.Signup)
		users.POST("/login", loginHandler.Login)
	}
	r.POST("/auth/refresh", authMiddleware.RequireAuth(), authHandler.GetRefreshToken)
	// 4) Start
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
