<template>
  <v-app>
    <v-app-bar v-if="isAuthenticated" color="primary" density="compact">
      <v-app-bar-nav-icon @click="drawer = !drawer" />
      <v-app-bar-title>littURL</v-app-bar-title>
      <v-btn variant="tonal" color="error" @click="logout">Logout</v-btn>
    </v-app-bar>

    <v-navigation-drawer v-if="isAuthenticated" v-model="drawer" :rail="false">
      <v-list>
        <v-list-item title="Home" prepend-icon="mdi-home" :to="'/ui/home'" />
        <v-list-item title="Settings" prepend-icon="mdi-cog" :to="'/ui/settings'" />
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <router-view />
    </v-main>

    <v-snackbar v-model="snackbar" :timeout="4000" location="bottom center">
      {{ snackMessage }}
      <template #actions>
        <v-btn color="secondary" variant="text" @click="snackbar = false">OK</v-btn>
      </template>
    </v-snackbar>
  </v-app>
</template>

<script>
import auth from './auth'
import axios from 'axios'

export default {
  name: 'App',
  provide () {
    return {
      showSnack: (msg) => this.errorSnackBar(msg),
      postJson: (obj, url) => this.postJson(obj, url)
    }
  },
  data () {
    return {
      drawer: true,
      snackbar: false,
      snackMessage: ''
    }
  },
  computed: {
    isAuthenticated () {
      return auth.isAuthenticated()
    }
  },
  methods: {
    logout () {
      auth.Logout()
    },
    errorSnackBar (msg) {
      this.snackMessage = msg
      this.snackbar = true
    },
    async postJson (object, apiUrl) {
      try {
        const response = await axios.post(apiUrl, object, {
          headers: auth.getAuthHeader()
        })
        return response.data
      } catch (e) {
        return e.response ? e.response.data : { comment: e.message, success: false }
      }
    }
  }
}
</script>
