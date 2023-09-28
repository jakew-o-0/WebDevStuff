import { createRouter, createWebHistory } from 'vue-router'
import Base from '../components/Base.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: '/',
      component: Base
    }
  ]
})

export default router
