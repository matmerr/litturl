import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import Settings from '@/components/Settings'
import Home from '@/components/Home'
import Config from '@/components/Config'
import auth from '../auth'

const ROOT_API = '/'
const STATUS_API = ROOT_API + 'api/status'
const CONFIG_API = ROOT_API + 'api/config'
const LOGIN_API = ROOT_API + 'api/user/login'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/ui',
      name: '/ui',
      component: Home,
      beforeEnter: requireAuthenticated
    },

    {
      path: '/ui/home',
      name: 'home',
      component: Home,
      beforeEnter: requireAuthenticated
    },
    {
      path: '/ui/settings',
      name: 'settings',
      component: Settings,
      beforeEnter: requireAuthenticated
    },
    {
      path: '/ui/login',
      name: 'login',
      component: Login,
      beforeEnter: isReady
    },
    {
      path: '/ui/config',
      name: 'config',
      component: Config,
      beforeEnter: needConfig
    }
  ]
})

export default router

function requireAuthenticated (to, from, next) {
  if (auth.isAuthenticated()) {
    next()
  } else {
    router.replace('/ui/login')
  }
}

function isReady (to, from, next) {
  statusRedirect(to, from, next, true, 'server ready', CONFIG_API)
}

function needConfig (to, from, next) {
  statusRedirect(to, from, next, false, 'config missing', LOGIN_API)
}

function statusRedirect (to, from, next, resBool, resComment, ifNotRedirectTo) {
  var xhr = new XMLHttpRequest()
  xhr.open('GET', STATUS_API, true)
  xhr.onreadystatechange = function (e) {
    if (xhr.readyState === 4) {
      var response = JSON.parse(xhr.responseText)
      if (response.ready === resBool && response.comment === resComment) {
        next()
      } else {
        router.replace(ifNotRedirectTo)
      }
    }
  }
  xhr.send()
}
