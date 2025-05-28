.PHONY: build run clean test deps docker-build docker-run

# 应用名称
APP_NAME := gin-grom
VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date +%Y-%m-%d\ %H:%M:%S)
GO_VERSION := $(shell go version | awk '{print $$3}')

# 构建标志
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GoVersion=$(GO_VERSION)"

# 默认目标
all: build

# 安装依赖
deps:
	go mod download
	go mod tidy

# 构建应用
build: deps
	go build $(LDFLAGS) -o bin/$(APP_NAME) cmd/server/main.go

# 运行应用
run: build
	./bin/$(APP_NAME) -config configs/config.yaml

# 开发模式运行
dev:
	go run cmd/server/main.go -config configs/config.yaml

# 清理构建文件
clean:
	rm -rf bin/
	rm -rf logs/

# 运行测试
test:
	go test -v ./...

# 运行测试并生成覆盖率报告
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# 格式化代码
fmt:
	go fmt ./...

# 代码检查
lint:
	golangci-lint run

# 构建 Docker 镜像
docker-build:
	docker build -t $(APP_NAME):$(VERSION) .
	docker tag $(APP_NAME):$(VERSION) $(APP_NAME):latest

# 运行 Docker 容器
docker-run:
	docker run -p 8080:8080 --name $(APP_NAME) $(APP_NAME):latest

# 停止并删除 Docker 容器
docker-stop:
	docker stop $(APP_NAME) || true
	docker rm $(APP_NAME) || true

# 生成 API 文档
docs:
	swag init -g cmd/server/main.go -o docs/

# 数据库迁移
migrate:
	go run cmd/server/main.go -config configs/config.yaml migrate

# 创建新的迁移文件
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations $$name

# 帮助信息
help:
	@echo "Available commands:"
	@echo "  deps           - Download dependencies"
	@echo "  build          - Build the application"
	@echo "  run            - Build and run the application"
	@echo "  dev            - Run in development mode"
	@echo "  clean          - Clean build files"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  fmt            - Format code"
	@echo "  lint           - Run linter"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  docker-stop    - Stop and remove Docker container"
	@echo "  docs           - Generate API documentation"
	@echo "  help           - Show this help message"