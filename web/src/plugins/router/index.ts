import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
  history: createWebHashHistory(),
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
        }
      ],
    },
  ]
})

export default router
