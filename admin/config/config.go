package config

import (
	"os"
	"time"
)

// Config 包含应用程序的所有配置
type Config struct {
	// 服务器配置
	Port    string
	Env     string
	GinMode string

	// 数据库配置
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBCharset  string
	DBDSN      string

	// JWT配置
	JWTSecret    string
	JWTExpiresIn time.Duration
}

// LoadConfig 从环境变量加载配置
func LoadConfig() *Config {
	cfg := &Config{
		// 服务器配置
		Port:    getEnv("PORT", "8080"),
		Env:     getEnv("ENV", "development"),
		GinMode: getEnv("GIN_MODE", "debug"),

		// 数据库配置
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "crabshop"),
		DBCharset:  getEnv("DB_CHARSET", "utf8mb4"),

		// JWT配置
		JWTSecret: getEnv("JWT_SECRET", "default-secret-key-change-in-production"),
	}

	// 构建数据库DSN
	cfg.DBDSN = buildDSN(cfg)

	// 解析JWT过期时间
	expStr := getEnv("JWT_EXPIRATION", "24h")
	jwtExp, err := time.ParseDuration(expStr)
	if err != nil {
		jwtExp = 24 * time.Hour // 默认24小时
	}
	cfg.JWTExpiresIn = jwtExp

	return cfg
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// buildDSN 构建数据库连接字符串
func buildDSN(cfg *Config) string {
	return cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?charset=" + cfg.DBCharset + "&parseTime=True&loc=Local"
} 