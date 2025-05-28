# Gin-GORM 脚手架

一个基于 Gin 和 GORM 的 Go Web 应用脚手架，提供了完整的用户认证、权限管理和 RESTful API 功能。

## 特性

- 🚀 基于 Gin 的高性能 Web 框架
- 🗄️ GORM ORM 支持 MySQL 和 PostgreSQL
- 🔐 JWT 认证和授权
- 🛡️ 中间件支持（CORS、日志、恢复等）
- 📝 结构化日志记录
- ⚙️ 配置管理（支持环境变量）
- 🐳 Docker 和 Docker Compose 支持
- 📊 数据库迁移
- 🧪 单元测试支持

## 项目结构

```
.
├── cmd/
│   └── server/          # 应用程序入口
├── internal/
│   ├── config/          # 配置管理
│   ├── database/        # 数据库连接
│   ├── handler/         # HTTP 处理器
│   ├── middleware/      # 中间件
│   ├── model/           # 数据模型
│   ├── service/         # 业务逻辑
│   └── utils/           # 工具函数
├── pkg/
│   ├── auth/            # 认证相关
│   └── response/        # 响应格式
├── configs/             # 配置文件
├── docs/                # 文档
├── scripts/             # 脚本文件
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

## 快速开始

### 环境要求

- Go 1.21+
- MySQL 8.0+ 或 PostgreSQL 12+
- Docker (可选)

### 本地开发

1. 克隆项目
```bash
git clone https://github.com/dengxing/gin-grom.git
cd gin-grom
```

2. 安装依赖
```bash
make deps
```

3. 配置数据库
编辑 `configs/config.yaml` 文件，配置数据库连接信息。

4. 运行应用
```bash
make dev
```

应用将在 `http://localhost:8080` 启动。

### 使用 Docker

1. 使用 Docker Compose 启动所有服务
```bash
docker-compose up -d
```

这将启动：
- Web 应用 (端口 8080)
- MySQL 数据库 (端口 3306)
- Redis (端口 6379)
- Adminer 数据库管理工具 (端口 8081)

## API 文档

### 认证相关

#### 用户注册
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}
```

#### 用户登录
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

### 用户相关

#### 获取个人信息
```http
GET /api/v1/user/profile
Authorization: Bearer <token>
```

#### 更新个人信息
```http
PUT /api/v1/user/profile
Authorization: Bearer <token>
Content-Type: application/json

{
  "username": "newusername",
  "email": "newemail@example.com",
  "avatar": "http://example.com/avatar.jpg"
}
```

### 管理员相关

#### 获取用户列表
```http
GET /api/v1/admin/users?page=1&size=10
Authorization: Bearer <admin_token>
```

#### 获取指定用户
```http
GET /api/v1/admin/users/{id}
Authorization: Bearer <admin_token>
```

#### 更新用户
```http
PUT /api/v1/admin/users/{id}
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "username": "newusername",
  "email": "newemail@example.com",
  "status": 1,
  "role": "user"
}
```

#### 删除用户
```http
DELETE /api/v1/admin/users/{id}
Authorization: Bearer <admin_token>
```

## 配置

配置文件位于 `configs/config.yaml`，支持以下配置项：

- `server`: 服务器配置
- `database`: 数据库配置
- `jwt`: JWT 配置
- `log`: 日志配置
- `cors`: CORS 配置

也可以通过环境变量进行配置，环境变量前缀为 `GIN_GROM_`。

## 开发

### 可用命令

```bash
make deps           # 安装依赖
make build          # 构建应用
make run            # 构建并运行
make dev            # 开发模式运行
make test           # 运行测试
make test-coverage  # 运行测试并生成覆盖率报告
make fmt            # 格式化代码
make lint           # 代码检查
make clean          # 清理构建文件
make docker-build   # 构建 Docker 镜像
make docker-run     # 运行 Docker 容器
make help           # 显示帮助信息
```

### 添加新功能

1. 在 `internal/model/` 中定义数据模型
2. 在 `internal/service/` 中实现业务逻辑
3. 在 `internal/handler/` 中实现 HTTP 处理器
4. 在 `internal/handler/router.go` 中添加路由

## 部署

### 生产环境配置

1. 创建生产环境配置文件 `configs/config.prod.yaml`
2. 设置环境变量或修改配置文件
3. 构建并部署应用

```bash
make build
./bin/gin-grom -config configs/config.prod.yaml
```

### Docker 部署

```bash
docker build -t gin-grom .
docker run -p 8080:8080 gin-grom
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License