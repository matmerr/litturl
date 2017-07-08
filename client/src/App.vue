<template>
  <div id="app">
    <md-toolbar v-if="checkAuth()">
      <md-button class="md-icon-button" @click.native="toggleLeftSidenav">
        <md-icon>menu</md-icon>
      </md-button>
      <h2 class="md-title" style="flex: 1">littleURL</h2>
      <md-button class="md-raised md-warn">{{loginStatus}}</md-button>
    </md-toolbar>
  
    <md-sidenav class="md-left md-swipeable" ref="leftSidenav">
      <md-toolbar class="md-large">
        <div class="md-toolbar-container">
          <h3 class="md-title">Navigation</h3>
        </div>
  
        <div class="main-sidebar-links">
          <md-list class="md-dense">
            <md-list-item>
              <span>
              <router-link exact to="/ui/home">Home</router-link>
              </span>
            </md-list-item>
            <md-list-item>
              <span>
              <router-link exact to="/ui/settings">Settings</router-link>
              </span>
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
  data () {
    return {
      vertical: 'bottom',
      horizontal: 'center',
      duration: 4000,
      user: auth.user,
      err: '',
      contacts: [],
      loginStatus: 'Logout',
      urlform: ''
    }
  },
  methods: {
    toggleLeftSidenav () {
      this.$refs.leftSidenav.toggle()
    },
    closeLeftSidenav () {
      this.$refs.leftSidenav.close()
    },
    errorSnackBar (errormsg) {
      this.err = errormsg
      this.$refs.snackbar.open()
    },
    redirect (target) {
      console.log(this.router)
    },
    postJson (object, apiUrl, redirect) {
      return this.$http.post(apiUrl, object, {
        headers: auth.getAuthHeader()
      }).then(response => {
        if (redirect) {
          this.router.push(redirect)
        }
        // return jsonResponse
        return response.body
      }).catch(e => {
        return e
      })
    },
    checkAuth () {
      return auth.isAuthenticated()
    }
  }

}
</script>

<style lang="scss">
$sizebar-size: 280px;
</style>
