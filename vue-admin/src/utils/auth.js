import Cookies from 'js-cookie'

const TOKEN_KEY = 'vue-admin-token'

// 获取 token
export function getToken() {
  return Cookies.get(TOKEN_KEY)
}

// 设置 token
export function setToken(token) {
  return Cookies.set(TOKEN_KEY, token, { expires: 7 }) // 7天过期
}

// 移除 token
export function removeToken() {
  return Cookies.remove(TOKEN_KEY)
}

// 检查是否有 token
export function hasToken() {
  return !!getToken()
}