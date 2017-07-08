<template>
  <div id="home">
    <md-layout md-gutter>
      <md-layout md-column>
      <md-card>
        <md-card-header>
          <div class="md-title">
            littleURL
          </div>
          <div class="md-subhead">
            matmerr
          </div>
        </md-card-header>
        <md-card-content>
          <md-input-container>
            <label>Long URL</label>
            <md-input v-model="urlform"></md-input>
          </md-input-container>
          
          <br>
          <md-switch  v-model="showCustom" id="my-test5" name="showCustom" class="md-primary">Use Custom Short URL</md-switch>          
            <md-input-container v-if="showCustom">
              <label>Custom URL</label>
            <md-input v-model="custom"></md-input>
          </md-input-container>

          <br>

          <!--<md-button class="md-raised md-primary" @click.native="clearTable()">Clear Table</md-button>-->
          <md-button class="md-raised md-primary" @click.native="postURL()">Shorten</md-button>
        </md-card-content>
        </md-card>
    </md-layout>
    <md-layout md-column>
      <md-card>
        <md-table>
          <md-table-header>
            <md-table-row>
              <md-table-head>Short URL </md-table-head>
              <md-table-head>Success</md-table-head>
            </md-table-row>
          </md-table-header>
          <md-table-body>
            <md-table-row v-for="url in urlList" :key="url">
              <md-table-cell>{{url.newUrl}}</md-table-cell>
              <md-table-cell>{{url.success}}</md-table-cell>
            </md-table-row>
          </md-table-body>
        </md-table>
      </md-card>
    </md-layout>
  </md-layout>
  </div>
</template>

<script>

const API_ADDURL = '/api/url/add'

export default {
  name: 'app',
  data: function () {
    return {
      urlList: [],
      jsondata: '',
      urlform: '',
      showCustom: false,
      custom: ''
    }
  },
  methods: {
    postURL: function () {
      var ctx = this
      var data = Promise.resolve(this.$parent.postJson(JSON.stringify({ url: this.urlform }), API_ADDURL))
      data.then(result => {
        if (result) {
          if (result.status === 200) {
            ctx.urlList.push({newUrl: result.body.comment, success: result.body.success})
          } else if (result.status === 401) {
            ctx.$parent.errorSnackBar(result.status + ' ' + result.body.statusText)
            ctx.$parent.redirect('/login')
          } else {
            ctx.$parent.errorSnackBar(result.status + ' ' + result.body.statusText)
          }
        }
      })
    }
  }
}
</script>

<style>

#home {

  margin-top: 30px;
  margin-left: 30px;
  margin-right: 30px;
}
</style>
