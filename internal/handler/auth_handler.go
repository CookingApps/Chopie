package handler

import (
	// "Chopie/internal/service"
	"Chopie/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthHandler struct {
	handler service.AuthService
}

func (h *AuthHandler) CreateUser(c *gin.Context) {
	var userrequest CreateUserRequest

	// Handler logic goes here
	err := c.BindJSON(&userrequest)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "invalid request"})
		return
	}

	err = h.handler.Register(userrequest.Email, userrequest.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "Error creating user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "received", "email": userrequest.Email, "password": userrequest.Password})
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{handler: s}
}
