package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Username  string    `gorm:"size:50;not null;uniqueIndex" json:"username"`
	Password  string    `gorm:"size:100;not null" json:"password,omitempty"` // 仅在输出JSON时忽略
	Name      string    `gorm:"size:100" json:"name"`
	Email     string    `gorm:"size:100;index" json:"email"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Avatar    string    `gorm:"size:255" json:"avatar"`
	Role      string    `gorm:"size:20;default:user" json:"role"`
	LastLoginAt *time.Time `gorm:"type:datetime" json:"last_login_at"`
	LastLoginIP string    `gorm:"size:45" json:"last_login_ip"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}

// BeforeSave GORM 钩子，在保存前自动加密密码
func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password == "" {
		return nil
	}

	// 检查密码是否已经加密
	if !isPasswordHashed(u.Password) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// ValidatePassword 验证密码
func (u *User) ValidatePassword(password string) error {
	if u.Password == "" {
		return errors.New("密码不能为空")
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("密码不正确")
		}
		return err
	}

	return nil
}

// isPasswordHashed 检查密码是否已经加密
func isPasswordHashed(password string) bool {
	// bcrypt 加密的密码总是以 $2a$、$2b$ 或 $2y$ 开头
	return len(password) == 60 && (password[0:4] == "$2a$" || password[0:4] == "$2b$" || password[0:4] == "$2y$")
} 