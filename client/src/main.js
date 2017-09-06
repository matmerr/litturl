// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router.js'

import VueResource from 'vue-resource'
Vue.use(VueResource)

import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.css'
Vue.use(VueMaterial)

Vue.config.productionTip = false

Vue.material.registerTheme({
  default: {
    primary: {
      color: 'cyan',
      hue: '900'
    },
    accent: 'red',
    warn: 'red',
    background: 'white'
  },
  white: {
    primary: 'white'
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {
    App
  }
})
