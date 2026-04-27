<template>
  <section class="py-24">
    <div class="mx-auto max-w-7xl px-6">
      <div class="mb-16 text-center">
        <div class="mb-3 text-xs font-semibold uppercase tracking-[0.3em] text-amber-400">计费说明</div>
        <h2 class="mb-4 text-3xl font-bold text-white md:text-4xl">简单透明，用多少付多少</h2>
      </div>

      <div class="grid gap-8 lg:grid-cols-2">
        <div class="space-y-6">
          <div class="rounded-3xl border border-amber-400/20 bg-amber-400/5 p-6">
            <div class="mb-2 text-2xl font-black text-amber-400">1 人民币 = 1 美元额度</div>
            <p class="text-sm text-zinc-400">
              充值 100 元即获得 $100 额度。官方汇率约 7.2，相当于 1.4 折使用全部 Claude / GPT 模型。
            </p>
          </div>

          <div class="grid gap-4 sm:grid-cols-3">
            <div class="rounded-2xl border border-white/6 bg-zinc-900 p-4">
              <div class="mb-2 font-semibold text-white">按量计费</div>
              <p class="text-xs leading-relaxed text-zinc-500">每次 API 调用按实际 Token 消耗扣费，用多少扣多少，额度永不过期。</p>
            </div>
            <div class="rounded-2xl border border-white/6 bg-zinc-900 p-4">
              <div class="mb-2 font-semibold text-white">月卡更划算</div>
              <p class="text-xs leading-relaxed text-zinc-500">月卡用户每日享有固定额度，折合低至 ¥0.31/USD，额度每日零点刷新。</p>
            </div>
            <div class="rounded-2xl border border-white/6 bg-zinc-900 p-4">
              <div class="mb-2 font-semibold text-white">VIP 等级折扣</div>
              <p class="text-xs leading-relaxed text-zinc-500">累计消费自动升级 VIP，最高 VIP8 享 0.88x 倍率，消费越多越便宜。</p>
            </div>
          </div>

          <div class="rounded-2xl border border-white/6 bg-zinc-900 p-5">
            <div class="mb-3 text-sm font-semibold text-zinc-300">举个例子</div>
            <p class="mb-3 text-xs text-zinc-500">使用 Claude Sonnet 4.6 进行一次编码对话（输入 5000 Token，输出 2000 Token）：</p>
            <div class="space-y-1.5 font-mono text-xs">
              <div class="text-zinc-400">输入: 5,000 ÷ 1,000,000 × $3 = <span class="text-white">$0.015</span></div>
              <div class="text-zinc-400">输出: 2,000 ÷ 1,000,000 × $15 = <span class="text-white">$0.030</span></div>
              <div class="mt-2 border-t border-white/5 pt-2 font-semibold text-amber-400">合计: $0.045（约 ¥0.045，不到 5 分钱）</div>
            </div>
            <p class="mt-3 text-xs text-zinc-500">充值 ¥1 就能进行约 22 次这样的对话。一杯奶茶的钱够用一周。</p>
          </div>
        </div>

        <div class="overflow-hidden rounded-3xl border border-white/6 bg-zinc-900">
          <div class="border-b border-white/5 px-6 py-4">
            <div class="font-semibold text-white">支持的模型</div>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full text-sm">
              <thead>
                <tr class="border-b border-white/5">
                  <th class="px-4 py-3 text-left text-xs font-medium text-zinc-500">模型</th>
                  <th class="px-4 py-3 text-right text-xs font-medium text-zinc-500">输入</th>
                  <th class="px-4 py-3 text-right text-xs font-medium text-zinc-500">输出</th>
                  <th class="hidden px-4 py-3 text-right text-xs font-medium text-zinc-500 sm:table-cell">Cache 写</th>
                  <th class="hidden px-4 py-3 text-right text-xs font-medium text-zinc-500 sm:table-cell">Cache 读</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="(model, index) in models"
                  :key="model.name"
                  :class="[
                    'border-b border-white/5 transition-colors hover:bg-white/2',
                    index % 2 === 0 ? '' : 'bg-white/[0.01]',
                  ]"
                >
                  <td class="px-4 py-3">
                    <div class="font-medium text-white">{{ model.name }}</div>
                    <div class="text-xs text-zinc-600">{{ model.provider }}</div>
                  </td>
                  <td class="px-4 py-3 text-right text-zinc-300">{{ model.input }}</td>
                  <td class="px-4 py-3 text-right text-zinc-300">{{ model.output }}</td>
                  <td class="hidden px-4 py-3 text-right text-zinc-500 sm:table-cell">{{ model.cacheWrite }}</td>
                  <td class="hidden px-4 py-3 text-right text-zinc-500 sm:table-cell">{{ model.cacheRead }}</td>
                </tr>
              </tbody>
            </table>
          </div>
          <div class="px-6 py-3 text-xs text-zinc-600">MTok = 百万 Token。价格与官方一致。</div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
const models = [
  { name: 'Claude Opus 4.7', provider: 'Anthropic', input: '$5', output: '$25', cacheWrite: '$6.25', cacheRead: '$0.5' },
  { name: 'Claude Opus 4.6', provider: 'Anthropic', input: '$5', output: '$25', cacheWrite: '$6.25', cacheRead: '$0.5' },
  { name: 'Claude Sonnet 4.6', provider: 'Anthropic', input: '$3', output: '$15', cacheWrite: '$3.75', cacheRead: '$0.3' },
  { name: 'Claude Haiku 4.5', provider: 'Anthropic', input: '$0.8', output: '$4', cacheWrite: '$1', cacheRead: '$0.08' },
  { name: 'GPT-5.4', provider: 'OpenAI', input: '$5', output: '$15', cacheWrite: '—', cacheRead: '—' },
  { name: 'GPT-5', provider: 'OpenAI', input: '$5', output: '$15', cacheWrite: '—', cacheRead: '—' },
  { name: 'GPT-5 Mini', provider: 'OpenAI', input: '$1.5', output: '$6', cacheWrite: '—', cacheRead: '—' },
  { name: 'GPT-4o', provider: 'OpenAI', input: '$2.5', output: '$10', cacheWrite: '—', cacheRead: '—' },
  { name: 'o3', provider: 'OpenAI', input: '$10', output: '$40', cacheWrite: '—', cacheRead: '—' },
  { name: 'o3-pro', provider: 'OpenAI', input: '$20', output: '$80', cacheWrite: '—', cacheRead: '—' },
  { name: 'o4-mini', provider: 'OpenAI', input: '$1.1', output: '$4.4', cacheWrite: '—', cacheRead: '—' },
  { name: 'Codex Mini', provider: 'OpenAI', input: '$1.5', output: '$6', cacheWrite: '—', cacheRead: '—' },
]
</script>
