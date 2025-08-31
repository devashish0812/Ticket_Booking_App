package services

import (
	"context"
	"fmt"
	"time"

	"github.com/devashish0812/user-service/config"
	"github.com/devashish0812/user-service/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
)

type Claims struct {
	Userid    string `bson:"userid" json:"userid"`
	Name      string `json:"name" bson:"name"`
	Role      string `json:"role" bson:"role"`
	SessionId string `json:"sessionid" bson:"sessionid"`
	jwt.RegisteredClaims
}

type AuthService struct {
	SecretKey string
	con       *config.MongoConfig
}

func NewAuthService(jwtSecret string, con *config.MongoConfig) *AuthService {
	fmt.Println("sign key:", jwtSecret)

	return &AuthService{SecretKey: jwtSecret,
		con: con}

}

func (a *AuthService) GenerateToken(user models.User) (string, string, error) {
	fmt.Println("sign key:", a.SecretKey)
	accessClaims := &Claims{
		Userid: user.Userid.Hex(),
		Name:   user.Name,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshClaims := &Claims{
		Userid:    user.Userid.Hex(),
		SessionId: uuid.New().String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", "", err
	}

	return accessStr, refreshStr, nil

}

func (a *AuthService) ValidateToken(ctx context.Context, userid string, refreshToken string) (string, string, error) {

	filter := bson.M{"userId": userid, "refreshToken": refreshToken}
	fmt.Println(userid, refreshToken)
	var result bson.M
	err := a.con.TokenCol.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return "", "", err
	}

	return result["name"].(string), result["role"].(string), nil

}

func (a *AuthService) RefreshToken(ctx *gin.Context) (string, error) {
	nameVal, exists := ctx.Get("name")
	if !exists {
		return "", fmt.Errorf("unauthorized")
	}
	fmt.Println(nameVal)
	roleVal, ok := ctx.Get("role")
	if !ok {
		return "", fmt.Errorf("unauthorized")
	}
	userid, ok := ctx.Get("userid")
	if !ok {
		return "", fmt.Errorf("unauthorized")
	}
	accessClaims := &Claims{
		Userid: userid.(string),
		Name:   nameVal.(string),
		Role:   roleVal.(string),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", err
	}

	return accessStr, nil

}
