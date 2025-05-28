# Vue 3 Admin Dashboard

基于 Vue 3 + Vite + Element Plus 的现代化后台管理系统前端，配合 Gin-GORM 后端使用。

## 特性

- 🚀 **现代化技术栈**: Vue 3 + Vite + Element Plus
- 📱 **响应式设计**: 支持桌面端和移动端
- 🎨 **主题切换**: 支持明暗主题切换
- 🔐 **权限管理**: 基于角色的权限控制
- 📊 **数据可视化**: 仪表盘和图表展示
- 🛠 **开发体验**: 热重载、TypeScript 支持、ESLint + Prettier

## 技术栈

- **框架**: Vue 3.4 (Composition API)
- **构建工具**: Vite 5.0
- **UI 框架**: Element Plus 2.4
- **状态管理**: Pinia 2.1
- **路由**: Vue Router 4.2
- **HTTP 客户端**: Axios 1.6
- **样式**: SCSS
- **图标**: Element Plus Icons

## 项目结构

```
vue-admin/
├── public/                 # 静态资源
├── src/
│   ├── api/               # API 接口
│   │   ├── request.js     # Axios 配置
│   │   └── user.js        # 用户相关 API
│   ├── assets/            # 资源文件
│   │   └── styles/        # 样式文件
│   ├── components/        # 组件
│   │   └── layout/        # 布局组件
│   ├── router/            # 路由配置
│   ├── stores/            # Pinia 状态管理
│   ├── utils/             # 工具函数
│   ├── views/             # 页面组件
│   │   ├── auth/          # 认证页面
│   │   ├── dashboard/     # 仪表盘
│   │   ├── user/          # 用户相关
│   │   ├── admin/         # 管理员功能
│   │   └── error/         # 错误页面
│   ├── App.vue            # 根组件
│   └── main.js            # 入口文件
├── .env                   # 环境变量
├── vite.config.js         # Vite 配置
└── package.json           # 依赖配置
```

## 功能模块

### 认证系统
- 用户登录/注册
- JWT Token 管理
- 路由守卫
- 权限验证

### 用户管理
- 个人资料编辑
- 密码修改
- 头像上传

### 管理员功能
- 用户列表管理
- 用户状态控制
- 批量操作

### 仪表盘
- 数据统计展示
- 系统状态监控
- 快捷操作入口

## 开发指南

### 环境要求

- Node.js >= 16.0.0
- npm >= 8.0.0

### 安装依赖

```bash
npm install
```

### 开发模式

```bash
npm run dev
```

访问 http://localhost:12000

### 构建生产版本

```bash
npm run build
```

### 预览生产版本

```bash
npm run preview
```

## 配置说明

### 环境变量

- `.env` - 通用配置
- `.env.development` - 开发环境配置
- `.env.production` - 生产环境配置

### API 配置

默认 API 地址为 `http://localhost:8080/api`，可通过环境变量 `VITE_API_BASE_URL` 修改。

### 代理配置

开发环境下，Vite 会自动代理 `/api` 请求到后端服务器。

## 与后端集成

本前端项目设计用于与 Gin-GORM 后端配合使用，支持以下 API 接口：

### 认证接口
- `POST /api/auth/login` - 用户登录
- `POST /api/auth/register` - 用户注册
- `GET /api/auth/profile` - 获取用户信息

### 用户接口
- `PUT /api/user/profile` - 更新用户信息
- `PUT /api/user/password` - 修改密码

### 管理员接口
- `GET /api/admin/users` - 获取用户列表
- `POST /api/admin/users` - 创建用户
- `PUT /api/admin/users/:id` - 更新用户
- `DELETE /api/admin/users/:id` - 删除用户

## 部署

### Docker 部署

```dockerfile
FROM node:18-alpine as builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

### Nginx 配置

```nginx
server {
    listen 80;
    server_name localhost;
    root /usr/share/nginx/html;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    location /api {
        proxy_pass http://backend:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## 开发规范

### 代码风格

项目使用 ESLint + Prettier 进行代码格式化：

```bash
npm run lint    # 检查代码
npm run format  # 格式化代码
```

### 组件规范

- 使用 Composition API
- 组件名使用 PascalCase
- 文件名使用 kebab-case
- 样式使用 scoped

### 提交规范

使用 Conventional Commits 规范：

- `feat:` 新功能
- `fix:` 修复
- `docs:` 文档
- `style:` 样式
- `refactor:` 重构
- `test:` 测试
- `chore:` 构建/工具

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！