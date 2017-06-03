import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import Home from '@/components/Home'
import Signup from '@/components/Signup'
import auth from '../auth'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: '/',
      component: Home,
      beforeEnter: requireAuthenticated
    },

    {
      path: '/home',
      name: 'home',
      component: Home,
      beforeEnter: requireAuthenticated
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/signup',
      name: 'signup',
      component: Signup
    }
  ]
})

export default router

function requireAuthenticated (to, from, next) {
  if (auth.isAuthenticated()) {
    next()
  } else {
    router.replace('/login')
  }
}
