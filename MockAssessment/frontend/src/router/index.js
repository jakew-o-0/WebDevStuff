import { createRouter, createWebHistory } from 'vue-router'
import IndexPage from '../views/IndexPage.vue'
import LoginPage from '../views/LoginPage.vue'
import SignUpPage from '../views/SignUpPage.vue'
import StudentHomePage from '../views/StudentHomePage.vue'
import TeacherHomePage from '../views/TeacherHomePage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'indexpage',
      component: IndexPage
    },
    {
      path: '/login',
      name: 'LoginPage',
      component: LoginPage
    },
    {
      path: '/login/new',
      name: 'SignUpPage',
      component: SignUpPage 
    },
    {
      path: '/student/home',
      name: 'StudentHomePage',
      component: StudentHomePage 
    },
    {
      path: '/teacher/home',
      name: 'TeacherHomePage',
      component: TeacherHomePage
    },
  ]
})

export default router
