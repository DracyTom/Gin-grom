import { defineStore } from 'pinia'
import { authAPI } from '@/api/user'
import { getToken, setToken, removeToken } from '@/utils/auth'
import { ElMessage } from 'element-plus'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: getToken(),
    userInfo: null,
    permissions: []
  }),
  
  getters: {
    isLoggedIn: (state) => !!state.token,
    isAdmin: (state) => state.userInfo?.role === 'admin',
    username: (state) => state.userInfo?.username || '',
    avatar: (state) => state.userInfo?.avatar || '',
    email: (state) => state.userInfo?.email || ''
  },
  
  actions: {
    // 登录
    async login(loginForm) {
      try {
        const { token, user } = await authAPI.login(loginForm)
        
        this.token = token
        this.userInfo = user
        setToken(token)
        
        ElMessage.success('登录成功')
        return Promise.resolve()
      } catch (error) {
        return Promise.reject(error)
      }
    },
    
    // 注册
    async register(registerForm) {
      try {
        await authAPI.register(registerForm)
        ElMessage.success('注册成功，请登录')
        return Promise.resolve()
      } catch (error) {
        return Promise.reject(error)
      }
    },
    
    // 获取用户信息
    async getUserInfo() {
      try {
        const userInfo = await authAPI.getProfile()
        this.userInfo = userInfo
        return userInfo
      } catch (error) {
        this.logout()
        return Promise.reject(error)
      }
    },
    
    // 更新用户信息
    async updateUserInfo(data) {
      try {
        const userInfo = await authAPI.updateProfile(data)
        this.userInfo = userInfo
        ElMessage.success('更新成功')
        return userInfo
      } catch (error) {
        return Promise.reject(error)
      }
    },
    
    // 登出
    logout() {
      this.token = ''
      this.userInfo = null
      this.permissions = []
      removeToken()
    },
    
    // 检查认证状态
    async checkAuth() {
      if (this.token) {
        try {
          await this.getUserInfo()
        } catch (error) {
          this.logout()
        }
      }
    }
  }
})