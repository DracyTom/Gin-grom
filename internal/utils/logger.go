package utils

import (
	"os"

	"github.com/dengxing/gin-grom/internal/config"
	"github.com/sirupsen/logrus"
)

func InitLogger(cfg *config.LogConfig) error {
	// 设置日志级别
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	// 设置日志格式
	if cfg.Format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	// 设置输出
	if cfg.Output == "file" && cfg.Filename != "" {
		file, err := os.OpenFile(cfg.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		logrus.SetOutput(file)
	} else {
		logrus.SetOutput(os.Stdout)
	}

	return nil
}