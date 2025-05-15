package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/crabshop/admin/config"
	"github.com/crabshop/admin/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims 自定义JWT声明结构
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken 为用户生成JWT Token
func GenerateToken(user *models.User, cfg *config.Config) (string, error) {
	// 设置token过期时间
	expirationTime := time.Now().Add(cfg.JWTExpiresIn)

	// 创建声明
	claims := &JWTClaims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "crabshop-admin",
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	// 使用签名方法和密钥创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken 验证JWT Token
func ValidateToken(tokenString string, cfg *config.Config) (*JWTClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("无法解析token: %w", err)
	}

	// 验证token是否有效
	if !token.Valid {
		return nil, errors.New("无效的token")
	}

	// 获取token中的声明
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("无法获取token声明")
	}

	return claims, nil
}