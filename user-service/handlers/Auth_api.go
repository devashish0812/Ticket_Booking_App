package handlers

import (
	"net/http"
	"time"

	"github.com/devashish0812/user-service/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) GetRefreshToken(c *gin.Context) {

	accessToken, err := h.service.RefreshToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   int((15 * time.Minute).Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "token refreshed successfully",
	})
}
