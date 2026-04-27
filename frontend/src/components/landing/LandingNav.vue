<template>
  <header
    :class="[
      'fixed inset-x-0 top-0 z-50 transition-all duration-300',
      scrolled
        ? 'border-b border-white/5 bg-black/90 backdrop-blur-md'
        : 'bg-transparent',
    ]"
  >
    <nav class="mx-auto flex max-w-7xl items-center justify-between px-6 py-4">
      <router-link to="/home" class="flex items-center gap-2.5">
        <div
          class="flex h-9 w-9 items-center justify-center rounded-xl bg-gradient-to-br from-amber-300 via-amber-400 to-amber-600 shadow-lg shadow-amber-400/20"
        >
          <span class="text-sm font-black text-black">S</span>
        </div>
        <div>
          <div class="text-base font-extrabold tracking-wide text-white">ShadowOne</div>
          <div class="text-[10px] uppercase tracking-[0.35em] text-zinc-500">API Gateway</div>
        </div>
      </router-link>

      <div class="hidden items-center gap-7 md:flex">
        <a href="#models" class="text-sm text-zinc-400 transition-colors hover:text-white">模型</a>
        <a href="#pricing" class="text-sm text-zinc-400 transition-colors hover:text-white">定价</a>
        <a
          v-if="docUrl"
          :href="docUrl"
          target="_blank"
          rel="noopener noreferrer"
          class="text-sm text-zinc-400 transition-colors hover:text-white"
        >
          文档
        </a>
        <router-link
          to="/affiliate"
          class="text-sm text-zinc-400 transition-colors hover:text-white"
        >
          分销合伙人
        </router-link>
      </div>

      <div class="flex items-center gap-2">
        <template v-if="isAuthenticated">
          <router-link
            :to="dashboardPath"
            class="inline-flex items-center gap-2 rounded-full bg-amber-400 px-4 py-2 text-sm font-bold text-black transition-all hover:bg-amber-300 hover:shadow-lg hover:shadow-amber-400/20"
          >
            进入控制台
            <svg class="h-3.5 w-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.25">
              <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 19.5l15-15m0 0H8.25m11.25 0v11.25" />
            </svg>
          </router-link>
        </template>
        <template v-else>
          <router-link
            to="/login"
            class="rounded-full border border-white/15 px-4 py-2 text-sm text-white transition-colors hover:border-white/30 hover:bg-white/5"
          >
            登录
          </router-link>
          <router-link
            to="/register"
            class="rounded-full bg-amber-400 px-4 py-2 text-sm font-bold text-black transition-all hover:bg-amber-300 hover:shadow-lg hover:shadow-amber-400/20"
          >
            立即注册
          </router-link>
        </template>
      </div>
    </nav>
  </header>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { useAppStore, useAuthStore } from '@/stores'

const authStore = useAuthStore()
const appStore = useAppStore()

const scrolled = ref(false)
const isAuthenticated = computed(() => authStore.isAuthenticated)
const dashboardPath = computed(() => (authStore.isAdmin ? '/admin/dashboard' : '/dashboard'))
const docUrl = computed(() => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || '')

function onScroll() {
  scrolled.value = window.scrollY > 80
}

onMounted(() => {
  onScroll()
  window.addEventListener('scroll', onScroll, { passive: true })
})

onUnmounted(() => {
  window.removeEventListener('scroll', onScroll)
})
</script>
