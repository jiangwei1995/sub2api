<template>
  <section class="relative flex min-h-screen items-center overflow-hidden pt-16">
    <div class="pointer-events-none absolute inset-0">
      <div class="absolute left-1/2 top-1/3 h-[38rem] w-[38rem] -translate-x-1/2 -translate-y-1/2 rounded-full bg-amber-500/7 blur-[140px]"></div>
      <div class="absolute right-[-6rem] top-1/2 h-[20rem] w-[20rem] rounded-full bg-orange-500/8 blur-[120px]"></div>
    </div>

    <div class="relative mx-auto w-full max-w-7xl px-6 py-24">
      <div class="flex flex-col items-center gap-16 lg:flex-row lg:items-center lg:gap-20">
        <div class="flex-1 text-center lg:text-left">
          <div class="mb-6 inline-flex items-center gap-2 rounded-full border border-amber-400/20 bg-amber-400/6 px-4 py-1.5">
            <span class="h-1.5 w-1.5 rounded-full bg-amber-400"></span>
            <span class="text-xs font-semibold tracking-[0.3em] text-amber-400">SHADOW ONE — AI API GATEWAY</span>
          </div>

          <h1 class="mb-4 text-5xl font-black leading-none text-white md:text-6xl lg:text-7xl">
            ShadowOne
          </h1>
          <p class="mb-3 text-xl text-zinc-300 md:text-2xl">
            稳定、实惠、开箱即用的 AI API 中转服务
          </p>
          <p class="mb-10 text-sm text-zinc-500">
            Shadow = 低调稳定幕后运行 · One = 一个 Key 调用全部模型
          </p>

          <div class="flex flex-wrap items-center justify-center gap-3 lg:justify-start">
            <router-link
              :to="isAuthenticated ? dashboardPath : '/register'"
              class="inline-flex items-center gap-2 rounded-full bg-amber-400 px-6 py-3 text-sm font-bold text-black transition-all hover:scale-[1.02] hover:bg-amber-300 hover:shadow-lg hover:shadow-amber-400/25"
            >
              {{ isAuthenticated ? '进入控制台' : '立即体验' }}
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 19.5l15-15m0 0H8.25m11.25 0v11.25" />
              </svg>
            </router-link>

            <a
              v-if="docUrl"
              :href="docUrl"
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex items-center gap-2 rounded-full border border-white/20 px-6 py-3 text-sm font-medium text-white transition-all hover:scale-[1.02] hover:border-white/40 hover:bg-white/5"
            >
              查看文档
              <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 6.042A8.967 8.967 0 006 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 016 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 016-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0018 18a8.967 8.967 0 00-6 2.292m0-14.25v14.25" />
              </svg>
            </a>
          </div>

          <div class="mt-12 flex flex-wrap items-center justify-center gap-5 lg:justify-start">
            <div class="text-center">
              <div class="text-lg font-bold text-amber-400">1 RMB = 1 USD</div>
              <div class="text-xs text-zinc-500">官方 1.4 折</div>
            </div>
            <div class="hidden h-8 w-px bg-white/10 sm:block"></div>
            <div class="text-center">
              <div class="text-lg font-bold text-amber-400">多线路智能调度</div>
              <div class="text-xs text-zinc-500">自动故障切换</div>
            </div>
            <div class="hidden h-8 w-px bg-white/10 sm:block"></div>
            <div class="text-center">
              <div class="text-lg font-bold text-amber-400">99.9% 稳定率</div>
              <div class="text-xs text-zinc-500">7×24 监控</div>
            </div>
          </div>
        </div>

        <div class="flex flex-1 justify-center lg:justify-end">
          <div class="terminal-window w-full max-w-xl">
            <div class="flex items-center gap-2 border-b border-white/5 bg-zinc-900 px-4 py-3">
              <span class="h-3 w-3 rounded-full bg-red-500"></span>
              <span class="h-3 w-3 rounded-full bg-yellow-500"></span>
              <span class="h-3 w-3 rounded-full bg-green-500"></span>
              <span class="ml-auto font-mono text-xs text-zinc-500">terminal</span>
            </div>
            <div class="bg-zinc-950 px-5 py-5 font-mono text-sm leading-loose">
              <div class="terminal-line line-1">
                <span class="text-zinc-500"># 一键调用 Claude API</span>
              </div>
              <div class="terminal-line line-2">
                <span class="text-green-400">$</span>
                <span class="ml-2 text-sky-400">curl</span>
                <span class="ml-1 text-violet-400">-X POST</span>
                <span class="ml-1 break-all text-amber-400">https://shadowone.net/v1/messages</span>
              </div>
              <div class="terminal-line line-3">
                <span class="ml-4 text-zinc-400">-H <span class="text-green-300">"Authorization: Bearer sk-xxx"</span> \</span>
              </div>
              <div class="terminal-line line-4">
                <span class="ml-4 text-zinc-400">-H <span class="text-green-300">"Content-Type: application/json"</span> \</span>
              </div>
              <div class="terminal-line line-5">
                <span class="ml-4 text-zinc-400">-d <span class="text-green-300">'{"model":"claude-sonnet-4-6","messages":[...]}'</span></span>
              </div>
              <div class="terminal-line line-6 mt-2">
                <span class="rounded bg-green-500/15 px-2 py-0.5 font-semibold text-green-400">200 OK</span>
                <span class="ml-2 text-amber-300">{"content":"Hello!"}</span>
              </div>
              <div class="terminal-line line-7 mt-1">
                <span class="text-green-400">$</span>
                <span class="terminal-cursor ml-1"></span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAppStore, useAuthStore } from '@/stores'

const authStore = useAuthStore()
const appStore = useAppStore()

const isAuthenticated = computed(() => authStore.isAuthenticated)
const dashboardPath = computed(() => (authStore.isAdmin ? '/admin/dashboard' : '/dashboard'))
const docUrl = computed(() => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || '')
</script>

<style scoped>
.terminal-window {
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  box-shadow: 0 28px 80px rgba(0, 0, 0, 0.45), 0 0 0 1px rgba(255, 255, 255, 0.04);
  transform: perspective(1100px) rotateX(2deg) rotateY(-3deg);
  transition: transform 0.4s ease;
}

.terminal-window:hover {
  transform: perspective(1100px) rotateX(0deg) rotateY(0deg) translateY(-4px);
}

.terminal-line {
  opacity: 0;
  animation: lineAppear 0.45s ease forwards;
}

.line-1 { animation-delay: 0.2s; }
.line-2 { animation-delay: 0.6s; }
.line-3 { animation-delay: 0.95s; }
.line-4 { animation-delay: 1.25s; }
.line-5 { animation-delay: 1.55s; }
.line-6 { animation-delay: 2s; }
.line-7 { animation-delay: 2.35s; }

@keyframes lineAppear {
  from {
    opacity: 0;
    transform: translateY(4px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.terminal-cursor {
  display: inline-block;
  width: 8px;
  height: 15px;
  vertical-align: middle;
  background: #22c55e;
  animation: blink 1s step-end infinite;
}

@keyframes blink {
  0%, 50% {
    opacity: 1;
  }

  51%, 100% {
    opacity: 0;
  }
}
</style>
