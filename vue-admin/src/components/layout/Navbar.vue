<template>
  <div class="navbar">
    <!-- 左侧 -->
    <div class="navbar-left">
      <el-button
        type="text"
        @click="toggleSidebar"
        class="sidebar-toggle"
      >
        <el-icon size="20">
          <Fold v-if="appStore.sidebarOpened" />
          <Expand v-else />
        </el-icon>
      </el-button>
      
      <!-- 面包屑 -->
      <el-breadcrumb separator="/" class="breadcrumb">
        <el-breadcrumb-item
          v-for="item in breadcrumbs"
          :key="item.path"
          :to="item.path"
        >
          {{ item.title }}
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    
    <!-- 右侧 -->
    <div class="navbar-right">
      <!-- 主题切换 -->
      <el-button
        type="text"
        @click="toggleTheme"
        class="theme-toggle"
      >
        <el-icon size="18">
          <Sunny v-if="appStore.theme === 'light'" />
          <Moon v-else />
        </el-icon>
      </el-button>
      
      <!-- 用户菜单 -->
      <el-dropdown @command="handleCommand" class="user-dropdown">
        <div class="user-info">
          <el-avatar
            :src="userStore.avatar"
            :size="32"
            class="user-avatar"
          >
            <el-icon><User /></el-icon>
          </el-avatar>
          <span class="username">{{ userStore.username }}</span>
          <el-icon class="dropdown-icon"><ArrowDown /></el-icon>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="profile">
              <el-icon><User /></el-icon>
              个人中心
            </el-dropdown-item>
            <el-dropdown-item command="settings">
              <el-icon><Setting /></el-icon>
              设置
            </el-dropdown-item>
            <el-dropdown-item divided command="logout">
              <el-icon><SwitchButton /></el-icon>
              退出登录
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from '@/stores/app'
import { useUserStore } from '@/stores/user'
import { ElMessageBox } from 'element-plus'
import {
  Fold,
  Expand,
  Sunny,
  Moon,
  User,
  ArrowDown,
  Setting,
  SwitchButton
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()
const userStore = useUserStore()

// 面包屑导航
const breadcrumbs = computed(() => {
  const matched = route.matched.filter(item => item.meta && item.meta.title)
  const breadcrumbs = []
  
  matched.forEach(item => {
    breadcrumbs.push({
      title: item.meta.title,
      path: item.path
    })
  })
  
  return breadcrumbs
})

// 切换侧边栏
const toggleSidebar = () => {
  appStore.toggleSidebar()
}

// 切换主题
const toggleTheme = () => {
  appStore.toggleTheme()
}

// 处理用户菜单命令
const handleCommand = (command) => {
  switch (command) {
    case 'profile':
      router.push('/profile')
      break
    case 'settings':
      // TODO: 实现设置页面
      break
    case 'logout':
      handleLogout()
      break
  }
}

// 退出登录
const handleLogout = () => {
  ElMessageBox.confirm(
    '确定要退出登录吗？',
    '提示',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(() => {
    userStore.logout()
    router.push('/login')
  })
}
</script>

<style lang="scss" scoped>
.navbar {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  
  .navbar-left {
    display: flex;
    align-items: center;
    
    .sidebar-toggle {
      margin-right: 20px;
      color: #5a5e66;
      
      &:hover {
        color: #409eff;
      }
    }
    
    .breadcrumb {
      font-size: 14px;
    }
  }
  
  .navbar-right {
    display: flex;
    align-items: center;
    gap: 15px;
    
    .theme-toggle {
      color: #5a5e66;
      
      &:hover {
        color: #409eff;
      }
    }
    
    .user-dropdown {
      .user-info {
        display: flex;
        align-items: center;
        cursor: pointer;
        padding: 5px 10px;
        border-radius: 4px;
        transition: background-color 0.3s;
        
        &:hover {
          background-color: #f5f7fa;
        }
        
        .user-avatar {
          margin-right: 8px;
        }
        
        .username {
          font-size: 14px;
          color: #303133;
          margin-right: 5px;
        }
        
        .dropdown-icon {
          font-size: 12px;
          color: #909399;
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .navbar {
    padding: 0 10px;
    
    .breadcrumb {
      display: none;
    }
    
    .username {
      display: none;
    }
  }
}
</style>