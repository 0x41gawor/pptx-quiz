<template>
  <div class="bye-view" @click="goHome">
    <div class="bye-card">
      <h1>{{ t('bye_title') }}</h1>
      <p>{{ t('bye_subtitle') }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from '@/i18n/lang'
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'

const { t } = useI18n()
const router = useRouter()

function goHome() {
  router.push({ name: 'hello' })
}

onMounted(() => {
  setTimeout(goHome, 10_000)
})
</script>

<style scoped>
.bye-view {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.bye-card {
  text-align: center;
  padding: 2rem 3rem;
  border-radius: 1.5rem;
  background: rgba(15, 23, 42, 0.9);
}
</style>
