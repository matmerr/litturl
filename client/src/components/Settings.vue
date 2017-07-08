<template>
  <div id="settings">
  
    <md-card>
      <md-card-header>
        <div class="md-title">
          Settings
        </div>
      </md-card-header>
      <md-card-content>
        <md-layout md-gutter>
          <md-input-container>
            <md-icon class="md-primary">perm_identity</md-icon>
            <label>Username</label>
            <md-input required v-model="settings.tinyaddress"></md-input>
          </md-input-container>
          <md-input-container md-has-password>
            <md-icon class="md-primary">lock</md-icon>
            <label>Password</label>
            <md-input required type="password" v-model="settings.tinyaddress"></md-input>
          </md-input-container>
          <md-layout md-align="end">
            <span class="md-caption">* indicates required</span>
          </md-layout>
  
        </md-layout>
        <md-layout md-align="center">
          <md-button class="md-raised md-primary" @click.native="PostSettings()">Update </md-button>
        </md-layout>
      </md-card-content>
  
    </md-card>
  
  </div>
</template>

<script>

import auth from '../auth'

export default {
  name: 'settings',
  data: function () {
    return {
      settings: {
        wordsSHA256: '',
        tinyaddress: '',
        db_type: '',
        db_address: '',
        db_port: 1
      }
    }
  },
  methods: {
    PostSettings: function () {
      var data = Promise.resolve(this.$parent.postJson(this.settings, '/api/settings', ''))
      var ctx = this
      data.then(function (result) {
        if (result) {
          ctx.$parent.errorSnackBar(result)
        }
      })
    },
    GetSettings: function () {
      var ctx = this
      var data = Promise.resolve(this.$http.get('/api/settings', {
        headers: auth.getAuthHeader()
      }).then(response => {
        return response
      }).catch(e => {
        return e
      }))
      data.then(result => {
        if (result) {
          ctx.settings = result.body
        }
      })
    }
  },
  beforeMount () {
    this.GetSettings()
  }
}
</script>

<style>
#settings {
  margin: auto;
  margin-top: 60px;
  max-width: 450px;
}
</style>
