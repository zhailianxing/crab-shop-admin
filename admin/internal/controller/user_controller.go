package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/crabshop/admin/internal/models"
	"github.com/crabshop/admin/internal/service"
	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	userService *service.UserService
}

// NewUserController 创建一个新的用户控制器
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Login 处理用户登录请求
func (ctrl *UserController) Login(c *gin.Context) {
	var req service.LoginRequest

	// 验证请求数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	// 获取客户端IP
	clientIP := c.ClientIP()

	// 调用服务登录
	resp, err := ctrl.userService.Login(&req, clientIP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 返回登录成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "登录成功",
		"data": gin.H{
			"token": resp.Token,
			"user":  resp.User,
		},
	})
}

// Register 处理用户注册请求
func (ctrl *UserController) Register(c *gin.Context) {
	// 读取原始请求体
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("读取请求体失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "无效的请求体", "code": 1})
		return
	}
	
	// 重置请求体，以便后续处理
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	
	log.Printf("Register请求体原始内容: %s", string(bodyBytes))

	var user models.User
	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		log.Printf("解析请求体失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "无效的请求参数", "code": 2})
		return
	}

	log.Printf("解析后的用户对象: %+v", user)

	// 调用服务注册
	newUser, err := ctrl.userService.Register(&user)
	if err != nil {
		log.Printf("注册失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "code": 3})
		return
	}

	// 返回注册成功响应
	c.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"message": "注册成功",
		"data": gin.H{
			"user": newUser,
		},
	})
}

// GetProfile 获取当前用户信息
func (ctrl *UserController) GetProfile(c *gin.Context) {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "未认证", "code": 4})
		return
	}

	// 获取用户信息
	user, err := ctrl.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error(), "code": 5})
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"user": user,
		},
	})
}

// Logout 处理用户退出登录
func (c *UserController) Logout(ctx *gin.Context) {
	// 简单实现，仅返回成功消息
	// 前端收到成功响应后，应当清除本地存储的token
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "退出登录成功",
	})
}

// ModifyPassword 修改用户密码
func (c *UserController) ModifyPassword(ctx *gin.Context) {
	// 获取当前用户ID
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		return
	}
	
	// 解析请求参数
	var req struct {
		OldPwd   string `json:"oldPwd" binding:"required"`
		NewPwd   string `json:"newPwd" binding:"required,min=6"`
		ReNewPwd string `json:"reNewPwd" binding:"required,eqfield=NewPwd"`
	}
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}
	
	// 调用服务层修改密码
	err := c.userService.ModifyPassword(ctx, userID.(uint), req.OldPwd, req.NewPwd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "密码修改成功",
	})
} 