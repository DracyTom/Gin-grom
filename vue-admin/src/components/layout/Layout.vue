<template>
  <div class="layout-container">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside :width="sidebarWidth" class="sidebar-container">
        <Sidebar />
      </el-aside>
      
      <!-- 主要内容区域 -->
      <el-container>
        <!-- 顶部导航栏 -->
        <el-header height="60px" class="header-container">
          <Navbar />
        </el-header>
        
        <!-- 主要内容 -->
        <el-main class="main-content">
          <router-view v-slot="{ Component }">
            <transition name="fade" mode="out-in">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import { computed, onMounted, onUnmounted } from 'vue'
import { useAppStore } from '@/stores/app'
import Sidebar from './Sidebar.vue'
import Navbar from './Navbar.vue'

const appStore = useAppStore()

const sidebarWidth = computed(() => {
  return appStore.sidebarOpened ? '200px' : '64px'
})

// 响应式处理
const handleResize = () => {
  const width = window.innerWidth
  if (width < 768) {
    appStore.setDevice('mobile')
    appStore.closeSidebar(true)
  } else {
    appStore.setDevice('desktop')
  }
}

onMounted(() => {
  handleResize()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
})
</script>

<style lang="scss" scoped>
.layout-container {
  height: 100vh;
  
  .el-container {
    height: 100%;
  }
  
  .sidebar-container {
    background-color: #304156;
    transition: width 0.3s;
    overflow: hidden;
  }
  
  .header-container {
    background-color: #fff;
    border-bottom: 1px solid #e4e7ed;
    padding: 0;
    display: flex;
    align-items: center;
  }
  
  .main-content {
    background-color: #f5f5f5;
    padding: 20px;
    overflow-y: auto;
  }
}

@media (max-width: 768px) {
  .sidebar-container {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 1000;
    height: 100vh;
  }
  
  .main-content {
    padding: 10px;
  }
}
</style>