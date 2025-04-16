package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/services"
	
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service services.AuthServiceInterface
}

func NewAuthController(service services.AuthServiceInterface) *AuthController {
	return &AuthController{Service: service}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}
	if err := ctrl.Service.RegisterUser(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}
	user, err := ctrl.Service.LoginUser(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	utils.SetSessionCookie(c, user.ID.String())
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

func (ctrl *AuthController) Me(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	user, err := ctrl.Service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (ctrl *AuthController) ForgotPassword(c *gin.Context) {
	var req dto.ForgotPasswordRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	user, err := ctrl.Service.FindUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Email not found"})
		return
	}

	otp := utils.GenerateOTP(6)

	key := "forgot_pass:" + user.Email
	err = config.RedisClient.Set(config.Ctx, key, otp, 15*time.Minute).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save token"})
		return
	}

	subject := "Reset Your Password"
	body := fmt.Sprintf("Your reset code is: <b>%s</b><br/>It will expire in 15 minutes.", otp)
	if err := utils.SendEmail(user.Email, subject, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to your email"})
}

func (ctrl *AuthController) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	var email string
	keys := config.RedisClient.Keys(config.Ctx, "forgot_pass:*").Val()
	for _, key := range keys {
		val := config.RedisClient.Get(config.Ctx, key).Val()
		if val == req.Token {
			email = strings.TrimPrefix(key, "forgot_pass:")
			break
		}
	}

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid or expired token"})
		return
	}

	if err := ctrl.Service.ResetPassword(email, req.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to reset password"})
		return
	}
	_ = config.RedisClient.Del(config.Ctx, "forgot_pass:"+email).Err()

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

func (ctrl *AuthController) ChangePassword(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.ChangePasswordRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := ctrl.Service.ChangePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func (ctrl *AuthController) GetAllUsers(c *gin.Context) {
	search := c.Query("search")
	page := utils.GetQueryInt(c, "page", 1)
	limit := utils.GetQueryInt(c, "limit", 10)

	users, total, err := ctrl.Service.GetAllUsersPaginated(search, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"meta": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func (ctrl *AuthController) GetUserByIDAdmin(c *gin.Context) {
	id := c.Param("id")

	user, err := ctrl.Service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
