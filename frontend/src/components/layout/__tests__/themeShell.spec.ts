import { describe, expect, it } from 'vitest'
import appLayoutSource from '@/components/layout/AppLayout.vue?raw'
import authLayoutSource from '@/components/layout/AuthLayout.vue?raw'

describe('shared theme shell', () => {
  it('reuses the shared loading-theme shell in auth and app layouts', () => {
    expect(authLayoutSource).toContain('theme-shell')
    expect(authLayoutSource).toContain('theme-shell-backdrop')
    expect(appLayoutSource).toContain('theme-shell')
    expect(appLayoutSource).toContain('theme-shell-backdrop')
  })
})
