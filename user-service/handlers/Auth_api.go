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

	c.SetCookie("access_token", accessToken, int((15 * time.Minute).Seconds()), "/", "localhost", true, true)

	// respond with success message only
	c.JSON(http.StatusOK, gin.H{
		"message": "token refreshed successfully",
	})
}
