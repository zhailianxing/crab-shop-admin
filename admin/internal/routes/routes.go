package routes

import (
	"time"

	"github.com/crabshop/admin/config"
	"github.com/crabshop/admin/internal/controller"
	"github.com/crabshop/admin/internal/middleware"
	"github.com/crabshop/admin/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter 设置所有路由
func SetupRouter(db *gorm.DB) *gin.Engine {
	// 初始化配置
	cfg := config.LoadConfig()

	// 创建路由
	router := gin.Default()

	// 配置CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 创建服务
	userService := service.NewUserService(db, cfg)

	// 创建控制器
	userController := controller.NewUserController(userService)

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API路由组
	api := router.Group("/api")
	{
		// 无需认证的路由
		api.POST("/user/login", userController.Login)
		api.POST("/user/register", userController.Register)

		// 需要认证的路由
		auth := api.Group("/") //Group()中的参数是路由前缀。 
		auth.Use(middleware.JWTAuth(cfg))
		{
			// 用户相关
			// 最终的路由是：/api/user/profile。 即 Group("/api") + Group("/") + GET("/user/profile")
			auth.GET("/user/profile", userController.GetProfile)

			// 管理员相关
			admin := auth.Group("/admin")
			admin.Use(middleware.RequireRole("admin"))
			{
				// TODO: 添加管理员路由
			}
		}
	}

	return router
} 