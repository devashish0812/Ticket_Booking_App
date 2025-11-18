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
	LoginUser(ctx context.Context, user models.User, authService AuthService) (string, string, models.User, error)
}

type loginService struct {
	con *config.MongoConfig
}

func NewLoginService(con *config.MongoConfig) LoginService {
	return &loginService{con: con}
}

func (s *loginService) LoginUser(ctx context.Context, user models.User, authService AuthService) (string, string, models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	filter := bson.M{"name": user.Name, "password": user.Password}
	err := s.con.UserCol.FindOne(ctx, filter).Decode(&user)
	payload := models.User{}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			println("No document found matching the filter.")
		} else {
			println("Error finding document:", err)
		}
		return "", "", payload, err
	}
	accessToken, refreshToken, errToken := authService.GenerateToken(user)
	if errToken != nil {
		return "", "", payload, errToken
	}

	// Prepare the update document
	update := bson.M{
		"$set": bson.M{
			"userId":       user.Userid.Hex(),
			"name":         user.Name,
			"role":         user.Role,
			"refreshToken": refreshToken,
		},
	}

	filterToken := bson.M{"userId": user.Userid}
	opts := options.Update().SetUpsert(true)
	_, errUpsert := s.con.TokenCol.UpdateOne(ctx, filterToken, update, opts)
	if errUpsert != nil {
		return "", "", payload, errUpsert
	}

	payload.Role = user.Role
	payload.Name = user.Name
	return accessToken, refreshToken, payload, nil
}
