<template>
  <v-container class="mt-4" style="max-width:600px">
    <v-card>
      <v-card-title>Settings</v-card-title>
      <v-card-text>
        <v-text-field label="Words Hash" v-model="settings.wordsSHA256" disabled prepend-icon="mdi-key-variant" />
        <v-text-field label="Database Type" v-model="settings.db_type" disabled prepend-icon="mdi-database" />
        <v-text-field :label="settings.db_type + ' Address'" v-model="settings.db_address" disabled prepend-icon="mdi-server" />
        <v-text-field :label="settings.db_type + ' Port'" v-model="settings.db_port" disabled prepend-icon="mdi-numeric" />
        <v-text-field label="Short URL Hostname" v-model="settings.tinyaddress" required prepend-icon="mdi-link" />
        <p class="text-caption text-right">* indicates required</p>
      </v-card-text>
      <v-card-actions class="justify-center">
        <v-btn color="primary" variant="elevated" @click="PostSettings">Update</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script>
import auth from '../auth'

export default {
  name: 'Settings',
  inject: ['postJson', 'showSnack'],
  data () {
    return {
      settings: { wordsSHA256: '', tinyaddress: '', db_type: '', db_address: '', db_port: 1 }
    }
  },
  async beforeMount () {
    const s = await auth.GetSettings()
    if (s) this.settings = s
  },
  methods: {
    async PostSettings () {
      const result = await this.postJson(this.settings, '/api/settings')
      if (result) {
        this.showSnack(result.comment)
        if (this.settings.tinyaddress.slice(-1) !== '/') {
          this.settings.tinyaddress += '/'
        }
        localStorage.setItem('tinyaddress', this.settings.tinyaddress)
      }
    }
  }
}
</script>
