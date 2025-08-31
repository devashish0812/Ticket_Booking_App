package services

import (
	"context"
	"time"

	"github.com/devashish0812/user-service/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error
}

type userService struct {
	col *mongo.Collection
}

func NewUserService(col *mongo.Collection) UserService {
	return &userService{col: col}
}

func (s *userService) CreateUser(ctx context.Context, user models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := s.col.InsertOne(ctx, user)
	return err
}
