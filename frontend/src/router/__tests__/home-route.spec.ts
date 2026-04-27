import { beforeEach, describe, expect, it, vi } from 'vitest'

const authStore = {
  isAuthenticated: false,
  isAdmin: false,
  isSimpleMode: false,
  hasPendingAuthSession: false,
  checkAuth: vi.fn(),
}

const appStore = {
  backendModeEnabled: false,
  cachedPublicSettings: null,
  siteName: 'Sub2API',
  docUrl: '',
}

const adminSettingsStore = {
  customMenuItems: [],
}

vi.mock('@/stores/auth', () => ({
  useAuthStore: () => authStore,
}))

vi.mock('@/stores/app', () => ({
  useAppStore: () => appStore,
}))

vi.mock('@/stores/adminSettings', () => ({
  useAdminSettingsStore: () => adminSettingsStore,
}))

vi.mock('@/composables/useNavigationLoading', () => ({
  useNavigationLoadingState: () => ({
    startNavigation: vi.fn(),
    endNavigation: vi.fn(),
    isLoading: { value: false },
  }),
}))

vi.mock('@/composables/useRoutePrefetch', () => ({
  useRoutePrefetch: () => ({
    triggerPrefetch: vi.fn(),
    cancelPendingPrefetch: vi.fn(),
    resetPrefetchState: vi.fn(),
  }),
}))

const { default: router } = await import('@/router')

describe('home route', () => {
  beforeEach(() => {
    authStore.checkAuth.mockReset()
  })

  it('lazy-loads the ShadowOne landing page for /home', () => {
    const homeRoute = router.getRoutes().find((route) => route.path === '/home')

    expect(homeRoute).toBeDefined()
    expect(homeRoute?.components?.default?.toString()).toContain('HomeLandingView.vue')
  })
})
