<template>
  <div class="dashboard">
    <!-- 欢迎信息 -->
    <div class="welcome-card content-card">
      <div class="welcome-content">
        <div class="welcome-text">
          <h2>欢迎回来，{{ userStore.username }}！</h2>
          <p>今天是 {{ currentDate }}，祝您工作愉快！</p>
        </div>
        <div class="welcome-avatar">
          <el-avatar :size="80" :src="userStore.avatar">
            <el-icon size="40"><User /></el-icon>
          </el-avatar>
        </div>
      </div>
    </div>
    
    <!-- 统计卡片 -->
    <el-row :gutter="20" class="stats-row">
      <el-col :xs="24" :sm="12" :md="6" v-for="stat in stats" :key="stat.title">
        <div class="stat-card content-card">
          <div class="stat-icon" :style="{ backgroundColor: stat.color }">
            <el-icon size="24">
              <component :is="stat.icon" />
            </el-icon>
          </div>
          <div class="stat-content">
            <div class="stat-value">{{ stat.value }}</div>
            <div class="stat-title">{{ stat.title }}</div>
          </div>
        </div>
      </el-col>
    </el-row>
    
    <!-- 图表区域 -->
    <el-row :gutter="20" class="charts-row">
      <el-col :xs="24" :lg="12">
        <div class="chart-card content-card">
          <h3>用户增长趋势</h3>
          <div class="chart-placeholder">
            <el-icon size="64" color="#ddd"><TrendCharts /></el-icon>
            <p>图表组件待集成</p>
          </div>
        </div>
      </el-col>
      
      <el-col :xs="24" :lg="12">
        <div class="chart-card content-card">
          <h3>系统状态</h3>
          <div class="system-status">
            <div class="status-item">
              <span class="status-label">CPU 使用率</span>
              <el-progress :percentage="cpuUsage" :color="getProgressColor(cpuUsage)" />
            </div>
            <div class="status-item">
              <span class="status-label">内存使用率</span>
              <el-progress :percentage="memoryUsage" :color="getProgressColor(memoryUsage)" />
            </div>
            <div class="status-item">
              <span class="status-label">磁盘使用率</span>
              <el-progress :percentage="diskUsage" :color="getProgressColor(diskUsage)" />
            </div>
          </div>
        </div>
      </el-col>
    </el-row>
    
    <!-- 快捷操作 -->
    <div class="quick-actions content-card">
      <h3>快捷操作</h3>
      <el-row :gutter="20">
        <el-col :xs="12" :sm="8" :md="6" v-for="action in quickActions" :key="action.title">
          <div class="action-item" @click="handleQuickAction(action.action)">
            <el-icon size="32" :color="action.color">
              <component :is="action.icon" />
            </el-icon>
            <span>{{ action.title }}</span>
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { formatTime } from '@/utils'
import {
  User,
  UserFilled,
  DataAnalysis,
  Monitor,
  Setting,
  TrendCharts,
  Plus,
  Edit,
  View,
  Refresh
} from '@element-plus/icons-vue'

const router = useRouter()
const userStore = useUserStore()

// 当前日期
const currentDate = computed(() => {
  return formatTime(new Date(), 'YYYY年MM月DD日 dddd')
})

// 统计数据
const stats = ref([
  {
    title: '总用户数',
    value: '1,234',
    icon: UserFilled,
    color: '#409eff'
  },
  {
    title: '今日访问',
    value: '567',
    icon: DataAnalysis,
    color: '#67c23a'
  },
  {
    title: '系统运行',
    value: '99.9%',
    icon: Monitor,
    color: '#e6a23c'
  },
  {
    title: '活跃用户',
    value: '89',
    icon: User,
    color: '#f56c6c'
  }
])

// 系统状态
const cpuUsage = ref(45)
const memoryUsage = ref(67)
const diskUsage = ref(23)

// 快捷操作
const quickActions = ref([
  {
    title: '新增用户',
    icon: Plus,
    color: '#409eff',
    action: 'addUser'
  },
  {
    title: '编辑资料',
    icon: Edit,
    color: '#67c23a',
    action: 'editProfile'
  },
  {
    title: '查看日志',
    icon: View,
    color: '#e6a23c',
    action: 'viewLogs'
  },
  {
    title: '系统设置',
    icon: Setting,
    color: '#f56c6c',
    action: 'settings'
  }
])

// 获取进度条颜色
const getProgressColor = (percentage) => {
  if (percentage < 50) return '#67c23a'
  if (percentage < 80) return '#e6a23c'
  return '#f56c6c'
}

// 处理快捷操作
const handleQuickAction = (action) => {
  switch (action) {
    case 'addUser':
      if (userStore.isAdmin) {
        router.push('/admin/users')
      }
      break
    case 'editProfile':
      router.push('/profile')
      break
    case 'viewLogs':
      // TODO: 实现日志查看
      break
    case 'settings':
      // TODO: 实现系统设置
      break
  }
}

// 模拟数据更新
onMounted(() => {
  const updateSystemStatus = () => {
    cpuUsage.value = Math.floor(Math.random() * 100)
    memoryUsage.value = Math.floor(Math.random() * 100)
    diskUsage.value = Math.floor(Math.random() * 100)
  }
  
  // 每30秒更新一次系统状态
  const timer = setInterval(updateSystemStatus, 30000)
  
  // 组件卸载时清除定时器
  onUnmounted(() => {
    clearInterval(timer)
  })
})
</script>

<style lang="scss" scoped>
.dashboard {
  .welcome-card {
    margin-bottom: 20px;
    
    .welcome-content {
      display: flex;
      align-items: center;
      justify-content: space-between;
      
      .welcome-text {
        h2 {
          color: #303133;
          margin-bottom: 8px;
          font-size: 24px;
        }
        
        p {
          color: #909399;
          font-size: 14px;
        }
      }
    }
  }
  
  .stats-row {
    margin-bottom: 20px;
    
    .stat-card {
      display: flex;
      align-items: center;
      padding: 20px;
      
      .stat-icon {
        width: 60px;
        height: 60px;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-right: 15px;
        color: white;
      }
      
      .stat-content {
        .stat-value {
          font-size: 24px;
          font-weight: bold;
          color: #303133;
          margin-bottom: 5px;
        }
        
        .stat-title {
          font-size: 14px;
          color: #909399;
        }
      }
    }
  }
  
  .charts-row {
    margin-bottom: 20px;
    
    .chart-card {
      h3 {
        margin-bottom: 20px;
        color: #303133;
      }
      
      .chart-placeholder {
        height: 300px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        color: #909399;
        
        p {
          margin-top: 10px;
          font-size: 14px;
        }
      }
      
      .system-status {
        .status-item {
          margin-bottom: 20px;
          
          .status-label {
            display: block;
            margin-bottom: 8px;
            font-size: 14px;
            color: #606266;
          }
        }
      }
    }
  }
  
  .quick-actions {
    h3 {
      margin-bottom: 20px;
      color: #303133;
    }
    
    .action-item {
      display: flex;
      flex-direction: column;
      align-items: center;
      padding: 20px;
      border-radius: 8px;
      cursor: pointer;
      transition: all 0.3s;
      
      &:hover {
        background-color: #f5f7fa;
        transform: translateY(-2px);
      }
      
      span {
        margin-top: 10px;
        font-size: 14px;
        color: #606266;
      }
    }
  }
}

@media (max-width: 768px) {
  .dashboard {
    .welcome-card .welcome-content {
      flex-direction: column;
      text-align: center;
      
      .welcome-avatar {
        margin-top: 20px;
      }
    }
    
    .stats-row .stat-card {
      margin-bottom: 10px;
    }
  }
}
</style>