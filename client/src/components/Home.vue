<template>
  <div id="home">
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
          <md-button class="md-raised md-primary" @click.native="clearTable()">Clear Table</md-button>
          <md-button class="md-raised md-primary" @click.native="postJson()">Shorten</md-button>
        </md-card-content>
        <md-table>
          <md-table-header>
            <md-table-row>
              <md-table-head>Short URL </md-table-head>
              <md-table-head>Success</md-table-head>
              <md-table-head>Email</md-table-head>
            </md-table-row>
          </md-table-header>
          <md-table-body>
            <md-table-row v-for="url in urlList" :key="url">
              <md-table-cell>{{url.newUrl}}</md-table-cell>
              <md-table-cell>{{url.success}}</md-table-cell>
              <md-table-cell>{{url.creator}}</md-table-cell>
            </md-table-row>
          </md-table-body>
        </md-table>
      </md-card>
  
  
   
  </div>
</template>

<script>

const API_ADDURL = '/api/add'

export default {
  name: 'app',
  data: function () {
    return {
      urlList: [],
      jsondata: '',
      urlform: ''
    }
  },
  methods: {
    fillTable: function () {
      this.contacts.push({newUrl: this.jsondata, success: 'Merrick', creator: 'test@test.com'})
    },
    clearTable: function () {
      this.$parent.errorSnackBar('alright')
      this.contacts.splice(0, this.url_list.length)
    },

    postJson: function () {
      var data = Promise.resolve(this.$parent.postJson(JSON.stringify({ url: this.urlform }), API_ADDURL))
      var ctx = this
      data.then(function (result) {
        if (result) {
          console.log(ctx)
          console.log(result)
          // ctx.urlList.push({newUrl: result.comment, success: result.success, creator: 'test@test.com'})
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
