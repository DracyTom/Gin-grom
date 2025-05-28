import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    // 侧边栏状态
    sidebar: {
      opened: true,
      withoutAnimation: false
    },
    // 设备类型
    device: 'desktop',
    // 主题
    theme: 'light',
    // 语言
    language: 'zh-cn',
    // 页面加载状态
    loading: false
  }),
  
  getters: {
    sidebarOpened: (state) => state.sidebar.opened,
    isMobile: (state) => state.device === 'mobile'
  },
  
  actions: {
    // 切换侧边栏
    toggleSidebar() {
      this.sidebar.opened = !this.sidebar.opened
      this.sidebar.withoutAnimation = false
    },
    
    // 关闭侧边栏
    closeSidebar(withoutAnimation = false) {
      this.sidebar.opened = false
      this.sidebar.withoutAnimation = withoutAnimation
    },
    
    // 设置设备类型
    setDevice(device) {
      this.device = device
    },
    
    // 切换主题
    toggleTheme() {
      this.theme = this.theme === 'light' ? 'dark' : 'light'
      document.documentElement.className = this.theme
    },
    
    // 设置语言
    setLanguage(language) {
      this.language = language
    },
    
    // 设置加载状态
    setLoading(loading) {
      this.loading = loading
    }
  },
  
  persist: {
    key: 'app-settings',
    storage: localStorage,
    paths: ['sidebar', 'theme', 'language']
  }
})