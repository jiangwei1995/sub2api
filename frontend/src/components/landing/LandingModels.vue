<template>
  <section id="models" class="py-24">
    <div class="mx-auto max-w-7xl px-6">
      <div class="mb-16 text-center">
        <div class="mb-3 inline-flex items-center gap-2 text-xs font-semibold uppercase tracking-[0.3em] text-amber-400">
          <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9.813 15.904L9 18.75l-.813-2.846a4.5 4.5 0 00-3.09-3.09L2.25 12l2.846-.813a4.5 4.5 0 003.09-3.09L9 5.25l.813 2.846a4.5 4.5 0 003.09 3.09L15.75 12l-2.846.813a4.5 4.5 0 00-3.09 3.09z" />
          </svg>
          支持模型
        </div>
        <h2 class="mb-4 text-3xl font-bold text-white md:text-4xl">一个 Key，调用所有模型</h2>
        <p class="text-zinc-400">Claude 全系列满血支持，GPT / o3 / Codex 同步可用</p>
      </div>

      <div class="mb-10 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
        <div
          v-for="model in featuredModels"
          :key="model.name"
          class="group rounded-3xl border border-white/6 bg-zinc-900/55 p-6 transition-all duration-300 hover:-translate-y-1 hover:border-amber-400/25 hover:bg-zinc-900 hover:shadow-lg hover:shadow-amber-400/6"
        >
          <div class="mb-4 flex items-center justify-between">
            <span class="rounded-full bg-zinc-800 px-2.5 py-1 text-xs font-medium text-zinc-400">
              {{ model.provider }}
            </span>
            <span class="text-xs text-zinc-600">{{ model.series }}</span>
          </div>

          <h3 class="mb-2 font-bold text-white">{{ model.name }}</h3>
          <p class="mb-4 text-xs leading-relaxed text-zinc-500">{{ model.desc }}</p>

          <div class="mb-4 flex flex-wrap gap-1.5">
            <span
              v-for="tag in model.tags"
              :key="tag"
              class="rounded bg-amber-400/10 px-2 py-0.5 text-xs font-medium text-amber-400"
            >
              {{ tag }}
            </span>
          </div>

          <div class="flex gap-4 border-t border-white/5 pt-4">
            <div>
              <div class="text-xs text-zinc-500">输入</div>
              <div class="text-sm font-semibold text-white">{{ model.inputPrice }}</div>
            </div>
            <div>
              <div class="text-xs text-zinc-500">输出</div>
              <div class="text-sm font-semibold text-white">{{ model.outputPrice }}</div>
            </div>
          </div>
        </div>
      </div>

      <div class="flex flex-wrap items-center justify-center gap-2">
        <span class="text-sm text-zinc-500">同时支持</span>
        <span
          v-for="model in moreModels"
          :key="model"
          class="rounded-full border border-white/10 bg-zinc-900 px-3 py-1 text-xs text-zinc-400"
        >
          {{ model }}
        </span>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
const featuredModels = [
  {
    provider: 'Anthropic',
    series: '旗舰推理',
    name: 'Claude Opus 4.6',
    desc: '最强推理模型，适合复杂分析、架构设计与高价值工程任务。',
    tags: ['推理', '编码', '最强'],
    inputPrice: '$5/MTok',
    outputPrice: '$25/MTok',
  },
  {
    provider: 'Anthropic',
    series: '主力编码',
    name: 'Claude Sonnet 4.6',
    desc: '日常主力编码模型，速度与质量的平衡最适合长期开发。',
    tags: ['编码', '高效', '主力'],
    inputPrice: '$3/MTok',
    outputPrice: '$15/MTok',
  },
  {
    provider: 'Anthropic',
    series: '轻量极速',
    name: 'Claude Haiku 4.5',
    desc: '最快模型，适合轻量补全、批量处理和快速迭代。',
    tags: ['快速', '轻量'],
    inputPrice: '$0.8/MTok',
    outputPrice: '$4/MTok',
  },
  {
    provider: 'OpenAI',
    series: '通用旗舰',
    name: 'GPT-5.4',
    desc: 'OpenAI 旗舰模型，兼顾编码、推理与多模态协作。',
    tags: ['推理', '编码', '多模态'],
    inputPrice: '$5/MTok',
    outputPrice: '$15/MTok',
  },
]

const moreModels = ['GPT-5', 'GPT-4o', 'o3', 'o3-pro', 'o4-mini', 'Codex Mini', '更多持续接入...']
</script>
