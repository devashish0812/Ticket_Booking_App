package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket-service/services"
)

type SeatLockHandler struct {
	service *services.SeatLockService
}

func NewSeatLockHandler(service *services.SeatLockService) *SeatLockHandler {
	return &SeatLockHandler{service: service}
}

func (h *SeatLockHandler) HandleLockSeat(c *gin.Context) {
	var req struct {
		SeatID string `json:"seatId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	userID, ok := userIDInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	err := h.service.LockSeat(c.Request.Context(), req.SeatID, userID)

	if err != nil {
		if err.Error() == "seat_already_locked" {
			c.JSON(http.StatusConflict, gin.H{"error": "Seat is already booked"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seat locked successfully"})
}
