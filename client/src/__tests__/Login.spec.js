import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import Login from '../components/Login.vue'

const vuetify = createVuetify({ components, directives })

const global = {
  plugins: [vuetify],
  provide: {
    showSnack: () => {}
  }
}

describe('Login', () => {
  it('renders username and password fields', () => {
    const wrapper = mount(Login, { global })
    expect(wrapper.html()).toContain('Username')
    expect(wrapper.html()).toContain('Password')
  })

  it('renders login button', () => {
    const wrapper = mount(Login, { global })
    expect(wrapper.html()).toContain('Login')
  })
})
