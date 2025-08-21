package main

import (
	"github.com/devashish0812/user-service/config"
	"github.com/devashish0812/user-service/handlers"
	"github.com/devashish0812/user-service/services"
	"github.com/joho/godotenv"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1) Initialize Mongo via config (reads MONGO_URI from env)
	_ = godotenv.Load()
	mongoCfg := config.InitMongo()

	// 2) Wire layers
	userService := services.NewUserService(mongoCfg.UserCol) // pass collection, not db
	userHandler := handlers.NewUserHandler(userService)

	// 3) Routes
	r := gin.Default()
	users := r.Group("/users")
	{
		users.POST("/signup", userHandler.Signup)
	}

	// 4) Start
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
