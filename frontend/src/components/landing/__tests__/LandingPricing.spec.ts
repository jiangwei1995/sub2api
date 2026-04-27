import { mount } from '@vue/test-utils'
import { describe, expect, it, vi } from 'vitest'
import LandingPricing from '@/components/landing/LandingPricing.vue'

vi.mock('@/stores', () => ({
  useAuthStore: () => ({
    isAuthenticated: false,
    isAdmin: false,
  }),
}))

describe('LandingPricing', () => {
  it('shows pay-as-you-go plans by default and switches to monthly plans', async () => {
    const wrapper = mount(LandingPricing, {
      global: {
        stubs: {
          'router-link': {
            props: ['to'],
            template: '<a :data-to="to"><slot /></a>',
          },
        },
      },
    })

    expect(wrapper.text()).toContain('标准')
    expect(wrapper.text()).toContain('¥100')
    expect(wrapper.text()).not.toContain('标准版')

    await wrapper.get('button:nth-of-type(2)').trigger('click')

    expect(wrapper.text()).toContain('标准版')
    expect(wrapper.text()).toContain('¥499')
    expect(wrapper.text()).toContain('$50')
  })
})
