import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'
import NProgress from 'nprogress'

// 路由配置
const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/auth/Login.vue'),
    meta: {
      title: '登录',
      requiresAuth: false
    }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/auth/Register.vue'),
    meta: {
      title: '注册',
      requiresAuth: false
    }
  },
  {
    path: '/',
    component: () => import('@/components/layout/Layout.vue'),
    redirect: '/dashboard',
    meta: {
      requiresAuth: true
    },
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/Index.vue'),
        meta: {
          title: '仪表盘',
          icon: 'Dashboard'
        }
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('@/views/user/Profile.vue'),
        meta: {
          title: '个人中心',
          icon: 'User'
        }
      },
      {
        path: 'admin',
        name: 'Admin',
        component: () => import('@/views/admin/Index.vue'),
        meta: {
          title: '用户管理',
          icon: 'UserFilled',
          requiresAdmin: true
        },
        children: [
          {
            path: 'users',
            name: 'UserManagement',
            component: () => import('@/views/admin/UserManagement.vue'),
            meta: {
              title: '用户列表',
              requiresAdmin: true
            }
          }
        ]
      }
    ]
  },
  {
    path: '/403',
    name: 'Forbidden',
    component: () => import('@/views/error/403.vue'),
    meta: {
      title: '权限不足'
    }
  },
  {
    path: '/404',
    name: 'NotFound',
    component: () => import('@/views/error/404.vue'),
    meta: {
      title: '页面不存在'
    }
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/404'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

// 路由守卫
router.beforeEach(async (to, from, next) => {
  NProgress.start()
  
  const userStore = useUserStore()
  const isLoggedIn = userStore.isLoggedIn
  
  // 设置页面标题
  if (to.meta.title) {
    document.title = `${to.meta.title} - Vue Admin`
  }
  
  // 不需要认证的页面
  if (!to.meta.requiresAuth) {
    // 如果已登录，访问登录页面时重定向到首页
    if (isLoggedIn && (to.path === '/login' || to.path === '/register')) {
      next('/')
    } else {
      next()
    }
    return
  }
  
  // 需要认证的页面
  if (!isLoggedIn) {
    ElMessage.warning('请先登录')
    next('/login')
    return
  }
  
  // 检查用户信息
  if (!userStore.userInfo) {
    try {
      await userStore.getUserInfo()
    } catch (error) {
      userStore.logout()
      next('/login')
      return
    }
  }
  
  // 检查管理员权限
  if (to.meta.requiresAdmin && !userStore.isAdmin) {
    ElMessage.error('权限不足')
    next('/403')
    return
  }
  
  next()
})

router.afterEach(() => {
  NProgress.done()
})

export default router