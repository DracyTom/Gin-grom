package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dengxing/gin-grom/internal/config"
	"github.com/dengxing/gin-grom/internal/database"
	"github.com/dengxing/gin-grom/internal/handler"
	"github.com/dengxing/gin-grom/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 解析命令行参数
	var configPath string
	flag.StringVar(&configPath, "config", "configs/config.yaml", "config file path")
	flag.Parse()

	// 加载配置
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志
	if err := utils.InitLogger(&cfg.Log); err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}

	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化数据库
	if err := database.InitDB(&cfg.Database); err != nil {
		logrus.Fatalf("Failed to init database: %v", err)
	}

	// 自动迁移数据库
	if err := database.AutoMigrate(); err != nil {
		logrus.Fatalf("Failed to migrate database: %v", err)
	}

	// 设置路由
	router := handler.SetupRouter()

	// 启动服务器
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	logrus.Infof("Server starting on %s", addr)

	// 优雅关闭
	go func() {
		if err := router.Run(addr); err != nil {
			logrus.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("Shutting down server...")

	// 关闭数据库连接
	if err := database.CloseDB(); err != nil {
		logrus.Errorf("Failed to close database: %v", err)
	}

	logrus.Info("Server exited")
}