import { flushPromises, mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import HomeLandingView from '@/views/HomeLandingView.vue'

const checkAuthMock = vi.fn()
const fetchPublicSettingsMock = vi.fn()

const authState = {
  isAuthenticated: false,
  isAdmin: false,
}

const appState = {
  publicSettingsLoaded: false,
  docUrl: 'https://docs.shadowone.test',
  cachedPublicSettings: {
    doc_url: 'https://docs.shadowone.test',
  },
}

vi.mock('@/stores', () => ({
  useAuthStore: () => ({
    get isAuthenticated() {
      return authState.isAuthenticated
    },
    get isAdmin() {
      return authState.isAdmin
    },
    checkAuth: checkAuthMock,
  }),
  useAppStore: () => ({
    get publicSettingsLoaded() {
      return appState.publicSettingsLoaded
    },
    get docUrl() {
      return appState.docUrl
    },
    get cachedPublicSettings() {
      return appState.cachedPublicSettings
    },
    fetchPublicSettings: fetchPublicSettingsMock,
  }),
}))

describe('HomeLandingView', () => {
  beforeEach(() => {
    checkAuthMock.mockReset()
    fetchPublicSettingsMock.mockReset()
    fetchPublicSettingsMock.mockResolvedValue(null)
    authState.isAuthenticated = false
    authState.isAdmin = false
    appState.publicSettingsLoaded = false
    appState.docUrl = 'https://docs.shadowone.test'
    appState.cachedPublicSettings = {
      doc_url: 'https://docs.shadowone.test',
    }
  })

  it('checks auth and loads public settings on mount when needed', async () => {
    mount(HomeLandingView, {
      global: {
        stubs: {
          'router-link': {
            props: ['to'],
            template: '<a :data-to="to"><slot /></a>',
          },
          transition: false,
        },
      },
    })

    await flushPromises()

    expect(checkAuthMock).toHaveBeenCalledTimes(1)
    expect(fetchPublicSettingsMock).toHaveBeenCalledTimes(1)
  })

  it('renders the ShadowOne landing sections from hero to footer', async () => {
    appState.publicSettingsLoaded = true

    const wrapper = mount(HomeLandingView, {
      global: {
        stubs: {
          'router-link': {
            props: ['to'],
            template: '<a :data-to="to"><slot /></a>',
          },
          transition: false,
        },
      },
    })

    await flushPromises()

    const text = wrapper.text()

    expect(text).toContain('ShadowOne')
    expect(text).toContain('用户需要的，我们都有')
    expect(text).toContain('一个 Key，调用所有模型')
    expect(text).toContain('兼容你喜欢的所有工具')
    expect(text).toContain('简单透明，用多少付多少')
    expect(text).toContain('按需付费，灵活选择')
    expect(text).toContain('FAQ')
    expect(text).toContain('3 分钟开始你的 AI 编码之旅')
  })
})
