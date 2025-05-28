package handler

import (
	"github.com/dengxing/gin-grom/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	// 中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "Server is running",
		})
	})

	// API 路由组
	api := r.Group("/api/v1")
	{
		// 用户相关路由
		userHandler := NewUserHandler()
		
		// 公开路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
		}

		// 需要认证的路由
		protected := api.Group("/")
		protected.Use(middleware.JWTAuth())
		{
			// 用户个人信息
			user := protected.Group("/user")
			{
				user.GET("/profile", userHandler.GetProfile)
				user.PUT("/profile", userHandler.UpdateProfile)
			}

			// 管理员路由
			admin := protected.Group("/admin")
			admin.Use(middleware.RequireRole("admin"))
			{
				users := admin.Group("/users")
				{
					users.GET("", userHandler.GetUsers)
					users.GET("/:id", userHandler.GetUser)
					users.PUT("/:id", userHandler.UpdateUser)
					users.DELETE("/:id", userHandler.DeleteUser)
				}
			}
		}
	}

	return r
}