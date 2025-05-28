<template>
  <div class="profile">
    <el-row :gutter="20">
      <!-- 左侧个人信息 -->
      <el-col :xs="24" :lg="8">
        <div class="profile-card content-card">
          <div class="profile-header">
            <el-avatar :size="100" :src="userInfo.avatar">
              <el-icon size="50"><User /></el-icon>
            </el-avatar>
            <h3>{{ userInfo.username }}</h3>
            <p>{{ userInfo.email }}</p>
            <el-tag :type="userInfo.role === 'admin' ? 'danger' : 'primary'">
              {{ userInfo.role === 'admin' ? '管理员' : '普通用户' }}
            </el-tag>
          </div>
          
          <div class="profile-info">
            <div class="info-item">
              <span class="label">用户ID:</span>
              <span class="value">{{ userInfo.id }}</span>
            </div>
            <div class="info-item">
              <span class="label">注册时间:</span>
              <span class="value">{{ formatTime(userInfo.created_at) }}</span>
            </div>
            <div class="info-item">
              <span class="label">最后更新:</span>
              <span class="value">{{ formatTime(userInfo.updated_at) }}</span>
            </div>
            <div class="info-item">
              <span class="label">账号状态:</span>
              <el-tag :type="userInfo.status === 1 ? 'success' : 'danger'">
                {{ userInfo.status === 1 ? '正常' : '禁用' }}
              </el-tag>
            </div>
          </div>
        </div>
      </el-col>
      
      <!-- 右侧编辑表单 -->
      <el-col :xs="24" :lg="16">
        <div class="edit-card content-card">
          <h3>编辑个人信息</h3>
          
          <el-form
            ref="profileFormRef"
            :model="profileForm"
            :rules="profileRules"
            label-width="100px"
            size="large"
          >
            <el-form-item label="用户名" prop="username">
              <el-input
                v-model="profileForm.username"
                placeholder="请输入用户名"
                clearable
              />
            </el-form-item>
            
            <el-form-item label="邮箱" prop="email">
              <el-input
                v-model="profileForm.email"
                placeholder="请输入邮箱"
                clearable
              />
            </el-form-item>
            
            <el-form-item label="头像" prop="avatar">
              <el-input
                v-model="profileForm.avatar"
                placeholder="请输入头像URL"
                clearable
              >
                <template #append>
                  <el-button @click="previewAvatar">预览</el-button>
                </template>
              </el-input>
            </el-form-item>
            
            <el-form-item>
              <el-button
                type="primary"
                :loading="loading"
                @click="handleUpdate"
              >
                保存修改
              </el-button>
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </el-form>
        </div>
        
        <!-- 修改密码 -->
        <div class="password-card content-card">
          <h3>修改密码</h3>
          
          <el-form
            ref="passwordFormRef"
            :model="passwordForm"
            :rules="passwordRules"
            label-width="100px"
            size="large"
          >
            <el-form-item label="当前密码" prop="currentPassword">
              <el-input
                v-model="passwordForm.currentPassword"
                type="password"
                placeholder="请输入当前密码"
                show-password
                clearable
              />
            </el-form-item>
            
            <el-form-item label="新密码" prop="newPassword">
              <el-input
                v-model="passwordForm.newPassword"
                type="password"
                placeholder="请输入新密码"
                show-password
                clearable
              />
            </el-form-item>
            
            <el-form-item label="确认密码" prop="confirmPassword">
              <el-input
                v-model="passwordForm.confirmPassword"
                type="password"
                placeholder="请确认新密码"
                show-password
                clearable
              />
            </el-form-item>
            
            <el-form-item>
              <el-button
                type="primary"
                :loading="passwordLoading"
                @click="handlePasswordUpdate"
              >
                修改密码
              </el-button>
              <el-button @click="resetPasswordForm">重置</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-col>
    </el-row>
    
    <!-- 头像预览对话框 -->
    <el-dialog
      v-model="avatarDialogVisible"
      title="头像预览"
      width="400px"
      center
    >
      <div class="avatar-preview">
        <el-avatar :size="200" :src="profileForm.avatar">
          <el-icon size="100"><User /></el-icon>
        </el-avatar>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { reactive, ref, computed, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { formatTime, validateEmail } from '@/utils'
import { ElMessage } from 'element-plus'
import { User } from '@element-plus/icons-vue'

const userStore = useUserStore()

const profileFormRef = ref()
const passwordFormRef = ref()
const loading = ref(false)
const passwordLoading = ref(false)
const avatarDialogVisible = ref(false)

// 用户信息
const userInfo = computed(() => userStore.userInfo || {})

// 个人信息表单
const profileForm = reactive({
  username: '',
  email: '',
  avatar: ''
})

// 密码表单
const passwordForm = reactive({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
})

// 个人信息验证规则
const profileRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 50, message: '用户名长度在 3 到 50 个字符', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (!validateEmail(value)) {
          callback(new Error('请输入正确的邮箱格式'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

// 密码验证规则
const passwordRules = {
  currentPassword: [
    { required: true, message: '请输入当前密码', trigger: 'blur' }
  ],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于 6 个字符', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请确认新密码', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== passwordForm.newPassword) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

// 初始化表单数据
const initForm = () => {
  if (userInfo.value) {
    profileForm.username = userInfo.value.username || ''
    profileForm.email = userInfo.value.email || ''
    profileForm.avatar = userInfo.value.avatar || ''
  }
}

// 更新个人信息
const handleUpdate = async () => {
  if (!profileFormRef.value) return
  
  try {
    await profileFormRef.value.validate()
    loading.value = true
    
    await userStore.updateUserInfo(profileForm)
  } catch (error) {
    console.error('更新失败:', error)
  } finally {
    loading.value = false
  }
}

// 修改密码
const handlePasswordUpdate = async () => {
  if (!passwordFormRef.value) return
  
  try {
    await passwordFormRef.value.validate()
    passwordLoading.value = true
    
    // TODO: 实现密码修改 API
    ElMessage.success('密码修改成功')
    resetPasswordForm()
  } catch (error) {
    console.error('密码修改失败:', error)
  } finally {
    passwordLoading.value = false
  }
}

// 重置个人信息表单
const resetForm = () => {
  initForm()
}

// 重置密码表单
const resetPasswordForm = () => {
  passwordForm.currentPassword = ''
  passwordForm.newPassword = ''
  passwordForm.confirmPassword = ''
  if (passwordFormRef.value) {
    passwordFormRef.value.clearValidate()
  }
}

// 预览头像
const previewAvatar = () => {
  if (!profileForm.avatar) {
    ElMessage.warning('请先输入头像URL')
    return
  }
  avatarDialogVisible.value = true
}

onMounted(() => {
  initForm()
})
</script>

<style lang="scss" scoped>
.profile {
  .profile-card {
    .profile-header {
      text-align: center;
      padding-bottom: 20px;
      border-bottom: 1px solid #ebeef5;
      margin-bottom: 20px;
      
      h3 {
        margin: 15px 0 5px;
        color: #303133;
      }
      
      p {
        color: #909399;
        margin-bottom: 10px;
      }
    }
    
    .profile-info {
      .info-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 10px 0;
        border-bottom: 1px solid #f5f7fa;
        
        &:last-child {
          border-bottom: none;
        }
        
        .label {
          color: #909399;
          font-size: 14px;
        }
        
        .value {
          color: #303133;
          font-size: 14px;
        }
      }
    }
  }
  
  .edit-card,
  .password-card {
    margin-bottom: 20px;
    
    h3 {
      margin-bottom: 20px;
      color: #303133;
      padding-bottom: 10px;
      border-bottom: 1px solid #ebeef5;
    }
  }
  
  .avatar-preview {
    text-align: center;
    padding: 20px;
  }
}

@media (max-width: 768px) {
  .profile {
    .profile-card {
      margin-bottom: 20px;
    }
  }
}
</style>