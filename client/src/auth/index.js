import router from '../router'
const API = '/api/'
const USER_API = API + 'user/'
const LOGIN_API = USER_API + 'login'
const SIGNUP_API = USER_API + 'create/'

export default {
  user: {
    authenticated: false
  },

  Login(context, creds, redirect) {
    return context.$http.post(LOGIN_API, creds).then(response => {
      localStorage.setItem('id_token', response.body.id_token)
      localStorage.setItem('access_token', response.body.access_token)
      this.user.authenticated = true

      if (redirect) {
        router.push(redirect)
      }
    }, response => {
      return response.body
    }).catch(e => {
      return e.message
    })
  },

  Signup(context, creds, redirect) {
    context.$http.post(SIGNUP_API, creds, (data) => {
      localStorage.setItem('id_token', data.id_token)
      localStorage.setItem('access_token', data.access_token)
      this.user.authenticated = true

      if (redirect) {
        context.$router.go(redirect)
      }
    }).error((err) => {
      context.error = err
    })
  },

  isAuthenticated() {
    var jwt = localStorage.getItem('id_token')
    if (jwt) {
      return true
    }
    return false
  },

  Logout() {
    localStorage.removeItem('id_token')
    localStorage.removeItem('access_token')
    this.user.authenticated = false
  },

  getAuthHeader() {
    return {
      'Authorization': 'Bearer ' + localStorage.getItem('access_token')
    }
  }

}
