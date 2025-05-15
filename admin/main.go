package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/crabshop/admin/config"
	"github.com/crabshop/admin/internal/database"
	"github.com/crabshop/admin/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 定义命令行参数
	env := flag.String("env", "", "环境配置文件 (dev, prod, local)，不指定则使用默认local.env")
	flag.Parse()

	// 根据参数加载环境变量
	if err := loadEnvFile(env); err != nil {
		log.Fatalf("加载环境配置文件失败: %v", err)
	}

	// 初始化配置
	cfg := config.LoadConfig()

	// 设置Gin模式
	gin.SetMode(cfg.GinMode)

	// 初始化数据库连接
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 初始化路由
	router := routes.SetupRouter(db)

	// 启动服务器
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting in %s mode on http://localhost%s", cfg.Env, serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// loadEnvFile 根据指定的环境加载对应的环境配置文件
func loadEnvFile(env *string) error {
	var envFile string

	// 确定要加载的环境文件
	switch *env {
	case "dev":
		envFile = "config/dev.env"
	case "prod":
		envFile = "config/prod.env"
	case "local":
		envFile = "config/local.env"
	default:
		envFile = "config/local.env" // 默认使用local.env
	}

	// 检查环境文件是否存在
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		return fmt.Errorf("环境配置文件 %s 不存在", envFile)
	}

	// 加载环境文件
	absPath, _ := filepath.Abs(envFile)
	log.Printf("Loading environment from: %s", absPath)
	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("加载环境配置文件 %s 失败: %w", envFile, err)
	}
	
	return nil
} 