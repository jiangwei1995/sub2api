<template>
  <section class="py-24">
    <div class="mx-auto max-w-7xl px-6">
      <div class="mb-16 text-center">
        <div class="mb-3 inline-flex items-center gap-2 text-xs font-semibold uppercase tracking-[0.3em] text-amber-400">
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M11.42 15.17L17.25 21A2.652 2.652 0 0021 17.25l-5.877-5.877M11.42 15.17l2.496-3.03c.317-.384.74-.626 1.208-.766M11.42 15.17l-4.655 5.653a2.548 2.548 0 11-3.586-3.586l6.837-5.63m5.108-.233c.55-.164 1.163-.188 1.743-.14a4.5 4.5 0 004.486-6.336l-3.276 3.277a3.004 3.004 0 01-2.25-2.25l3.276-3.276a4.5 4.5 0 00-6.336 4.486c.091 1.076-.071 2.264-.904 2.95l-.102.085m-1.745 1.437L5.909 7.5H4.5L2.25 3.75l1.5-1.5L7.5 4.5v1.409l4.26 4.26m-1.745 1.437l1.745-1.437m6.615 8.206L15.75 15.75M4.867 19.125h.008v.008h-.008v-.008z" />
          </svg>
          平台支持
        </div>
        <h2 class="mb-4 text-3xl font-bold text-white md:text-4xl">兼容你喜欢的所有工具</h2>
        <p class="text-zinc-400">兼容 OpenAI 格式，几乎支持所有主流 AI 编程工具和 IDE</p>
      </div>

      <div class="grid grid-cols-2 gap-4 sm:grid-cols-3 lg:grid-cols-4">
        <div
          v-for="tool in tools"
          :key="tool.name"
          class="group flex flex-col items-center gap-3 rounded-3xl border border-white/6 bg-zinc-900/55 p-6 text-center transition-all duration-300 hover:-translate-y-1 hover:border-amber-400/25 hover:bg-zinc-900"
        >
          <div
            class="flex h-12 w-12 items-center justify-center rounded-2xl text-lg font-bold text-white transition-transform duration-300 group-hover:scale-110"
            :class="tool.iconBg"
          >
            {{ tool.letter }}
          </div>
          <div>
            <div class="font-semibold text-white">{{ tool.name }}</div>
            <div class="mt-0.5 text-xs text-zinc-500">{{ tool.desc }}</div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
const tools = [
  { name: 'Claude Code', desc: 'Anthropic 官方 CLI', letter: 'CC', iconBg: 'bg-gradient-to-br from-orange-500 to-orange-700' },
  { name: 'Codex', desc: 'OpenAI 编码 CLI', letter: 'Cx', iconBg: 'bg-gradient-to-br from-green-600 to-green-800' },
  { name: 'OpenCode', desc: '开源多模型 CLI', letter: 'OC', iconBg: 'bg-gradient-to-br from-sky-500 to-sky-700' },
  { name: 'OpenClaw', desc: '龙虾 AI 助手', letter: 'OL', iconBg: 'bg-gradient-to-br from-red-500 to-red-700' },
  { name: 'Cursor', desc: 'AI-first 编辑器', letter: 'Cu', iconBg: 'bg-gradient-to-br from-zinc-600 to-zinc-800' },
  { name: 'VS Code', desc: 'Continue / Cline', letter: 'VS', iconBg: 'bg-gradient-to-br from-blue-500 to-blue-700' },
  { name: 'Windsurf', desc: 'AI 编码编辑器', letter: 'WS', iconBg: 'bg-gradient-to-br from-teal-500 to-teal-700' },
  { name: 'CherryStudio', desc: '桌面 AI 客户端', letter: 'CS', iconBg: 'bg-gradient-to-br from-pink-500 to-pink-700' },
]
</script>
