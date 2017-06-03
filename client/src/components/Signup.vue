<template>
  <div id="config">

    <md-card>
      <md-card-header>
        <div class="md-title">
          Create Admin User
        </div>
      </md-card-header>
      <md-card-content>
       
       <md-layout md-gutter>
      <md-input-container>
        <md-icon class="md-primary">perm_identity</md-icon>
          <label>Username</label>
        <md-input required ></md-input>
      </md-input-container>
      <md-input-container md-has-password>
        <md-icon class="md-primary">lock</md-icon>
          <label>Password</label>
        <md-input required type="password"></md-input>
      </md-input-container>
      <md-layout md-align="end">
        <span class="md-caption">* indicates required</span>
      </md-layout>

      </md-layout>
      </md-card-content>
      
    </md-card>
    
     
    

  </div>
</template>

<script>
// import Hello from './components/Hello'

export default {
  data: function () {
    return {
      contacts: [],
      urlform: ''
    }
  },
  methods: {
    postJson: function () {
      var xmlhttp = new XMLHttpRequest()   // new HttpRequest instance
      xmlhttp.open('POST', '/add')
      xmlhttp.setRequestHeader('Content-Type', 'application/json;charset=UTF-8')

      var cxt = this
      xmlhttp.onreadystatechange = function () {
        if (xmlhttp.readyState === 4 && xmlhttp.status === 200) {
          var json = JSON.parse(xmlhttp.responseText)
          console.log(json.comment + ', ' + json.success)
          cxt.contacts.push({firstname: json.comment, lastname: json.success, email: 'test@test.com'})
        }
      }
      xmlhttp.send(JSON.stringify({ apikey: 'lol', url: cxt.urlform }))
    }

  }
}
</script>

<style>
#config {
  margin-top: 60px;
  margin-left: 240px;
  margin-right: 240px;
}
</style>
