<template>
  <v-container style="max-width:500px;margin-top:60px">
    <v-card class="mb-4">
      <v-card-title>Create First User</v-card-title>
      <v-card-text>
        <v-text-field v-model="config.username" label="New Username" prepend-icon="mdi-account" required />
        <v-text-field v-model="config.password" label="New Password" type="password" prepend-icon="mdi-lock" required />
      </v-card-text>
    </v-card>

    <v-card>
      <v-card-title>Initial Server Setup</v-card-title>
      <v-card-text>
        <v-text-field v-model="config.tinyaddress" label="Short URL Address" placeholder="https://litt.url" required />
        <v-select v-model="config.db_type" label="Database Type" :items="['Redis']" required />
        <v-text-field v-model="config.db_address" label="Database Address" required />
        <v-text-field v-model="config.db_port" label="Database Port" type="number" required />
        <p class="text-caption text-right">* indicates required</p>
      </v-card-text>
      <v-card-actions class="justify-center">
        <v-btn color="primary" variant="elevated" @click="Initialize">Start!</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script>
import axios from 'axios'
import router from '../router'

export default {
  name: 'Config',
  inject: ['showSnack'],
  data () {
    return {
      config: { username: '', password: '', group: 'admin', tinyaddress: '', db_type: 'Redis', db_address: 'redis', db_port: 6379 }
    }
  },
  methods: {
    async Initialize () {
      try {
        const response = await axios.post('/api/config', this.config)
        this.showSnack(response.data.comment)
        if (response.data.success) {
          setTimeout(() => router.push('/ui/login'), 2000)
        }
      } catch (e) {
        this.showSnack(e.response ? e.response.data.comment : e.message)
      }
    }
  }
}
</script>
