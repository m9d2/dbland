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
        },
        {
          path: '/setting',
          component: () => import('@/views/setting/index.vue'),
          children: [
            {
              path: '/setting',
              component: () => import('@/views/setting/base/index.vue')
            },
            {
              path: '/about',
              component: () => import('@/views/setting/about/index.vue')
            },
          ],
        },
      ],
    },
  ]
})

export default router
