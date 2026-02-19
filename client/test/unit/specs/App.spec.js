import Vue from 'vue'
import VueMaterial from 'vue-material'

// Verify that key dependencies can be imported and used.
describe('Dependency smoke tests', () => {
  it('Vue is a valid constructor', () => {
    expect(Vue).to.be.a('function')
  })

  it('Vue.extend creates a component constructor', () => {
    const Component = Vue.extend({
      template: '<div>hello</div>'
    })
    expect(Component).to.be.a('function')
  })

  it('vue-material plugin installs without errors', () => {
    Vue.use(VueMaterial)
    expect(Vue.material).to.exist
  })

  it('Vue instance mounts and renders data', () => {
    const vm = new Vue({
      data: { message: 'litturl' },
      template: '<span>{{ message }}</span>'
    }).$mount()
    expect(vm.$el.textContent).to.equal('litturl')
    vm.$destroy()
  })
})
