<template>
  <section id="pricing" class="py-24">
    <div class="mx-auto max-w-7xl px-6">
      <div class="mb-16 text-center">
        <div class="mb-3 text-xs font-semibold uppercase tracking-[0.3em] text-amber-400">定价方案</div>
        <h2 class="mb-4 text-3xl font-bold text-white md:text-4xl">按需付费，灵活选择</h2>
        <p class="text-zinc-400">按量付费永不过期，月卡用户每天享受大额度刷新</p>
      </div>

      <div class="mb-10 flex justify-center">
        <div class="inline-flex rounded-full border border-white/10 bg-zinc-900 p-1">
          <button
            v-for="tab in tabs"
            :key="tab.key"
            @click="activeTab = tab.key"
            :class="[
              'rounded-full px-5 py-2 text-sm font-medium transition-all',
              activeTab === tab.key ? 'bg-amber-400 text-black' : 'text-zinc-400 hover:text-white',
            ]"
          >
            {{ tab.label }}
          </button>
        </div>
      </div>

      <div v-if="activeTab === 'paygo'" class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <div
          v-for="plan in paygoPlans"
          :key="plan.name"
          :class="[
            'relative rounded-3xl border p-6 transition-all duration-300',
            plan.popular
              ? 'border-amber-400/40 bg-amber-400/5 shadow-lg shadow-amber-400/10'
              : 'border-white/6 bg-zinc-900/50 hover:-translate-y-1 hover:border-white/10 hover:bg-zinc-900',
          ]"
        >
          <div v-if="plan.popular" class="absolute -top-3 left-1/2 -translate-x-1/2">
            <span class="rounded-full bg-amber-400 px-3 py-1 text-xs font-bold text-black">热门</span>
          </div>

          <div class="mb-1 text-xs font-semibold uppercase tracking-[0.2em] text-zinc-500">{{ plan.name }}</div>
          <div class="mb-1 text-3xl font-black text-white">¥{{ plan.price }}</div>
          <div class="mb-4 text-sm text-zinc-500">$ {{ plan.price }} USD 额度</div>

          <div class="mb-5 space-y-2">
            <div class="flex items-center gap-2 text-sm text-zinc-300"><span class="text-amber-400">✓</span>${{ plan.price }} 额度</div>
            <div class="flex items-center gap-2 text-sm text-zinc-300"><span class="text-amber-400">✓</span>永不过期</div>
            <div class="flex items-center gap-2 text-sm text-zinc-300"><span class="text-amber-400">✓</span>支持全部模型</div>
            <div v-if="plan.priority" class="flex items-center gap-2 text-sm text-zinc-300"><span class="text-amber-400">✓</span>优先支持</div>
          </div>

          <router-link
            :to="purchasePath"
            :class="[
              'block w-full rounded-xl py-2.5 text-center text-sm font-semibold transition-colors',
              plan.popular
                ? 'bg-amber-400 text-black hover:bg-amber-300'
                : 'border border-white/10 text-white hover:border-white/20 hover:bg-white/5',
            ]"
          >
            立即充值
          </router-link>
        </div>
      </div>

      <div v-else class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <div
          v-for="plan in monthlyPlans"
          :key="plan.name"
          :class="[
            'relative rounded-3xl border p-6 transition-all duration-300',
            plan.popular
              ? 'border-amber-400/40 bg-amber-400/5 shadow-lg shadow-amber-400/10'
              : 'border-white/6 bg-zinc-900/50 hover:-translate-y-1 hover:border-white/10 hover:bg-zinc-900',
          ]"
        >
          <div v-if="plan.popular" class="absolute -top-3 left-1/2 -translate-x-1/2">
            <span class="rounded-full bg-amber-400 px-3 py-1 text-xs font-bold text-black">推荐</span>
          </div>
          <div v-if="plan.badge" class="absolute right-4 top-4">
            <span class="rounded-full bg-zinc-800 px-2.5 py-1 text-xs font-medium text-zinc-400">{{ plan.badge }}</span>
          </div>

          <div class="mb-1 text-xs font-semibold uppercase tracking-[0.2em] text-zinc-500">{{ plan.name }}</div>
          <div class="mb-4 flex items-end gap-1">
            <span class="text-3xl font-black text-white">¥{{ plan.price }}</span>
            <span class="mb-1 text-zinc-500">/月</span>
          </div>

          <div class="mb-5 grid grid-cols-3 gap-2 rounded-2xl bg-black/30 p-3 text-center">
            <div>
              <div class="text-xs text-zinc-500">每日额度</div>
              <div class="text-sm font-bold text-amber-400">${{ plan.daily }}</div>
            </div>
            <div>
              <div class="text-xs text-zinc-500">月总额度</div>
              <div class="text-sm font-bold text-white">${{ plan.monthly }}</div>
            </div>
            <div>
              <div class="text-xs text-zinc-500">折合</div>
              <div class="text-sm font-bold text-white">¥{{ plan.rate }}/USD</div>
            </div>
          </div>

          <div class="mb-5 space-y-2">
            <div class="flex items-center gap-2 text-sm text-zinc-300"><span class="text-amber-400">✓</span>${{ plan.daily }}/天额度</div>
            <div class="flex items-center gap-2 text-sm text-zinc-300"><span class="text-amber-400">✓</span>每日刷新</div>
            <div class="flex items-center gap-2 text-sm text-zinc-300"><span class="text-amber-400">✓</span>支持全部模型</div>
          </div>

          <router-link
            :to="purchasePath"
            :class="[
              'block w-full rounded-xl py-2.5 text-center text-sm font-semibold transition-colors',
              plan.popular
                ? 'bg-amber-400 text-black hover:bg-amber-300'
                : 'border border-white/10 text-white hover:border-white/20 hover:bg-white/5',
            ]"
          >
            立即订阅
          </router-link>
        </div>
      </div>

      <div class="mt-8 text-center">
        <router-link
          to="/purchase"
          class="inline-flex items-center gap-1.5 text-sm text-zinc-400 transition-colors hover:text-amber-400"
        >
          查看全部套餐
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 19.5l15-15m0 0H8.25m11.25 0v11.25" />
          </svg>
        </router-link>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useAuthStore } from '@/stores'

const authStore = useAuthStore()
const activeTab = ref<'paygo' | 'monthly'>('paygo')
const purchasePath = computed(() => (authStore.isAuthenticated ? '/purchase' : '/register'))

const tabs = [
  { key: 'paygo', label: '按量付费' },
  { key: 'monthly', label: '月卡订阅' },
] as const

const paygoPlans = [
  { name: '基础', price: 50, popular: false, priority: false },
  { name: '标准', price: 100, popular: true, priority: false },
  { name: '进阶', price: 500, popular: false, priority: false },
  { name: '专业', price: 1000, popular: false, priority: false },
  { name: '企业', price: 5000, popular: false, priority: true },
]

const monthlyPlans = [
  { name: '入门版', price: 199, daily: 15, monthly: 450, rate: '0.44', popular: false, badge: '' },
  { name: '轻量版', price: 339, daily: 30, monthly: 900, rate: '0.38', popular: false, badge: '' },
  { name: '标准版', price: 499, daily: 50, monthly: 1500, rate: '0.33', popular: true, badge: '' },
  { name: '高级版', price: 1188, daily: 120, monthly: 3600, rate: '0.33', popular: false, badge: '进阶' },
  { name: '团队版', price: 1888, daily: 200, monthly: 6000, rate: '0.31', popular: false, badge: '团队' },
  { name: '商业版', price: 4688, daily: 500, monthly: 15000, rate: '0.31', popular: false, badge: '' },
]
</script>
