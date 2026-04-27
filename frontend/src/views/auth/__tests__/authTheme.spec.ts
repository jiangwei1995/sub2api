import { readFileSync } from 'node:fs'
import { describe, expect, it } from 'vitest'
import authLayoutSource from '@/components/layout/AuthLayout.vue?raw'
import loginViewSource from '@/views/auth/LoginView.vue?raw'
import registerViewSource from '@/views/auth/RegisterView.vue?raw'
import linuxDoOAuthSource from '@/components/auth/LinuxDoOAuthSection.vue?raw'
import wechatOAuthSource from '@/components/auth/WechatOAuthSection.vue?raw'
import oidcOAuthSource from '@/components/auth/OidcOAuthSection.vue?raw'

const stylesSource = readFileSync(`${process.cwd()}/src/style.css`, 'utf8')

describe('auth loading theme alignment', () => {
  it('uses dedicated auth theme classes across the auth shell, forms, and oauth actions', () => {
    expect(stylesSource).toContain('.auth-panel')
    expect(stylesSource).toContain('.auth-input')
    expect(stylesSource).toContain('.auth-submit-button')
    expect(stylesSource).toContain('.auth-oauth-button')

    expect(authLayoutSource).toContain('auth-panel')

    expect(loginViewSource).toContain('auth-heading')
    expect(loginViewSource).toContain('auth-subheading')
    expect(loginViewSource).toContain('auth-input')
    expect(loginViewSource).toContain('auth-submit-button')

    expect(registerViewSource).toContain('auth-heading')
    expect(registerViewSource).toContain('auth-subheading')
    expect(registerViewSource).toContain('auth-input')
    expect(registerViewSource).toContain('auth-submit-button')

    expect(linuxDoOAuthSource).toContain('auth-oauth-button')
    expect(wechatOAuthSource).toContain('auth-oauth-button')
    expect(oidcOAuthSource).toContain('auth-oauth-button')
  })
})
