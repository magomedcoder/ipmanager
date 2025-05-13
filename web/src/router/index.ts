import { createRouter, createWebHistory } from 'vue-router'
import { isLoggedIn } from '@/utils/cookie'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'HomeView',
      component: () => import('../views/home.vue')
    },
    {
      path: '/services',
      name: 'servicesView',
      component: () => import('../views/services.vue'),
      meta: {
        title: 'Сервисы'
      }
    },
    {
      path: '/vlans',
      name: 'vlanView',
      component: () => import('../views/vlans.vue'),
      meta: {
        title: 'VLAN'
      }
    },
    {
      path: '/customers',
      name: 'CustomerView',
      component: () => import('../views/customers.vue'),
      meta: {
        title: 'Клиенты'
      }
    },
    {
      path: '/users',
      name: 'UserView',
      component: () => import('../views/users.vue'),
      meta: {
        title: 'Пользователи'
      }
    },
    {
      path: '/login',
      name: 'LoginView',
      component: () => import('../views/login.vue'),
      meta: {
        title: 'Вход'
      }
    },
  ]
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title || 'IP Manager'
  if (!isLoggedIn() && to.fullPath !== '/login') {
    next('/login')
    return
  }

  if (isLoggedIn() && to.fullPath == '/login') {
    next('/')
    return
  }

  next()
})

export default router
