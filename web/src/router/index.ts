import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/home/index.vue'),
      children: [
        {
          path: '/',
          component: () => import('@/views/database/index.vue')
        },
        {
          path: '/connect',
          component: () => import('@/views/connect/index.vue')
        },
        {
          path: '/chart',
          component: () => import('@/views/chart/index.vue')
        },
        {
          path: '/setting',
          component: () => import('@/views/setting/index.vue')
        },
      ],
    },
  ]
})

export default router
