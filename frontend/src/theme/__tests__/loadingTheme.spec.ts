import { describe, expect, it } from 'vitest'
import tailwindConfig from '../../../tailwind.config.js'

describe('loading theme tokens', () => {
  it('uses the refreshed loading-style primary palette and gradients', () => {
    const theme = tailwindConfig.theme.extend

    expect(theme.colors.primary[500]).toBe('#12bea9')
    expect(theme.colors.primary[700]).toBe('#0f7d76')
    expect(theme.backgroundImage['gradient-primary']).toBe(
      'linear-gradient(135deg, #12bea9 0%, #0a9b8f 100%)'
    )
    expect(theme.backgroundImage['mesh-gradient']).toContain('rgba(18, 190, 169, 0.16)')
    expect(theme.boxShadow.glow).toBe('0 0 20px rgba(18, 190, 169, 0.28)')
  })
})
