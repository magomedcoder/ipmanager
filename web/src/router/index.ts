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
      path: '/login',
      name: 'LoginView',
      component: () => import('../views/login.vue')
    }
  ]
})

// router.beforeEach((to, from, next) => {
//   document.title = to.meta.title || 'IP Manager'
//   if (!isLoggedIn() && to.fullPath !== '/login') {
//     next('/login')
//     return
//   }
//
//   if (isLoggedIn() && to.fullPath == '/login') {
//     next('/')
//     return
//   }
//
//   next()
// })

export default router
