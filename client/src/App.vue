<template>
  <div id="app">
    <md-toolbar v-if="checkAuth()">
      <md-button class="md-icon-button" @click.native="toggleLeftSidenav">
        <md-icon>menu</md-icon>
      </md-button>
      <h2 class="md-title" style="flex: 1">littleURL</h2>
      <md-button class="md-raised md-warn">{{loginStatus}}</md-button>
    </md-toolbar>
  
    <md-sidenav class="md-left md-swipeable" ref="leftSidenav" @open="open('Left')" @close="close('Left')">
      <md-toolbar class="md-large">
        <div class="md-toolbar-container">
          <h3 class="md-title">Navigation</h3>
        </div>
  
        <div class="main-sidebar-links">
          <md-list class="md-dense">
            <md-list-item>
              <router-link exact to="/home">Home</router-link>
            </md-list-item>
  
          </md-list>
        </div>
  
      </md-toolbar>
    </md-sidenav>
    <md-snackbar :md-position="vertical + ' ' + horizontal" ref="snackbar" :md-duration="duration">
      <span>{{err}}</span>
      <md-button class="md-accent" md-theme="light-blue" @click.native="$refs.snackbar.close()">OK</md-button>
  
    </md-snackbar>
  
    <router-view></router-view>
  </div>
</template>

<script>
import auth from './auth'
export default {
  data() {
    return {
      vertical: 'bottom',
      horizontal: 'center',
      duration: 4000,
      user: auth.user,
      err: '',
      contacts: [],
      loginStatus: 'Login',
      urlform: ''
    }
  },
  methods: {
    toggleLeftSidenav() {
      this.$refs.leftSidenav.toggle()
    },
    closeLeftSidenav() {
      this.$refs.leftSidenav.close()
    },
    open(ref) {
      console.log('Opened: ' + ref)
    },
    close(ref) {
      console.log('Closed: ' + ref)
    },
    errorSnackBar(errormsg) {
      this.err = errormsg
      this.$refs.snackbar.open()
    },
    postJson(object, apiUrl, redirect) {
      return this.$http.post(apiUrl, object).then(response => {
        // var jsonResponse = JSON.parse(response.body)

        // console.log(auth.getAuthHeader())
        // return jsonResponse
        return response.body
      }, {
        headers: {'Authorization': 'okokok'}
      }).catch(e => {
        console.log(auth.getAuthHeader())
        return e
      })
    },
    checkAuth() {
      return auth.isAuthenticated()
    }
  }

}
</script>

<style lang="scss">
@import 'stylesheets/variables.scss';

$sizebar-size: 280px;

[v-cloak] {
  display: none;
}
</style>
