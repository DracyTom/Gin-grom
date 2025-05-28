<template>
  <div class="sidebar">
    <!-- Logo -->
    <div class="logo-container">
      <div class="logo">
        <el-icon v-if="!appStore.sidebarOpened" size="24">
          <Setting />
        </el-icon>
        <span v-else class="logo-text">Vue Admin</span>
      </div>
    </div>
    
    <!-- 菜单 -->
    <el-menu
      :default-active="activeMenu"
      :collapse="!appStore.sidebarOpened"
      :unique-opened="true"
      background-color="#304156"
      text-color="#bfcbd9"
      active-text-color="#409eff"
      router
    >
      <el-menu-item index="/dashboard">
        <el-icon><House /></el-icon>
        <template #title>仪表盘</template>
      </el-menu-item>
      
      <el-menu-item index="/profile">
        <el-icon><User /></el-icon>
        <template #title>个人中心</template>
      </el-menu-item>
      
      <el-sub-menu v-if="userStore.isAdmin" index="/admin">
        <template #title>
          <el-icon><UserFilled /></el-icon>
          <span>用户管理</span>
        </template>
        <el-menu-item index="/admin/users">用户列表</el-menu-item>
      </el-sub-menu>
    </el-menu>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useAppStore } from '@/stores/app'
import { useUserStore } from '@/stores/user'
import { House, User, UserFilled, Setting } from '@element-plus/icons-vue'

const route = useRoute()
const appStore = useAppStore()
const userStore = useUserStore()

const activeMenu = computed(() => {
  const { path } = route
  if (path.startsWith('/admin')) {
    return '/admin/users'
  }
  return path
})
</script>

<style lang="scss" scoped>
.sidebar {
  height: 100%;
  
  .logo-container {
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #2b2f3a;
    
    .logo {
      display: flex;
      align-items: center;
      color: #fff;
      font-size: 18px;
      font-weight: bold;
      
      .logo-text {
        margin-left: 8px;
      }
    }
  }
  
  .el-menu {
    border-right: none;
    height: calc(100vh - 60px);
    overflow-y: auto;
    
    &:not(.el-menu--collapse) {
      width: 200px;
    }
  }
}
</style>