<template>
  <v-container style="max-width:500px;margin-top:60px">
    <v-card>
      <v-card-title>littURL</v-card-title>
      <v-card-subtitle>"The Little URL Shortener"</v-card-subtitle>
      <v-card-text>
        <v-text-field
          v-model="credentials.username"
          label="Username"
          prepend-icon="mdi-account"
          required
        />
        <v-text-field
          v-model="credentials.password"
          label="Password"
          type="password"
          prepend-icon="mdi-lock"
          required
        />
        <p class="text-caption text-right">* indicates required</p>
      </v-card-text>
      <v-card-actions class="justify-center">
        <v-btn color="primary" variant="elevated" @click="doLogin">Login</v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<script>
import auth from '../auth'

export default {
  name: 'Login',
  inject: ['showSnack'],
  data () {
    return {
      credentials: { username: '', password: '', group: 'admin' }
    }
  },
  methods: {
    async doLogin () {
      const err = await auth.Login(this.credentials)
      if (err) {
        this.showSnack(err.comment || 'Login failed')
      }
    }
  }
}
</script>
