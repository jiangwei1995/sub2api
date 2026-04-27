import { mount } from '@vue/test-utils'
import { describe, expect, it } from 'vitest'
import LandingFaq from '@/components/landing/LandingFaq.vue'

describe('LandingFaq', () => {
  it('opens one answer at a time', async () => {
    const wrapper = mount(LandingFaq)
    const buttons = wrapper.findAll('button')
    const firstButton = buttons.find((button) => button.text().includes('ShadowOne 是什么？'))
    const secondButton = buttons.find((button) => button.text().includes('国内能直接用吗？需要梯子吗？'))

    expect(firstButton).toBeDefined()
    expect(secondButton).toBeDefined()

    const firstAnswer = firstButton!.element.parentElement?.querySelector('div.border-t') as HTMLElement
    const secondAnswer = secondButton!.element.parentElement?.querySelector('div.border-t') as HTMLElement

    expect(firstAnswer.style.display).toBe('none')
    expect(secondAnswer.style.display).toBe('none')

    await firstButton!.trigger('click')
    expect(firstAnswer.style.display).toBe('')
    expect(secondAnswer.style.display).toBe('none')

    await secondButton!.trigger('click')
    expect(firstAnswer.style.display).toBe('none')
    expect(secondAnswer.style.display).toBe('')
  })
})
