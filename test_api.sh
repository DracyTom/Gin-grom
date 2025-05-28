#!/bin/bash

# API 测试脚本
BASE_URL="http://localhost:8080"

echo "=== Gin-GORM 脚手架 API 测试 ==="
echo

# 测试健康检查
echo "1. 测试健康检查..."
curl -s "$BASE_URL/health" | jq .
echo

# 测试用户注册
echo "2. 测试用户注册..."
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/api/v1/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }')
echo $REGISTER_RESPONSE | jq .
echo

# 测试用户登录
echo "3. 测试用户登录..."
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }')
echo $LOGIN_RESPONSE | jq .

# 提取 token
TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.data.token')
echo "Token: $TOKEN"
echo

# 测试获取个人信息
echo "4. 测试获取个人信息..."
curl -s -X GET "$BASE_URL/api/v1/user/profile" \
  -H "Authorization: Bearer $TOKEN" | jq .
echo

# 测试更新个人信息
echo "5. 测试更新个人信息..."
curl -s -X PUT "$BASE_URL/api/v1/user/profile" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "updateduser",
    "email": "updated@example.com",
    "avatar": "http://example.com/avatar.jpg"
  }' | jq .
echo

# 测试管理员登录
echo "6. 测试管理员登录..."
ADMIN_LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }')
echo $ADMIN_LOGIN_RESPONSE | jq .

# 提取管理员 token
ADMIN_TOKEN=$(echo $ADMIN_LOGIN_RESPONSE | jq -r '.data.token')
echo "Admin Token: $ADMIN_TOKEN"
echo

# 测试获取用户列表
echo "7. 测试获取用户列表..."
curl -s -X GET "$BASE_URL/api/v1/admin/users?page=1&size=10" \
  -H "Authorization: Bearer $ADMIN_TOKEN" | jq .
echo

echo "=== 测试完成 ==="