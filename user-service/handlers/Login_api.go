package handlers

import (
	"net/http"
	"time"

	"github.com/devashish0812/user-service/models"
	"github.com/devashish0812/user-service/services"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	service     services.LoginService
	authservice services.AuthService
}

func NewLoginHandler(service services.LoginService, authservice services.AuthService) *LoginHandler {
	return &LoginHandler{service: service, authservice: authservice}
}

func (h *LoginHandler) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.service.LoginUser(c.Request.Context(), user, h.authservice)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
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

	// set refresh token cookie (long expiry)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		Domain:   "",
		MaxAge:   int((7 * 24 * time.Hour).Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	// respond with success message only
	c.JSON(http.StatusOK, gin.H{
		"message": "logged in successfully",
	})
}
