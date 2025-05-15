package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/crabshop/admin/config"
	"github.com/golang-migrate/migrate/v4"
	mysqldriver "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// RunMigrations 使用golang-migrate运行数据库迁移
func RunMigrations(cfg *config.Config, direction string) error {
	// 获取原始数据库连接
	sqlDB, err := GetSQLDB(cfg)
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// 准备迁移驱动
	driver, err := mysqldriver.WithInstance(sqlDB, &mysqldriver.Config{})
	if err != nil {
		return fmt.Errorf("failed to create mysql driver: %w", err)
	}

	// 创建迁移实例
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // 迁移文件的位置
		cfg.DBName,          // 数据库名称
		driver,              // 数据库驱动
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// 根据方向执行迁移
	if direction == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to apply migrations: %w", err)
		}
		log.Println("Migrations applied successfully")
	} else if direction == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to rollback migrations: %w", err)
		}
		log.Println("Migrations rolled back successfully")
	} else {
		return fmt.Errorf("invalid migration direction: %s", direction)
	}

	return nil
}

// GetSQLDB 返回底层的sql.DB连接
func GetSQLDB(cfg *config.Config) (*sql.DB, error) {
	// 连接到数据库
	db, err := gorm.Open(mysql.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 获取底层的sql.DB连接
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	return sqlDB, nil
} 