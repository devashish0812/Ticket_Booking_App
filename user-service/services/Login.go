package services

import (
	"context"
	"time"

	"github.com/devashish0812/user-service/config"
	"github.com/devashish0812/user-service/models"

	//	"github.com/devashish0812/user-service/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoginService interface {
	LoginUser(ctx context.Context, user models.User, authService AuthService) (string, string, error)
}

type loginService struct {
	con *config.MongoConfig
}

func NewLoginService(con *config.MongoConfig) LoginService {
	return &loginService{con: con}
}

func (s *loginService) LoginUser(ctx context.Context, user models.User, authService AuthService) (string, string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	filter := bson.M{"name": user.Name, "password": user.Password}
	err := s.con.UserCol.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			println("No document found matching the filter.")
		} else {
			println("Error finding document:", err)
		}
		return "", "", err
	}
	accessToken, refreshToken, errToken := authService.GenerateToken(user)
	if errToken != nil {
		return "", "", errToken
	}

	// Prepare the update document
	update := bson.M{
		"$set": bson.M{
			"userId":       user.Userid.Hex(),
			"name":         user.Name,
			"role":         user.Role,
			"refreshToken": refreshToken,
			"createdAt":    time.Now(),
		},
	}

	filterToken := bson.M{"userId": user.Userid}
	opts := options.Update().SetUpsert(true)
	_, errUpsert := s.con.TokenCol.UpdateOne(ctx, filterToken, update, opts)
	if errUpsert != nil {
		return "", "", errUpsert
	}

	return accessToken, refreshToken, nil
}
