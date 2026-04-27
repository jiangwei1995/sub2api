<template>
  <section class="py-24">
    <div class="mx-auto max-w-7xl px-6">
      <div class="relative overflow-hidden rounded-[2rem] border border-amber-400/20 bg-gradient-to-br from-amber-950/45 via-zinc-900 to-zinc-950 px-8 py-16 text-center">
        <div class="pointer-events-none absolute left-1/2 top-0 h-64 w-64 -translate-x-1/2 -translate-y-1/2 rounded-full bg-amber-400/10 blur-3xl"></div>

        <h2 class="relative mb-4 text-3xl font-black text-white md:text-4xl">3 分钟开始你的 AI 编码之旅</h2>
        <p class="relative mb-2 text-zinc-400">注册账号 → 充值 → 创建 API Key → 开始使用。就是这么简单。</p>
        <p class="relative mb-8 text-xs text-zinc-600">最低 ¥20 起充 · 额度永不过期 · 支持微信 / 支付宝 / USDT</p>

        <router-link
          :to="isAuthenticated ? dashboardPath : '/register'"
          class="inline-flex items-center gap-2 rounded-full bg-amber-400 px-8 py-3.5 text-base font-bold text-black transition-all hover:scale-[1.02] hover:bg-amber-300 hover:shadow-lg hover:shadow-amber-400/30"
        >
          {{ isAuthenticated ? '进入控制台' : '立即注册' }}
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 19.5l15-15m0 0H8.25m11.25 0v11.25" />
          </svg>
        </router-link>

        <p class="relative mt-6 text-xs text-zinc-600">
          对分销合作感兴趣？注册后访问
          <router-link to="/affiliate" class="text-amber-400 transition-colors hover:text-amber-300">分销中心</router-link>
          即可成为合伙人，佣金 20%-40%
        </p>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAuthStore } from '@/stores'

const authStore = useAuthStore()
const isAuthenticated = computed(() => authStore.isAuthenticated)
const dashboardPath = computed(() => (authStore.isAdmin ? '/admin/dashboard' : '/dashboard'))
</script>
