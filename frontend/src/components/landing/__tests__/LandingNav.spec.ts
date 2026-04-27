import { mount } from '@vue/test-utils'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import { nextTick } from 'vue'
import LandingNav from '@/components/landing/LandingNav.vue'

const authState = {
  isAuthenticated: false,
  isAdmin: false,
}

const appState = {
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
  }),
  useAppStore: () => ({
    get docUrl() {
      return appState.docUrl
    },
    get cachedPublicSettings() {
      return appState.cachedPublicSettings
    },
  }),
}))

describe('LandingNav', () => {
  beforeEach(() => {
    authState.isAuthenticated = false
    authState.isAdmin = false
    appState.docUrl = 'https://docs.shadowone.test'
    appState.cachedPublicSettings = {
      doc_url: 'https://docs.shadowone.test',
    }
    Object.defineProperty(window, 'scrollY', {
      value: 0,
      writable: true,
      configurable: true,
    })
  })

  it('shows login/register actions and the doc link for guests', () => {
    const wrapper = mount(LandingNav, {
      global: {
        stubs: {
          'router-link': {
            props: ['to'],
            template: '<a :data-to="to"><slot /></a>',
          },
        },
      },
    })

    expect(wrapper.text()).toContain('登录')
    expect(wrapper.text()).toContain('立即注册')
    expect(wrapper.text()).toContain('文档')
    expect(wrapper.text()).not.toContain('进入控制台')
  })

  it('shows the dashboard CTA for authenticated admins and applies the scrolled style', async () => {
    authState.isAuthenticated = true
    authState.isAdmin = true

    const wrapper = mount(LandingNav, {
      global: {
        stubs: {
          'router-link': {
            props: ['to'],
            template: '<a :data-to="to"><slot /></a>',
          },
        },
      },
    })

    expect(wrapper.find('[data-to="/admin/dashboard"]').exists()).toBe(true)
    expect(wrapper.text()).not.toContain('登录')

    window.scrollY = 120
    window.dispatchEvent(new Event('scroll'))
    await nextTick()

    expect(wrapper.get('header').classes()).toContain('bg-black/90')
  })
})
