import { createRouter, createWebHistory } from 'vue-router'
import IndexPage from '../views/IndexPage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'IndexPage',
      component: IndexPage
    },
  ]
})

export default router
