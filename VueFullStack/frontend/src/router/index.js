import { createRouter, createWebHistory } from 'vue-router'
import Base from '../views/Base.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Base',
      component: Base
    }
  ]
})

export default router
