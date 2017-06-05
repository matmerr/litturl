<template>
  <div id="login">
  
    <md-card>
      <md-card-header>
        <div class="md-title">
          login
        </div>
      </md-card-header>
      <md-card-content>
        <md-layout md-gutter>
          <md-input-container>
            <md-icon class="md-primary">perm_identity</md-icon>
            <label>Username</label>
            <md-input required v-model="credentials.username"></md-input>
          </md-input-container>
          <md-input-container md-has-password>
            <md-icon class="md-primary">lock</md-icon>
            <label>Password</label>
            <md-input required type="password" v-model="credentials.password"></md-input>
          </md-input-container>
          <md-layout md-align="end">
            <span class="md-caption">* indicates required</span>
          </md-layout>
  
        </md-layout>
        <md-layout md-align="center">
          <md-button class="md-raised md-primary" @click.native="Login()">Login</md-button>
        </md-layout>
      </md-card-content>
  
    </md-card>
  
  </div>
</template>

<script>
// import Hello from './components/Hello'

import auth from '../auth'

export default {
  name: 'login',
  data: function() {
    return {
      credentials: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    Login: function() {
      var creds = {
        username: this.credentials.username,
        password: this.credentials.password
      }
      var data = Promise.resolve(auth.Login(this, creds, '/home'))
      var ctx = this
      data.then(function(result) {
        if (result) {
          ctx.$parent.errorSnackBar(result)
        }
      })
    }
  }
}
</script>

<style>
#login {
  margin: auto;
  margin-top: 60px;
  max-width: 450px;
}
</style>

#margin-left: 240px;
#margin-right: 240px; 
