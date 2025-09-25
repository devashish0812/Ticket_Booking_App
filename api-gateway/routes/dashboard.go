package routes

import (
	"api-gateway/config"
	"api-gateway/proxy"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterDashboardRoutes(r *gin.Engine, cfg config.ServiceConfig) {
	r.GET("/events", func(c *gin.Context) {
		token, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(401, gin.H{"error": "no token found"})
			return
		}

		// Split the token into parts (header.payload.signature)
		parts := strings.Split(token, ".")
		if len(parts) < 2 {
			c.JSON(400, gin.H{"error": "invalid token"})
			return
		}

		// Decode payload (base64url)
		payload, err := base64.RawURLEncoding.DecodeString(parts[1])
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid payload"})
			return
		}

		// Unmarshal JSON into map
		var claims map[string]interface{}
		if err := json.Unmarshal(payload, &claims); err != nil {
			c.JSON(400, gin.H{"error": "invalid claims"})
			return
		}

		// Check role
		if claims["role"] != "organizer" && claims["role"] != "user" {
			c.JSON(403, gin.H{"error": "forbidden"})
			return
		}

		switch claims["role"] {
		case "user":
			proxy.Forward(cfg.EventService, "/events/getall")(c)
		case "organizer":
			proxy.Forward(cfg.EventService, "/events/getallForOrg")(c)
		}

	})

}
