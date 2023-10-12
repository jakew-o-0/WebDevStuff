import { createRouter, createWebHistory } from 'vue-router'
import UserHomePage from '../views/UserHomePage.vue'
import HomePage from '../views/HomePage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/home',
      name: 'UserHomePage',
      component: UserHomePage
    },
    {
      path: '/',
      name: 'HomePage',
      component: HomePage
    }
  ]
})

export default router
