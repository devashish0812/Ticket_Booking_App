package handlers

import (
	"fmt"
	"net/http"
	"time"

	// "github.com/devashish0812/user-service/services"
	"github.com/devashish0812/user-service/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	SecretKey   string
	authService services.AuthService
}

func NewAuthMiddleware(secretKey string, authService services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{SecretKey: secretKey,
		authService: authService}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Extract token from header
		tokenString, err := c.Cookie("refresh_token") // for /auth/refresh
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token cookie"})
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
		exp, ok := claims["exp"].(float64)
		if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
			c.JSON(401, gin.H{"error": "token expired"})
			c.Abort()
			return
		}

		userId, ok := claims["userid"].(string)
		if !ok || userId == "" {
			c.JSON(401, gin.H{"error": "invalid userId"})
			c.Abort()
			return
		}

		sessionId, ok := claims["sessionid"].(string)
		if !ok || sessionId == "" {
			c.JSON(401, gin.H{"error": "invalid sessionId"})
			c.Abort()
			return
		}
		c.Set("userid", userId)

		ctx := c.Request.Context()
		name, role, err := m.authService.ValidateToken(ctx, userId, tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("name", name)
		c.Set("role", role)
		c.Next()

	}
}
