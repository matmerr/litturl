<template>
  <div id="config">
    <div id="user_config">
      <md-card>
        <md-card-header>
          <div class="md-title">
            Create First User
          </div>
        </md-card-header>
        <md-card-content>
          <md-layout md-gutter>
            <md-input-container>
              <md-icon class="md-primary">perm_identity</md-icon>
              <label>New Username</label>
              <md-input required v-model="config.username"></md-input>
            </md-input-container>
            <md-input-container md-has-password>
              <md-icon class="md-primary">lock</md-icon>
              <label>New Password</label>
              <md-input required type="password" v-model="config.password"></md-input>
            </md-input-container>
  
          </md-layout>
        </md-card-content>
      </md-card>
    </div>
    <br>
    <div id="server_config">
      <md-card>
        <md-card-header>
          <div class="md-title">
            Initial Server Setup
          </div>
        </md-card-header>
        <md-card-content>
          <md-input-container>
            <label>Short URL Address</label>
            <md-input required v-model="config.tinyaddress" placeholder="Short URL Address (ex. https://litt.url)"></md-input>
          </md-input-container>
          <md-input-container>
  
            <label for="db">Database Type</label>
            <md-select required name="db" id="db" v-model="config.db_type">
              <md-option value="Redis">Redis
              </md-option>
              <md-option value="MongoDB">MongoDB</md-option>
            </md-select>
          </md-input-container>
  
          <md-input-container>
            <label>Database Address</label>
            <md-input required v-model="config.db_address"></md-input>
          </md-input-container>
          <md-input-container>
            <label>Database Port</label>
            <md-input type="number" required v-model="config.db_port"></md-input>
          </md-input-container>
          <md-layout md-align="end">
            <span class="md-caption">* indicates required</span>
          </md-layout>
          </md-layout>
          <md-layout md-align="center">
            <md-button class="md-raised md-primary" @click.native="Initialize()">Start!</md-button>
          </md-layout>
  
        </md-card-content>
      </md-card>
    </div>
  </div>
</template>

<script>

import router from '../router'

export default {
  name: 'config',
  data: function () {
    return {
      config: {
        username: '',
        password: '',
        tinyaddress: '',
        db_type: 'Redis',
        db_address: '192.168.91.137',
        db_port: 6379
      }
    }
  },
  methods: {
    Initialize: function () {
      var ctx = this
      var data = Promise.resolve(this.SendConfig(ctx))
      data.then(function (result) {
        if (result) {
          console.log(result)
          ctx.$parent.errorSnackBar(result)
        }
      })
    },
    SendConfig: function (ctx) {
      return ctx.$http.post('/api/config', this.config).then(response => {
        console.log('config accepted')
        router.push('/ui/login')
      }, response => {
        return response.body
      }).catch(e => {
        return e.message
      })
    }
  }
}
</script>

<style>
#config {
  margin: auto;
  margin-top: 60px;
  max-width: 450px;
}
</style>

#margin-left: 240px;
#margin-right: 240px; 
