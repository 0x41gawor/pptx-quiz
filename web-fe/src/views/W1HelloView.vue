<template>
  <div class="hello-view">
    <div class="hello-card">
      <h1>{{ t('hello_title') }}</h1>
      <p class="subtitle">{{ t('hello_subtitle') }}</p>

      <div class="flags">
        <button class="flag-button" @click="selectLanguage('pl')">
          🇵🇱
          <span>Polski</span>
        </button>
        <button class="flag-button" @click="selectLanguage('en')">
          🇬🇧
          <span>English</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Language } from '@/models/common'
import { useRouter } from 'vue-router'
import { useI18n } from '@/i18n/lang'
import { quizSession } from '@/stores/quizSession'

const router = useRouter()
const { setLanguage, t } = useI18n()

function selectLanguage(lang: Language) {
  setLanguage(lang)
  quizSession.language = lang
  router.push({ name: 'user-data' })
}
</script>

<style scoped>
.hello-view {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--color-bg-main);
}

.hello-card {
  text-align: center;
  padding: 2rem 3rem;
  border-radius: 1.5rem;
  background: rgba(15, 23, 42, 0.9);
}

.subtitle {
  margin-top: 0.5rem;
  margin-bottom: 1.5rem;
  color: var(--color-text-muted);
}

.flags {
  display: flex;
  gap: 1.5rem;
  justify-content: center;
}

.flag-button {
  border: none;
  border-radius: 999px;
  padding: 0.75rem 1.5rem;
  background: var(--color-bg-panel);
  color: var(--color-text-main);
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 1.1rem;
}
</style>
