<template>
  <section class="py-24">
    <div class="mx-auto max-w-7xl px-6">
      <div class="mb-16 text-center">
        <div class="mb-3 text-xs font-semibold uppercase tracking-[0.3em] text-amber-400">为什么选择 ShadowOne</div>
        <h2 class="mb-4 text-3xl font-bold text-white md:text-4xl">用户需要的，我们都有</h2>
        <p class="text-zinc-400">专为中国用户打造的 AI API 中转平台，免梯子、免注册海外账号</p>
      </div>

      <div class="grid gap-6 md:grid-cols-2">
        <div
          v-for="feature in features"
          :key="feature.title"
          class="group rounded-3xl border border-white/6 bg-zinc-900/60 p-8 transition-all duration-300 hover:-translate-y-1 hover:border-amber-400/25 hover:bg-zinc-900 hover:shadow-xl hover:shadow-amber-400/6"
        >
          <div
            class="mb-5 flex h-12 w-12 items-center justify-center rounded-2xl transition-transform duration-300 group-hover:scale-110"
            :class="feature.iconBg"
          >
            <component :is="feature.icon" class="h-6 w-6 text-white" />
          </div>
          <h3 class="mb-3 text-xl font-bold text-white">{{ feature.title }}</h3>
          <p class="leading-relaxed text-zinc-400">{{ feature.desc }}</p>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { h } from 'vue'

const IconShield = () => h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': '1.5' }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M9 12.75L11.25 15 15 9.75m-3-7.036A11.959 11.959 0 013.598 6 11.99 11.99 0 003 9.749c0 5.592 3.824 10.29 9 11.623 5.176-1.332 9-6.03 9-11.622 0-1.31-.21-2.571-.598-3.751h-.152c-3.196 0-6.1-1.248-8.25-3.285z' }),
])

const IconKey = () => h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': '1.5' }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M15.75 5.25a3 3 0 013 3m3 0a6 6 0 01-7.029 5.912c-.563-.097-1.159.026-1.563.43L10.5 17.25H8.25v2.25H6v2.25H2.25v-2.818c0-.597.237-1.17.659-1.591l6.499-6.499c.404-.404.527-1 .43-1.563A6 6 0 1121.75 8.25z' }),
])

const IconCoin = () => h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': '1.5' }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M12 6v12m-3-2.818l.879.659c1.171.879 3.07.879 4.242 0 1.172-.879 1.172-2.303 0-3.182C13.536 12.219 12.768 12 12 12c-.725 0-1.45-.22-2.003-.659-1.106-.879-1.106-2.303 0-3.182s2.9-.879 4.006 0l.415.33M21 12a9 9 0 11-18 0 9 9 0 0118 0z' }),
])

const IconChart = () => h('svg', { fill: 'none', viewBox: '0 0 24 24', stroke: 'currentColor', 'stroke-width': '1.5' }, [
  h('path', { 'stroke-linecap': 'round', 'stroke-linejoin': 'round', d: 'M3 13.125C3 12.504 3.504 12 4.125 12h2.25c.621 0 1.125.504 1.125 1.125v6.75C7.5 20.496 6.996 21 6.375 21h-2.25A1.125 1.125 0 013 19.875v-6.75zM9.75 8.625c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125v11.25c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V8.625zM16.5 4.125c0-.621.504-1.125 1.125-1.125h2.25C20.496 3 21 3.504 21 4.125v15.75c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V4.125z' }),
])

const features = [
  {
    icon: IconShield,
    iconBg: 'bg-gradient-to-br from-amber-500 to-amber-700',
    title: '满血不掺水',
    desc: '官方企业级通道直连，拒绝逆向阉割和以次充好。内置一键 Key 测试功能，模型真假立验。'
  },
  {
    icon: IconKey,
    iconBg: 'bg-gradient-to-br from-sky-500 to-sky-700',
    title: '一键接入',
    desc: '一个 API 密钥调用所有模型。支持 Claude Code、Codex、OpenClaw、Cursor 等主流工具。'
  },
  {
    icon: IconCoin,
    iconBg: 'bg-gradient-to-br from-green-500 to-green-700',
    title: '极致性价比',
    desc: '1 元 = 1 刀额度，约官方 1.4 折。企业级采购渠道让规模化红利直接回到开发者手里。'
  },
  {
    icon: IconChart,
    iconBg: 'bg-gradient-to-br from-violet-500 to-violet-700',
    title: '稳定透明',
    desc: '多线路智能调度，单线故障秒级切换。每笔调用可追溯，费用明细一目了然。'
  },
]
</script>
