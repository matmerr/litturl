<template>
  <div id="home">
    <md-layout md-gutter md-column>
      <md-layout md-column>
        <md-card>
          <md-card-header>
            <div class="md-title">
              Home
            </div>
          </md-card-header>
          <md-card-content>
            <md-input-container>
              <label>Long URL</label>
              <md-input v-model="urlform"></md-input>
            </md-input-container>
            <md-switch v-model="showCustom" name="showCustom" class="md-primary">Use Custom Short URL</md-switch>
            <md-layout v-if="showCustom">
              <md-layout md-flex="100">
                <md-input-container>
                  <label>Custom URL Mapping</label>
                  <md-input v-model="custom"></md-input>
                </md-input-container>
                Result: {{tinyaddress}}{{custom}}
              </md-layout>
            </md-layout>
            <md-layout>
            </md-layout>
            <md-layout md-gutter>
              <md-button class="md-raised md-primary" @click.native="postURL()">Shorten</md-button>
            </md-layout>
          </md-card-content>
        </md-card>
      </md-layout>
      <br>
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

  name: 'home',
  data: function () {
    return {
      urlList: [],
      jsondata: '',
      urlform: '',
      showCustom: false,
      custom: '',
      tinyaddress: localStorage.getItem('tinyaddress')
    }
  },
  methods: {
    postURL: function () {
      var ctx = this
      var data = Promise.resolve(this.$parent.postJson(JSON.stringify({ url: this.urlform, custom: this.custom }), API_ADDURL))
      data.then(result => {
        if (result) {
          ctx.urlList.push({ newUrl: result.comment, success: result.success })
        }
      })
    }
  }
}
</script>

<style>
#home {
  overflow: auto;
  margin-top: 30px;
  margin-left: 30px;
  margin-right: 30px;
}
</style>
