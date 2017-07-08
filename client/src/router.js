// Third external
import auth from './auth'
import VueRouter from 'vue-router';
import Vue from 'vue'
Vue.use(VueRouter)

// import components
import Login from '@/components/Login'
import Settings from '@/components/Settings'
import Home from '@/components/Home'
import Config from '@/components/Config'

// set consts for API paths
const ROOT_API = '/'
const STATUS_API = ROOT_API + 'api/status'
const CONFIG_UI = ROOT_API + 'ui/config'
const LOGIN_UI = ROOT_API + 'ui/login'

const main = [
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

let router = new VueRouter({
  mode: 'history',
  routes: main
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
  statusRedirect(to, from, next, true, 'server ready', CONFIG_UI)
}

function needConfig (to, from, next) {
  statusRedirect(to, from, next, false, 'config missing', LOGIN_UI)
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
