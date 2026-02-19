import { describe, it, expect, beforeEach, afterEach, vi } from 'vitest'

// Mock the router to avoid circular dependency issues
vi.mock('../router', () => ({
  default: {
    push: vi.fn()
  }
}))

// Mock axios
vi.mock('axios', () => ({
  default: {
    post: vi.fn(),
    get: vi.fn()
  }
}))

describe('auth', () => {
  let auth

  beforeEach(async () => {
    localStorage.clear()
    vi.resetModules()
    // Re-import after mocks are set up
    auth = (await import('../auth/index.js')).default
  })

  afterEach(() => {
    localStorage.clear()
    vi.clearAllMocks()
  })

  it('isAuthenticated returns false when no token', () => {
    expect(auth.isAuthenticated()).toBe(false)
  })

  it('isAuthenticated returns true when token exists', () => {
    localStorage.setItem('id_token', 'test-token')
    expect(auth.isAuthenticated()).toBe(true)
  })

  it('getAuthHeader returns correct Bearer header', () => {
    localStorage.setItem('access_token', 'my-access-token')
    const header = auth.getAuthHeader()
    expect(header.Authorization).toBe('Bearer my-access-token')
  })

  it('Logout clears localStorage', () => {
    localStorage.setItem('id_token', 'tok')
    localStorage.setItem('access_token', 'acc')
    auth.Logout()
    expect(localStorage.getItem('id_token')).toBeNull()
    expect(localStorage.getItem('access_token')).toBeNull()
  })
})
