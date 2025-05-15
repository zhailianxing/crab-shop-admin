package database

import (
	"fmt"
	"log"

	"github.com/crabshop/admin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 初始化数据库连接
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	// 配置GORM
	gormConfig := &gorm.Config{}

	// 在开发环境中启用详细日志
	if cfg.Env == "development" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	// 连接到数据库
	db, err := gorm.Open(mysql.Open(cfg.DBDSN), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")
	return db, nil
} 