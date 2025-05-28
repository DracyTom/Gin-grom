#!/bin/bash

# Gin-GROM 脚手架完整测试脚本
# 测试所有 API 端点和功能

BASE_URL="http://localhost:8080"
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}=== Gin-GROM 脚手架测试 ===${NC}"

# 测试健康检查
echo -e "\n${YELLOW}1. 测试健康检查${NC}"
response=$(curl -s "$BASE_URL/health")
echo "Response: $response"
if echo "$response" | jq -e '.status == "ok"' > /dev/null; then
    echo -e "${GREEN}✓ 健康检查通过${NC}"
else
    echo -e "${RED}✗ 健康检查失败${NC}"
    exit 1
fi

# 测试用户注册
echo -e "\n${YELLOW}2. 测试用户注册${NC}"
register_response=$(curl -s -X POST "$BASE_URL/api/v1/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser123",
    "email": "testuser123@example.com",
    "password": "password123"
  }')
echo "Response: $register_response"
if echo "$register_response" | jq -e '.code == 200' > /dev/null; then
    echo -e "${GREEN}✓ 用户注册成功${NC}"
else
    echo -e "${RED}✗ 用户注册失败${NC}"
fi

# 测试管理员注册
echo -e "\n${YELLOW}3. 测试管理员注册${NC}"
admin_register_response=$(curl -s -X POST "$BASE_URL/api/v1/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin123",
    "email": "admin123@example.com",
    "password": "admin123",
    "role": "admin"
  }')
echo "Response: $admin_register_response"
if echo "$admin_register_response" | jq -e '.code == 200 and .data.role == "admin"' > /dev/null; then
    echo -e "${GREEN}✓ 管理员注册成功${NC}"
else
    echo -e "${RED}✗ 管理员注册失败${NC}"
fi

# 测试用户登录
echo -e "\n${YELLOW}4. 测试用户登录${NC}"
login_response=$(curl -s -X POST "$BASE_URL/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser123",
    "password": "password123"
  }')
echo "Response: $login_response"
user_token=$(echo "$login_response" | jq -r '.data.token')
if [ "$user_token" != "null" ] && [ "$user_token" != "" ]; then
    echo -e "${GREEN}✓ 用户登录成功${NC}"
    echo "User Token: $user_token"
else
    echo -e "${RED}✗ 用户登录失败${NC}"
fi

# 测试管理员登录
echo -e "\n${YELLOW}5. 测试管理员登录${NC}"
admin_login_response=$(curl -s -X POST "$BASE_URL/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin123",
    "password": "admin123"
  }')
echo "Response: $admin_login_response"
admin_token=$(echo "$admin_login_response" | jq -r '.data.token')
if [ "$admin_token" != "null" ] && [ "$admin_token" != "" ]; then
    echo -e "${GREEN}✓ 管理员登录成功${NC}"
    echo "Admin Token: $admin_token"
else
    echo -e "${RED}✗ 管理员登录失败${NC}"
fi

# 测试获取个人信息
echo -e "\n${YELLOW}6. 测试获取个人信息${NC}"
profile_response=$(curl -s -X GET "$BASE_URL/api/v1/user/profile" \
  -H "Authorization: Bearer $user_token")
echo "Response: $profile_response"
if echo "$profile_response" | jq -e '.code == 200' > /dev/null; then
    echo -e "${GREEN}✓ 获取个人信息成功${NC}"
else
    echo -e "${RED}✗ 获取个人信息失败${NC}"
fi

# 测试更新个人信息
echo -e "\n${YELLOW}7. 测试更新个人信息${NC}"
update_profile_response=$(curl -s -X PUT "$BASE_URL/api/v1/user/profile" \
  -H "Authorization: Bearer $user_token" \
  -H "Content-Type: application/json" \
  -d '{
    "avatar": "https://example.com/new-avatar.jpg"
  }')
echo "Response: $update_profile_response"
if echo "$update_profile_response" | jq -e '.code == 200' > /dev/null; then
    echo -e "${GREEN}✓ 更新个人信息成功${NC}"
else
    echo -e "${RED}✗ 更新个人信息失败${NC}"
fi

# 测试管理员获取用户列表
echo -e "\n${YELLOW}8. 测试管理员获取用户列表${NC}"
users_response=$(curl -s -X GET "$BASE_URL/api/v1/admin/users" \
  -H "Authorization: Bearer $admin_token")
echo "Response: $users_response"
if echo "$users_response" | jq -e '.code == 200' > /dev/null; then
    echo -e "${GREEN}✓ 管理员获取用户列表成功${NC}"
    user_count=$(echo "$users_response" | jq '.total')
    echo "总用户数: $user_count"
else
    echo -e "${RED}✗ 管理员获取用户列表失败${NC}"
fi

# 测试普通用户访问管理员接口（应该被拒绝）
echo -e "\n${YELLOW}9. 测试权限控制（普通用户访问管理员接口）${NC}"
forbidden_response=$(curl -s -X GET "$BASE_URL/api/v1/admin/users" \
  -H "Authorization: Bearer $user_token")
echo "Response: $forbidden_response"
if echo "$forbidden_response" | jq -e '.code == 403' > /dev/null; then
    echo -e "${GREEN}✓ 权限控制正常（访问被拒绝）${NC}"
else
    echo -e "${RED}✗ 权限控制失败${NC}"
fi

# 测试无效 token
echo -e "\n${YELLOW}10. 测试无效 token${NC}"
invalid_token_response=$(curl -s -X GET "$BASE_URL/api/v1/user/profile" \
  -H "Authorization: Bearer invalid_token")
echo "Response: $invalid_token_response"
if echo "$invalid_token_response" | jq -e '.code == 401' > /dev/null; then
    echo -e "${GREEN}✓ 无效 token 处理正常${NC}"
else
    echo -e "${RED}✗ 无效 token 处理失败${NC}"
fi

# 测试重复注册
echo -e "\n${YELLOW}11. 测试重复注册${NC}"
duplicate_response=$(curl -s -X POST "$BASE_URL/api/v1/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser123",
    "email": "testuser123@example.com",
    "password": "password123"
  }')
echo "Response: $duplicate_response"
if echo "$duplicate_response" | jq -e '.code == 400' > /dev/null; then
    echo -e "${GREEN}✓ 重复注册检查正常${NC}"
else
    echo -e "${RED}✗ 重复注册检查失败${NC}"
fi

# 测试错误密码登录
echo -e "\n${YELLOW}12. 测试错误密码登录${NC}"
wrong_password_response=$(curl -s -X POST "$BASE_URL/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser123",
    "password": "wrongpassword"
  }')
echo "Response: $wrong_password_response"
if echo "$wrong_password_response" | jq -e '.code == 401' > /dev/null; then
    echo -e "${GREEN}✓ 错误密码检查正常${NC}"
else
    echo -e "${RED}✗ 错误密码检查失败${NC}"
fi

echo -e "\n${YELLOW}=== 测试完成 ===${NC}"
echo -e "${GREEN}Gin-GROM 脚手架功能测试完成！${NC}"
echo -e "${YELLOW}包含功能：${NC}"
echo "- ✓ 用户注册/登录"
echo "- ✓ JWT 认证"
echo "- ✓ 角色权限控制"
echo "- ✓ 个人信息管理"
echo "- ✓ 管理员功能"
echo "- ✓ 错误处理"
echo "- ✓ 数据验证"
echo "- ✓ 数据库集成"
echo "- ✓ 中间件系统"
echo "- ✓ RESTful API"