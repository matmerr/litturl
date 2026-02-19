import { createRouter, createWebHistory } from 'vue-router'
import auth from './auth'
import Login from '@/components/Login.vue'
import Settings from '@/components/Settings.vue'
import Home from '@/components/Home.vue'
import Config from '@/components/Config.vue'

const STATUS_API = '/api/status'
const CONFIG_UI = '/ui/config'
const LOGIN_UI = '/ui/login'

async function getStatus () {
  try {
    const res = await fetch(STATUS_API)
    return await res.json()
  } catch {
    return null
  }
}

async function requireAuthenticated (_to, _from, next) {
  if (auth.isAuthenticated()) {
    next()
  } else {
    next(LOGIN_UI)
  }
}

async function isReady (_to, _from, next) {
  const status = await getStatus()
  if (status && status.ready === true && status.comment === 'server ready') {
    next()
  } else {
    next(CONFIG_UI)
  }
}

async function needConfig (_to, _from, next) {
  const status = await getStatus()
  if (status && status.ready === false && status.comment === 'config missing') {
    next()
  } else {
    next(LOGIN_UI)
  }
}

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/ui', name: 'ui', component: Home, beforeEnter: requireAuthenticated },
    { path: '/ui/home', name: 'home', component: Home, beforeEnter: requireAuthenticated },
    { path: '/ui/settings', name: 'settings', component: Settings, beforeEnter: requireAuthenticated },
    { path: '/ui/login', name: 'login', component: Login, beforeEnter: isReady },
    { path: '/ui/config', name: 'config', component: Config, beforeEnter: needConfig },
    { path: '/:pathMatch(.*)*', redirect: '/ui' }
  ]
})

export default router
