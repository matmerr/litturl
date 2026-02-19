import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css'

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'light',
    themes: {
      light: {
        colors: {
          primary: '#006064',
          secondary: '#f44336',
          error: '#f44336'
        }
      }
    }
  },
  icons: {
    defaultSet: 'mdi'
  }
})

const app = createApp(App)
app.config.globalProperties.$axios = axios
app.use(router)
app.use(vuetify)
app.mount('#app')
