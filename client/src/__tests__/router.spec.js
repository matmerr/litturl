import { describe, it, expect } from 'vitest'
import router from '../router'

describe('router', () => {
  it('has correct routes', () => {
    const routes = router.getRoutes()
    const paths = routes.map(r => r.path)
    expect(paths).toContain('/ui')
    expect(paths).toContain('/ui/home')
    expect(paths).toContain('/ui/settings')
    expect(paths).toContain('/ui/login')
    expect(paths).toContain('/ui/config')
  })
})
