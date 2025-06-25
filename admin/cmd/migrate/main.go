package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/crabshop/admin/config"
	"github.com/crabshop/admin/internal/database"
	"github.com/joho/godotenv"
)

/*
# 使用本地环境配置
go run cmd/migrate/main.go -direction=up -env=local

# 使用开发环境配置
go run cmd/migrate/main.go -direction=up -env=dev

# 使用生产环境配置
go run cmd/migrate/main.go -direction=up -env=prod


select * from schema_migrations;
+---------+-------+
| version | dirty |
+---------+-------+
|       4 |     1 |
+---------+-------+
=> 版本号为4，dirty=1表示迁移失败（可能是sql文件有问题）

如果已经确认上一个版本3是正确的，则可以强制设置迁移版本：
UPDATE schema_migrations SET version = 3, dirty = 0;

接下来重新运行迁移：
go run cmd/migrate/main.go -direction=up -env=local

select * from schema_migrations;
+---------+-------+
| version | dirty |
+---------+-------+
|       4 |     0 |
*/
func main() {
	// 解析命令行参数
	direction := flag.String("direction", "up", "Migration direction (up or down)")
	env := flag.String("env", "", "环境配置文件 (dev, prod, local)，不指定则使用默认local.env")
	flag.Parse()

	// 验证方向
	if *direction != "up" && *direction != "down" {
		log.Fatalf("Invalid direction: %s. Must be 'up' or 'down'", *direction)
	}

	// 根据参数加载环境变量
	if err := loadEnvFile(env); err != nil {
		log.Fatalf("加载环境配置文件失败: %v", err)
	}

	// 初始化配置
	cfg := config.LoadConfig()

	// 运行迁移
	if err := database.RunMigrations(cfg, *direction); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Printf("Migration '%s' completed successfully", *direction)
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