import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Login from "@/views/Login";

import store from '../store'

const isAuthenticated = (to, from, next) => {
  if (store.getters.isAuthenticated) {
    next()
    return
  }
  next('/login')
}


const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/',
    name: 'Home',
    component: Home,
    beforeEnter: isAuthenticated,
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
