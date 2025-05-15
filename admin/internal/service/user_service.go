package service

import (
	"errors"
	"net"
	"time"

	"github.com/crabshop/admin/config"
	"github.com/crabshop/admin/internal/auth"
	"github.com/crabshop/admin/internal/models"
	"gorm.io/gorm"
)

// UserService 用户服务
type UserService struct {
	db  *gorm.DB
	cfg *config.Config
}

// NewUserService 创建一个新的用户服务
func NewUserService(db *gorm.DB, cfg *config.Config) *UserService {
	return &UserService{
		db:  db,
		cfg: cfg,
	}
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

// Login 用户登录
func (s *UserService) Login(req *LoginRequest, ip string) (*LoginResponse, error) {
	var user models.User

	// 查找用户
	result := s.db.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, result.Error
	}

	// 验证密码
	if err := user.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	// 生成Token
	token, err := auth.GenerateToken(&user, s.cfg)
	if err != nil {
		return nil, err
	}

	// 更新登录信息
	now := time.Now()
	user.LastLoginAt = &now
	user.LastLoginIP = ip
	s.db.Save(&user)

	return &LoginResponse{
		Token: token,
		User:  &user,
	}, nil
}

// Register 用户注册
func (s *UserService) Register(user *models.User) (*models.User, error) {
	// 确保密码不为空
	if user.Password == "" {
		return nil, errors.New("密码不能为空")
	}

	// 检查用户名是否已存在
	var existingUser models.User
	result := s.db.Where("username = ?", user.Username).First(&existingUser)
	if result.Error == nil {
		return nil, errors.New("用户名已存在")
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	// 设置默认角色
	if user.Role == "" {
		user.Role = "user"
	}

	// 保存用户
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	// 验证用户是否成功保存并返回
	var savedUser models.User
	if err := s.db.Where("username = ?", user.Username).First(&savedUser).Error; err != nil {
		return nil, errors.New("用户创建后无法检索")
	}

	return &savedUser, nil
}

// GetUserByID 通过ID获取用户
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, result.Error
	}

	return &user, nil
}

// ParseIP 解析IP地址
func ParseIP(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return ""
	}
	return ip.String()
} 