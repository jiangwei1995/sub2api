<template>
  <div
    class="min-h-screen overflow-x-hidden bg-[#0a0a0a] text-white"
    style="font-family: Manrope, 'PingFang SC', 'Microsoft YaHei', sans-serif;"
  >
    <LandingNav />
    <main class="relative">
      <div class="pointer-events-none absolute inset-0 overflow-hidden">
        <div class="absolute left-1/2 top-0 h-[32rem] w-[32rem] -translate-x-1/2 rounded-full bg-amber-500/8 blur-[140px]"></div>
        <div class="absolute right-[-8rem] top-[26rem] h-[24rem] w-[24rem] rounded-full bg-amber-300/5 blur-[120px]"></div>
        <div class="absolute left-[-10rem] top-[68rem] h-[24rem] w-[24rem] rounded-full bg-orange-500/5 blur-[120px]"></div>
      </div>

      <LandingHero />
      <LandingFeatures />
      <LandingModels />
      <LandingTools />
      <LandingBilling />
      <LandingPricing />
      <LandingFaq />
      <LandingCta />
    </main>
    <LandingFooter />
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import LandingBilling from '@/components/landing/LandingBilling.vue'
import LandingCta from '@/components/landing/LandingCta.vue'
import LandingFaq from '@/components/landing/LandingFaq.vue'
import LandingFeatures from '@/components/landing/LandingFeatures.vue'
import LandingFooter from '@/components/landing/LandingFooter.vue'
import LandingHero from '@/components/landing/LandingHero.vue'
import LandingModels from '@/components/landing/LandingModels.vue'
import LandingNav from '@/components/landing/LandingNav.vue'
import LandingPricing from '@/components/landing/LandingPricing.vue'
import LandingTools from '@/components/landing/LandingTools.vue'
import { useAppStore, useAuthStore } from '@/stores'

const authStore = useAuthStore()
const appStore = useAppStore()

onMounted(() => {
  authStore.checkAuth()

  if (!appStore.publicSettingsLoaded) {
    void appStore.fetchPublicSettings()
  }
})
</script>
