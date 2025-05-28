import request from './request'

// 用户认证相关
export const authAPI = {
  // 用户登录
  login(data) {
    return request.post('/v1/auth/login', data)
  },
  
  // 用户注册
  register(data) {
    return request.post('/v1/auth/register', data)
  },
  
  // 获取个人信息
  getProfile() {
    return request.get('/v1/user/profile')
  },
  
  // 更新个人信息
  updateProfile(data) {
    return request.put('/v1/user/profile', data)
  }
}

// 管理员用户管理相关
export const adminAPI = {
  // 获取用户列表
  getUserList(params) {
    return request.get('/v1/admin/users', { params })
  },
  
  // 获取指定用户信息
  getUserById(id) {
    return request.get(`/v1/admin/users/${id}`)
  },
  
  // 更新用户信息
  updateUser(id, data) {
    return request.put(`/v1/admin/users/${id}`, data)
  },
  
  // 删除用户
  deleteUser(id) {
    return request.delete(`/v1/admin/users/${id}`)
  }
}

// 系统相关
export const systemAPI = {
  // 健康检查
  healthCheck() {
    return request.get('/health')
  }
}