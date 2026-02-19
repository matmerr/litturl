import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import Home from '../components/Home.vue'

const vuetify = createVuetify({ components, directives })

const global = {
  plugins: [vuetify],
  provide: {
    postJson: async () => ({ comment: 'http://short.url/abc', success: true }),
    showSnack: () => {}
  }
}

describe('Home', () => {
  it('renders without crashing', () => {
    const wrapper = mount(Home, { global })
    expect(wrapper.exists()).toBe(true)
  })

  it('contains URL input field', () => {
    const wrapper = mount(Home, { global })
    expect(wrapper.html()).toContain('Long URL')
  })

  it('shows custom URL field when switch toggled', async () => {
    const wrapper = mount(Home, { global })
    // Set showCustom to true directly
    await wrapper.setData({ showCustom: true })
    expect(wrapper.html()).toContain('Custom URL Mapping')
  })
})
