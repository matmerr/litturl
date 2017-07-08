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
            <label>Words Hash</label>
            <md-input disabled v-model="settings.wordsSHA256"></md-input>
          </md-input-container>
          <md-input-container>
            <md-icon class="md-primary">perm_identity</md-icon>
            <label>Database Type</label>
            <md-input disabled v-model="settings.db_type"></md-input>
          </md-input-container> 
          <md-input-container>
            <md-icon class="md-primary">perm_identity</md-icon>
            <label>{{settings.db_type}} Address</label>
            <md-input disabled v-model="settings.db_address"></md-input>
          </md-input-container>
          <md-input-container>
            <md-icon class="md-primary">perm_identity</md-icon>
            <label>{{settings.db_type}} Port</label>
            <md-input disabled v-model="settings.db_port"></md-input>
          </md-input-container>            
          <md-input-container>
            <md-icon class="md-primary">perm_identity</md-icon>
            <label>Short URL Hostname</label>
            <md-input required v-model="settings.tinyaddress"></md-input>
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
          ctx.$parent.errorSnackBar(result.comment)
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
  beforeMount: function () {
    this.GetSettings()
  }
}
</script>

<style>
#settings {
  margin: auto;
  margin-top: 60px;
  max-width: 900px;
}
</style>
