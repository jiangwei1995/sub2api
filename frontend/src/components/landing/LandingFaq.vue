<template>
  <section class="py-24">
    <div class="mx-auto max-w-5xl px-6">
      <div class="mb-16 text-center">
        <div class="mb-3 text-xs font-semibold uppercase tracking-[0.3em] text-amber-400">常见问题</div>
        <h2 class="mb-4 text-3xl font-bold text-white md:text-4xl">FAQ</h2>
        <p class="text-zinc-400">还有其他问题？欢迎联系我们</p>
      </div>

      <div class="grid gap-8 md:grid-cols-2">
        <div v-for="group in faqGroups" :key="group.title">
          <h3 class="mb-4 text-sm font-semibold text-zinc-400">{{ group.title }}</h3>
          <div class="space-y-2">
            <div
              v-for="item in group.items"
              :key="item.q"
              class="overflow-hidden rounded-2xl border border-white/6 bg-zinc-900/50"
            >
              <button
                @click="toggle(item.q)"
                class="flex w-full items-center justify-between px-5 py-4 text-left text-sm font-medium text-white transition-colors hover:bg-zinc-900"
              >
                <span>{{ item.q }}</span>
                <svg
                  :class="[
                    'h-4 w-4 flex-shrink-0 text-zinc-400 transition-transform duration-200',
                    openItems.has(item.q) ? 'rotate-180' : '',
                  ]"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  stroke-width="2"
                >
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
                </svg>
              </button>
              <div
                v-show="openItems.has(item.q)"
                class="border-t border-white/5 px-5 py-4 text-sm leading-relaxed text-zinc-400"
              >
                {{ item.a }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const openItems = ref(new Set<string>())

function toggle(question: string) {
  if (openItems.value.has(question)) {
    openItems.value.delete(question)
    return
  }

  openItems.value.clear()
  openItems.value.add(question)
}

const faqGroups = [
  {
    title: '基础入门',
    items: [
      {
        q: 'ShadowOne 是什么？',
        a: 'ShadowOne 是一个 AI API 中转平台，让国内开发者无需梯子、无需注册海外账号，即可通过统一接口调用 Claude、GPT 等主流大模型。'
      },
      {
        q: '国内能直接用吗？需要梯子吗？',
        a: '完全不需要。ShadowOne 在国内有专属线路，直接访问 shadowone.net 即可，无需任何代理工具。'
      },
      {
        q: '支持哪些模型？',
        a: '支持 Claude 全系列、GPT-5.4 / GPT-5 / GPT-4o、o3 / o3-pro / o4-mini、Codex Mini 等，持续接入新模型。'
      },
      {
        q: '支持哪些工具和 IDE？',
        a: '兼容 OpenAI 格式，支持 Claude Code、Codex、OpenCode、OpenClaw、Cursor、VS Code、Windsurf、CherryStudio 等主流工具。'
      },
    ],
  },
  {
    title: '定价与计费',
    items: [
      {
        q: '为什么价格这么便宜（1块钱=1刀）？',
        a: '我们通过企业级渠道批量采购，将规模化红利直接让利给开发者。1 RMB = 1 USD 额度，相当于官方汇率的约 1.4 折。'
      },
      {
        q: '1 块钱大概能用多少？',
        a: '以 Claude Sonnet 4.6 为例，输入 5000 Token + 输出 2000 Token 约 $0.045，不到 5 分钱。'
      },
      {
        q: '实际扣费怎么计算？',
        a: '按 Token 数量计费，与官方计价方式一致。输入和输出 Token 分别计算，用多少扣多少。'
      },
      {
        q: '月卡的「每日额度」怎么用？用不完能累积吗？',
        a: '每日额度在当日零点自动刷新，当日未用完的额度不会累积到次日，系统优先消耗月卡额度。'
      },
    ],
  },
  {
    title: '功能与能力',
    items: [
      {
        q: '支持图片识别吗？',
        a: '支持。Claude 和 GPT-4o 等多模态模型均支持图片输入，可通过标准 API 的 image 类型传入图片。'
      },
      {
        q: '可以创建多个 API Key 吗？',
        a: '可以。登录控制台后可在 API Keys 页面创建多个 Key，方便隔离不同项目的用量统计。'
      },
      {
        q: '支持企业用户吗？',
        a: '支持。企业用户可使用商业版或企业版月卡，也可联系我们获取企业定制方案。'
      },
    ],
  },
  {
    title: '安全与稳定',
    items: [
      {
        q: '会封号吗？安全吗？',
        a: '我们使用官方企业级通道，请求通过平台网关转发，不需要你自己注册海外账号，也不会影响现有账户。'
      },
      {
        q: '服务稳定性怎么保障？',
        a: '我们部署了多条线路智能调度，单条线路故障时秒级自动切换，并进行 7×24 小时监控。'
      },
    ],
  },
]
</script>
