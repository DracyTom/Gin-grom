<template>
  <div class="user-management">
    <!-- 搜索和操作栏 -->
    <div class="toolbar content-card">
      <div class="search-section">
        <el-input
          v-model="searchForm.keyword"
          placeholder="搜索用户名或邮箱"
          style="width: 300px"
          clearable
          @keyup.enter="handleSearch"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        <el-select
          v-model="searchForm.role"
          placeholder="选择角色"
          style="width: 120px; margin-left: 10px"
          clearable
        >
          <el-option label="全部" value="" />
          <el-option label="用户" value="user" />
          <el-option label="管理员" value="admin" />
        </el-select>
        <el-select
          v-model="searchForm.status"
          placeholder="选择状态"
          style="width: 120px; margin-left: 10px"
          clearable
        >
          <el-option label="全部" value="" />
          <el-option label="正常" :value="1" />
          <el-option label="禁用" :value="0" />
        </el-select>
        <el-button type="primary" @click="handleSearch" style="margin-left: 10px">
          <el-icon><Search /></el-icon>
          搜索
        </el-button>
        <el-button @click="resetSearch">
          <el-icon><Refresh /></el-icon>
          重置
        </el-button>
      </div>
      
      <div class="action-section">
        <el-button type="primary" @click="showAddDialog">
          <el-icon><Plus /></el-icon>
          新增用户
        </el-button>
        <el-button type="danger" :disabled="!selectedUsers.length" @click="handleBatchDelete">
          <el-icon><Delete /></el-icon>
          批量删除
        </el-button>
      </div>
    </div>
    
    <!-- 用户表格 -->
    <div class="table-container content-card">
      <el-table
        v-loading="loading"
        :data="userList"
        @selection-change="handleSelectionChange"
        stripe
        style="width: 100%"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="username" label="用户名" min-width="120" />
        <el-table-column prop="email" label="邮箱" min-width="180" />
        <el-table-column prop="role" label="角色" width="100">
          <template #default="{ row }">
            <el-tag :type="row.role === 'admin' ? 'danger' : 'primary'">
              {{ row.role === 'admin' ? '管理员' : '普通用户' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '正常' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatTime(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="showEditDialog(row)">
              编辑
            </el-button>
            <el-button
              :type="row.status === 1 ? 'warning' : 'success'"
              size="small"
              @click="toggleUserStatus(row)"
            >
              {{ row.status === 1 ? '禁用' : '启用' }}
            </el-button>
            <el-button type="danger" size="small" @click="handleDelete(row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <!-- 分页 -->
      <div class="pagination-container">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.size"
          :page-sizes="[10, 20, 50, 100]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </div>
    
    <!-- 新增/编辑用户对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      @close="resetForm"
    >
      <el-form
        ref="userFormRef"
        :model="userForm"
        :rules="userRules"
        label-width="100px"
      >
        <el-form-item label="用户名" prop="username">
          <el-input v-model="userForm.username" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="userForm.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item v-if="!isEdit" label="密码" prop="password">
          <el-input
            v-model="userForm.password"
            type="password"
            placeholder="请输入密码"
            show-password
          />
        </el-form-item>
        <el-form-item label="角色" prop="role">
          <el-select v-model="userForm.role" placeholder="请选择角色" style="width: 100%">
            <el-option label="普通用户" value="user" />
            <el-option label="管理员" value="admin" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="userForm.status">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="头像" prop="avatar">
          <el-input v-model="userForm.avatar" placeholder="请输入头像URL（可选）" />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="submitLoading" @click="handleSubmit">
          确定
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { reactive, ref, onMounted, computed } from 'vue'
import { adminAPI, authAPI } from '@/api/user'
import { formatTime, validateEmail } from '@/utils'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Search,
  Refresh,
  Plus,
  Delete
} from '@element-plus/icons-vue'

// 响应式数据
const loading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const isEdit = ref(false)
const userList = ref([])
const selectedUsers = ref([])

// 搜索表单
const searchForm = reactive({
  keyword: '',
  role: '',
  status: ''
})

// 分页信息
const pagination = reactive({
  page: 1,
  size: 10,
  total: 0
})

// 用户表单
const userForm = reactive({
  id: null,
  username: '',
  email: '',
  password: '',
  role: 'user',
  status: 1,
  avatar: ''
})

const userFormRef = ref()

// 对话框标题
const dialogTitle = computed(() => {
  return isEdit.value ? '编辑用户' : '新增用户'
})

// 表单验证规则
const userRules = {
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
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于 6 个字符', trigger: 'blur' }
  ],
  role: [
    { required: true, message: '请选择角色', trigger: 'change' }
  ]
}

// 获取用户列表
const getUserList = async () => {
  try {
    loading.value = true
    const params = {
      page: pagination.page,
      size: pagination.size,
      ...searchForm
    }
    
    const data = await adminAPI.getUserList(params)
    userList.value = data.data || data || []
    pagination.total = data.total || userList.value.length
  } catch (error) {
    console.error('获取用户列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  getUserList()
}

// 重置搜索
const resetSearch = () => {
  searchForm.keyword = ''
  searchForm.role = ''
  searchForm.status = ''
  pagination.page = 1
  getUserList()
}

// 分页大小改变
const handleSizeChange = (size) => {
  pagination.size = size
  pagination.page = 1
  getUserList()
}

// 当前页改变
const handleCurrentChange = (page) => {
  pagination.page = page
  getUserList()
}

// 选择改变
const handleSelectionChange = (selection) => {
  selectedUsers.value = selection
}

// 显示新增对话框
const showAddDialog = () => {
  isEdit.value = false
  dialogVisible.value = true
}

// 显示编辑对话框
const showEditDialog = (row) => {
  isEdit.value = true
  userForm.id = row.id
  userForm.username = row.username
  userForm.email = row.email
  userForm.role = row.role
  userForm.status = row.status
  userForm.avatar = row.avatar || ''
  userForm.password = ''
  dialogVisible.value = true
}

// 重置表单
const resetForm = () => {
  userForm.id = null
  userForm.username = ''
  userForm.email = ''
  userForm.password = ''
  userForm.role = 'user'
  userForm.status = 1
  userForm.avatar = ''
  
  if (userFormRef.value) {
    userFormRef.value.clearValidate()
  }
}

// 提交表单
const handleSubmit = async () => {
  if (!userFormRef.value) return
  
  try {
    await userFormRef.value.validate()
    submitLoading.value = true
    
    if (isEdit.value) {
      // 编辑用户
      const { id, password, ...updateData } = userForm
      await adminAPI.updateUser(id, updateData)
      ElMessage.success('用户更新成功')
    } else {
      // 新增用户
      const { id, ...createData } = userForm
      await authAPI.register(createData)
      ElMessage.success('用户创建成功')
    }
    
    dialogVisible.value = false
    getUserList()
  } catch (error) {
    console.error('操作失败:', error)
  } finally {
    submitLoading.value = false
  }
}

// 切换用户状态
const toggleUserStatus = async (row) => {
  try {
    const newStatus = row.status === 1 ? 0 : 1
    await adminAPI.updateUser(row.id, { status: newStatus })
    ElMessage.success(`用户${newStatus === 1 ? '启用' : '禁用'}成功`)
    getUserList()
  } catch (error) {
    console.error('状态切换失败:', error)
  }
}

// 删除用户
const handleDelete = (row) => {
  ElMessageBox.confirm(
    `确定要删除用户 "${row.username}" 吗？`,
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async () => {
    try {
      await adminAPI.deleteUser(row.id)
      ElMessage.success('删除成功')
      getUserList()
    } catch (error) {
      console.error('删除失败:', error)
    }
  })
}

// 批量删除
const handleBatchDelete = () => {
  const usernames = selectedUsers.value.map(user => user.username).join('、')
  ElMessageBox.confirm(
    `确定要删除选中的 ${selectedUsers.value.length} 个用户吗？\n用户：${usernames}`,
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(async () => {
    try {
      const promises = selectedUsers.value.map(user => adminAPI.deleteUser(user.id))
      await Promise.all(promises)
      ElMessage.success('批量删除成功')
      getUserList()
    } catch (error) {
      console.error('批量删除失败:', error)
    }
  })
}

onMounted(() => {
  getUserList()
})
</script>

<style lang="scss" scoped>
.user-management {
  .toolbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
    padding: 20px;
    
    .search-section {
      display: flex;
      align-items: center;
    }
    
    .action-section {
      display: flex;
      gap: 10px;
    }
  }
  
  .table-container {
    .pagination-container {
      display: flex;
      justify-content: center;
      margin-top: 20px;
    }
  }
}

@media (max-width: 768px) {
  .user-management {
    .toolbar {
      flex-direction: column;
      gap: 15px;
      
      .search-section {
        flex-wrap: wrap;
        gap: 10px;
      }
    }
  }
}
</style>