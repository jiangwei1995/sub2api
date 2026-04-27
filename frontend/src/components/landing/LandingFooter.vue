<template>
  <footer class="border-t border-white/5 bg-black py-16">
    <div class="mx-auto max-w-7xl px-6">
      <div class="grid gap-10 md:grid-cols-5">
        <div class="md:col-span-2">
          <router-link to="/home" class="mb-4 flex items-center gap-2.5">
            <div class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-amber-400 to-amber-600">
              <span class="text-sm font-black text-black">S</span>
            </div>
            <span class="text-lg font-bold text-white">ShadowOne</span>
          </router-link>
          <p class="text-sm leading-relaxed text-zinc-500">
            稳定、实惠、开箱即用的 AI API 中转服务<br>
            专为中国开发者打造，无需梯子
          </p>
        </div>

        <div v-for="column in linkColumns" :key="column.title">
          <h4 class="mb-4 text-xs font-semibold uppercase tracking-[0.25em] text-zinc-400">{{ column.title }}</h4>
          <ul class="space-y-2.5">
            <li v-for="link in column.links" :key="link.label">
              <a
                v-if="link.kind === 'external' || link.kind === 'anchor'"
                :href="link.href"
                :target="link.kind === 'external' ? '_blank' : undefined"
                :rel="link.kind === 'external' ? 'noopener noreferrer' : undefined"
                class="text-sm text-zinc-500 transition-colors hover:text-white"
              >
                {{ link.label }}
              </a>
              <router-link
                v-else
                :to="link.href"
                class="text-sm text-zinc-500 transition-colors hover:text-white"
              >
                {{ link.label }}
              </router-link>
            </li>
          </ul>
        </div>
      </div>

      <div class="mt-12 flex flex-col items-center justify-between gap-4 border-t border-white/5 pt-8 sm:flex-row">
        <p class="text-xs text-zinc-600">© {{ currentYear }} ShadowOne. All rights reserved.</p>
        <p class="text-xs text-zinc-700">shadowone.net</p>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAppStore } from '@/stores'

type FooterLink = {
  label: string
  href: string
  kind: 'route' | 'external' | 'anchor'
}

const appStore = useAppStore()
const currentYear = computed(() => new Date().getFullYear())
const docUrl = computed(() => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || '')

const linkColumns = computed(() => {
  const productLinks: FooterLink[] = [
    { label: '按量付费', href: '/purchase', kind: 'route' },
    { label: '月卡订阅', href: '/purchase', kind: 'route' },
    { label: '模型定价', href: '#pricing', kind: 'anchor' },
    { label: '企业方案', href: '/purchase', kind: 'route' },
  ]

  const resourceLinks: FooterLink[] = docUrl.value
    ? [
        { label: '使用文档', href: docUrl.value, kind: 'external' },
        { label: '常见问题', href: '#pricing', kind: 'anchor' },
        { label: '分销合伙人', href: '/affiliate', kind: 'route' },
      ]
    : [
        { label: '常见问题', href: '#pricing', kind: 'anchor' },
        { label: '分销合伙人', href: '/affiliate', kind: 'route' },
      ]

  const contactLinks: FooterLink[] = [
    { label: '分销合作', href: '/affiliate', kind: 'route' },
    { label: '隐私政策', href: '/privacy-policy', kind: 'route' },
    { label: '服务条款', href: '/terms-of-service', kind: 'route' },
  ]

  return [
    { title: '产品', links: productLinks },
    { title: '资源', links: resourceLinks },
    { title: '联系我们', links: contactLinks },
  ]
})
</script>
