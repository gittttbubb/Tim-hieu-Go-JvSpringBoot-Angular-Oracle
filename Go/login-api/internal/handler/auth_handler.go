package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"login-api/internal/dto/request"
	"login-api/internal/service"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	// dùng var vì LoginRequest là 1 kiểu dữ liệu(type) không phải là 1 giá trị(value), nên không thể dùng := để khai báo và khởi tạo biến req
	var req request.LoginRequest
	// Giải mã JSON bằng c.ShouldBindJSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	token, err := h.service.Login(req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req  request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	err := h.service.Register(req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "register success",
	})
}
// USER
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
    userID := c.GetInt("userId")
    var req request.UpdateProfileRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{
            "message": "invalid request",
        })
        return
    }
    err := h.service.UpdateProfile(userID, req.Username)
    if err != nil {
        c.JSON(500, gin.H{
            "message": err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "message": "updated",
    })
}

func (h *AuthHandler) ChangePassword(c *gin.Context) {
    userID := c.GetInt("userId")
    var req request.ChangePasswordRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{
            "message": "invalid request",
        })
        return
    }
    err := h.service.ChangePassword(userID, req.Password)
    if err != nil {
        c.JSON(500, gin.H{
            "message": err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "message": "password changed",
    })
}

// ADMIN
func (h *AuthHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func (h *AuthHandler) PromoteToAdmin(c *gin.Context) {
	id, err := strconv.Atoi(
		c.Param("id"),
	)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id",
		})
		return
	}
	err = h.service.PromoteToAdmin(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "user promoted to admin",
	})
}