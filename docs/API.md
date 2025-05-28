# API 文档

## 基础信息

- **Base URL**: `http://localhost:8080`
- **认证方式**: Bearer Token (JWT)
- **Content-Type**: `application/json`

## 响应格式

所有 API 响应都遵循统一格式：

```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

### 状态码说明

- `200`: 成功
- `400`: 请求参数错误
- `401`: 未认证或认证失败
- `403`: 权限不足
- `404`: 资源不存在
- `500`: 服务器内部错误

## 公开接口

### 健康检查

**GET** `/health`

检查服务器状态。

**响应示例**:
```json
{
  "message": "Server is running",
  "status": "ok"
}
```

### 用户注册

**POST** `/api/v1/auth/register`

注册新用户。

**请求参数**:
```json
{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123",
  "role": "user"  // 可选，默认为 "user"，可设置为 "admin"
}
```

**参数说明**:
- `username`: 用户名，3-50字符，必填
- `email`: 邮箱地址，必填
- `password`: 密码，最少6字符，必填
- `role`: 角色，可选值：`user`、`admin`

**响应示例**:
```json
{
  "code": 200,
  "message": "User registered successfully",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "avatar": "",
    "status": 1,
    "role": "user",
    "created_at": "2025-05-28T06:25:31.437Z",
    "updated_at": "2025-05-28T06:25:31.437Z"
  }
}
```

### 用户登录

**POST** `/api/v1/auth/login`

用户登录获取 JWT Token。

**请求参数**:
```json
{
  "username": "testuser",
  "password": "password123"
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "avatar": "",
      "status": 1,
      "role": "user",
      "created_at": "2025-05-28T06:25:31.437Z",
      "updated_at": "2025-05-28T06:25:31.437Z"
    }
  }
}
```

## 用户接口（需要认证）

所有用户接口都需要在请求头中包含 JWT Token：

```
Authorization: Bearer <your-jwt-token>
```

### 获取个人信息

**GET** `/api/v1/user/profile`

获取当前登录用户的个人信息。

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "avatar": "",
    "status": 1,
    "role": "user",
    "created_at": "2025-05-28T06:25:31.437Z",
    "updated_at": "2025-05-28T06:25:31.437Z"
  }
}
```

### 更新个人信息

**PUT** `/api/v1/user/profile`

更新当前登录用户的个人信息。

**请求参数**:
```json
{
  "username": "newusername",  // 可选
  "email": "newemail@example.com",  // 可选
  "avatar": "https://example.com/avatar.jpg"  // 可选
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "Profile updated successfully",
  "data": {
    "id": 1,
    "username": "newusername",
    "email": "newemail@example.com",
    "avatar": "https://example.com/avatar.jpg",
    "status": 1,
    "role": "user",
    "created_at": "2025-05-28T06:25:31.437Z",
    "updated_at": "2025-05-28T06:26:10.972Z"
  }
}
```

## 管理员接口（需要管理员权限）

### 获取用户列表

**GET** `/api/v1/admin/users`

获取所有用户列表（分页）。

**查询参数**:
- `page`: 页码，默认为 1
- `size`: 每页数量，默认为 10

**请求示例**:
```
GET /api/v1/admin/users?page=1&size=10
```

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "avatar": "",
      "status": 1,
      "role": "user",
      "created_at": "2025-05-28T06:25:31.437Z",
      "updated_at": "2025-05-28T06:25:31.437Z"
    }
  ],
  "total": 5,
  "page": 1,
  "size": 10
}
```

### 获取指定用户信息

**GET** `/api/v1/admin/users/:id`

获取指定用户的详细信息。

**路径参数**:
- `id`: 用户ID

**响应示例**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "avatar": "",
    "status": 1,
    "role": "user",
    "created_at": "2025-05-28T06:25:31.437Z",
    "updated_at": "2025-05-28T06:25:31.437Z"
  }
}
```

### 更新指定用户信息

**PUT** `/api/v1/admin/users/:id`

更新指定用户的信息。

**路径参数**:
- `id`: 用户ID

**请求参数**:
```json
{
  "username": "newusername",  // 可选
  "email": "newemail@example.com",  // 可选
  "avatar": "https://example.com/avatar.jpg",  // 可选
  "status": 1,  // 可选，0=禁用，1=启用
  "role": "admin"  // 可选，user 或 admin
}
```

**响应示例**:
```json
{
  "code": 200,
  "message": "User updated successfully",
  "data": {
    "id": 1,
    "username": "newusername",
    "email": "newemail@example.com",
    "avatar": "https://example.com/avatar.jpg",
    "status": 1,
    "role": "admin",
    "created_at": "2025-05-28T06:25:31.437Z",
    "updated_at": "2025-05-28T06:26:10.972Z"
  }
}
```

### 删除指定用户

**DELETE** `/api/v1/admin/users/:id`

删除指定用户。

**路径参数**:
- `id`: 用户ID

**响应示例**:
```json
{
  "code": 200,
  "message": "User deleted successfully"
}
```

## 错误响应

### 参数验证错误

```json
{
  "code": 400,
  "message": "username already exists"
}
```

### 认证失败

```json
{
  "code": 401,
  "message": "invalid token"
}
```

### 权限不足

```json
{
  "code": 403,
  "message": "forbidden"
}
```

### 资源不存在

```json
{
  "code": 404,
  "message": "user not found"
}
```

### 服务器错误

```json
{
  "code": 500,
  "message": "internal server error"
}
```

## 使用示例

### 完整的用户注册和登录流程

```bash
# 1. 注册用户
curl -X POST "http://localhost:8080/api/v1/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'

# 2. 用户登录
curl -X POST "http://localhost:8080/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'

# 3. 使用返回的 token 访问受保护的接口
curl -X GET "http://localhost:8080/api/v1/user/profile" \
  -H "Authorization: Bearer <your-jwt-token>"
```

### 管理员操作示例

```bash
# 1. 注册管理员账户
curl -X POST "http://localhost:8080/api/v1/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "email": "admin@example.com",
    "password": "admin123",
    "role": "admin"
  }'

# 2. 管理员登录
curl -X POST "http://localhost:8080/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'

# 3. 获取用户列表
curl -X GET "http://localhost:8080/api/v1/admin/users" \
  -H "Authorization: Bearer <admin-jwt-token>"

# 4. 更新用户角色
curl -X PUT "http://localhost:8080/api/v1/admin/users/1" \
  -H "Authorization: Bearer <admin-jwt-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "role": "admin"
  }'
```