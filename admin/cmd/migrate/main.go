package main

import (
	"flag"
	"log"

	"github.com/crabshop/admin/config"
	"github.com/crabshop/admin/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// 解析命令行参数
	direction := flag.String("direction", "up", "Migration direction (up or down)")
	flag.Parse()

	// 验证方向
	if *direction != "up" && *direction != "down" {
		log.Fatalf("Invalid direction: %s. Must be 'up' or 'down'", *direction)
	}

	// 初始化配置
	cfg := config.LoadConfig()

	// 运行迁移
	if err := database.RunMigrations(cfg, *direction); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Printf("Migration '%s' completed successfully", *direction)
} 