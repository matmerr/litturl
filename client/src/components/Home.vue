<template>
  <v-container class="mt-4">
    <v-row>
      <v-col cols="12">
        <v-card>
          <v-card-title>Home</v-card-title>
          <v-card-text>
            <v-text-field label="Long URL" v-model="urlform" />
            <v-switch v-model="showCustom" label="Use Custom Short URL" color="primary" />
            <template v-if="showCustom">
              <v-text-field label="Custom URL Mapping" v-model="custom" />
              <p>Result: {{ tinyaddress }}{{ custom }}</p>
            </template>
            <v-btn color="primary" @click="postURL">Shorten</v-btn>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-row class="mt-4">
      <v-col cols="12">
        <v-card>
          <v-table>
            <thead>
              <tr>
                <th>Short URL</th>
                <th>Success</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(url, i) in urlList" :key="i">
                <td>{{ url.newUrl }}</td>
                <td>{{ url.success }}</td>
              </tr>
            </tbody>
          </v-table>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
const API_ADDURL = '/api/url/add'

export default {
  name: 'Home',
  inject: ['postJson', 'showSnack'],
  data () {
    return {
      urlList: [],
      urlform: '',
      showCustom: false,
      custom: '',
      tinyaddress: localStorage.getItem('tinyaddress') || ''
    }
  },
  methods: {
    async postURL () {
      const result = await this.postJson({ url: this.urlform, custom: this.custom }, API_ADDURL)
      if (result) {
        this.urlList.push({ newUrl: result.comment, success: result.success })
      }
    }
  }
}
</script>
