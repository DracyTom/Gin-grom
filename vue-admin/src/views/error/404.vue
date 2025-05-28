<template>
  <div class="error-page">
    <div class="error-content">
      <div class="error-code">404</div>
      <div class="error-title">页面不存在</div>
      <div class="error-description">
        抱歉，您访问的页面不存在或已被删除
      </div>
      <div class="error-actions">
        <el-button type="primary" @click="goBack">
          返回上一页
        </el-button>
        <el-button @click="goHome">
          回到首页
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goBack = () => {
  router.go(-1)
}

const goHome = () => {
  router.push('/')
}
</script>

<style lang="scss" scoped>
.error-page {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f5f5f5;
  
  .error-content {
    text-align: center;
    
    .error-code {
      font-size: 120px;
      font-weight: bold;
      color: #409eff;
      margin-bottom: 20px;
    }
    
    .error-title {
      font-size: 32px;
      color: #303133;
      margin-bottom: 16px;
    }
    
    .error-description {
      font-size: 16px;
      color: #909399;
      margin-bottom: 40px;
    }
    
    .error-actions {
      display: flex;
      gap: 16px;
      justify-content: center;
    }
  }
}
</style>