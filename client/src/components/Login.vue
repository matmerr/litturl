<template>
  <div id="login">
  
    <md-card>
      <md-card-header>
        <div class="md-title">
          littURL
        </div>
        <div class="md-subhead">
          "The Little URL Shortener"
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

import auth from '../auth'

export default {
  name: 'login',
  data: function () {
    return {
      credentials: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    Login () {
      var creds = {
        username: this.credentials.username,
        password: this.credentials.password
      }
      var data = Promise.resolve(auth.Login(this, creds, '/ui/home'))
      var ctx = this
      data.then(function (result) {
        if (result) {
          console.log(result)
          ctx.$parent.errorSnackBar(result)
        }
      })
    }
  }
}
</script>

<style>
#login {
  width: 50%;
  margin: 0 auto;
  margin-top: 60px;
  max-width: 500px;
}
</style>
