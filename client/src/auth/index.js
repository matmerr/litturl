import axios from 'axios'
import router from '../router'

const LOGIN_API = '/api/user/login'

export default {
  user: {
    authenticated: false
  },

  async Login (creds) {
    try {
      const response = await axios.post(LOGIN_API, creds)
      localStorage.setItem('id_token', response.data.id_token)
      localStorage.setItem('access_token', response.data.access_token)
      this.user.authenticated = true
      await router.push('/ui/home')
      return null
    } catch (e) {
      return e.response ? e.response.data : { comment: e.message }
    }
  },

  isAuthenticated () {
    return !!localStorage.getItem('id_token')
  },

  getAuthHeader () {
    return { Authorization: 'Bearer ' + localStorage.getItem('access_token') }
  },

  async GetSettings () {
    try {
      const response = await axios.get('/api/settings', { headers: this.getAuthHeader() })
      const body = response.data
      if (body && body.tinyaddress) {
        localStorage.setItem('tinyaddress', body.tinyaddress)
      }
      return body
    } catch {
      return null
    }
  },

  Logout () {
    localStorage.removeItem('id_token')
    localStorage.removeItem('access_token')
    this.user.authenticated = false
    router.push('/ui/login')
  }
}
