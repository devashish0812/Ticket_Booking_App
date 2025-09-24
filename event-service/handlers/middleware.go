package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	SecretKey string
}

func NewAuthMiddleware(secretKey string) *AuthMiddleware {
	return &AuthMiddleware{SecretKey: secretKey}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Extract token from header
		tokenString, err := c.Cookie("access_token")
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing access token cookie"})
			c.Abort()
			return
		}
		// 2. Validate token (signature + expiry)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(m.SecretKey), nil
			//The library parses the JWT and calls our callback to get the secret key;
			//it then uses that key to verify the signature and, if valid, returns the parsed token.
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		userId, ok := claims["userid"].(string)
		if !ok || userId == "" {
			c.JSON(401, gin.H{"error": "invalid userId"})
			c.Abort()
			return
		}
		c.Set("userid", userId)

		name, ok := claims["name"].(string)
		if !ok || name == "" {
			c.JSON(401, gin.H{"error": "invalid name"})
			c.Abort()
			return
		}
		c.Set("name", name)

		role, ok := claims["role"].(string)
		if !ok || role == "" {
			c.JSON(401, gin.H{"error": "invalid roll"})
			c.Abort()
			return
		}
		c.Set("role", role)
		c.Next()

	}
}
