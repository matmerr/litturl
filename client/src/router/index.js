import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import Home from '@/components/Home'
import Config from '@/components/Config'
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
      component: Login,
      beforeEnter: isReady
    },
    {
      path: '/config',
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
    router.replace('/login')
  }
}

function isReady (to, from, next) {
  statusRedirect(to, from, next, true, 'server ready', '/config')
}

function needConfig (to, from, next) {
  statusRedirect(to, from, next, false, 'config missing', '/login')
}

function statusRedirect (to, from, next, resBool, resComment, ifNotRedirectTo) {
  var xhr = new XMLHttpRequest()
  xhr.open('GET', '/api/status', true)
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
